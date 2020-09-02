package tupu

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	//"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"mime/multipart"
	"net/http"
	//"sort"
	"net"
	"strings"
	"time"
)

type Request struct {
	url        string
	secretId   string
	privateKey *rsa.PrivateKey
}

func NewRequest(url string, secretId string, privateKey *rsa.PrivateKey) (req *Request) {
	req = &Request{
		url:        url,
		secretId:   secretId,
		privateKey: privateKey,
	}
	return
}

func (req *Request) CheckSingleImage(imgBuf *bytes.Buffer, imgName string, connTimeout time.Duration, readTimeout time.Duration) (string, error) {
	//timestamp := time.Now().Format(time.RFC1123)
	timestamp := fmt.Sprintf("%v", time.Now().UnixNano()/1000000)

	nonce_int, err := rand.Int(rand.Reader, big.NewInt(9999999999))
	if err != nil {
		return "", err
	}
	nonce := nonce_int.String()
	sign_params := []string{
		//imgHashStr,
		req.secretId,
		timestamp,
		nonce,
	}
	//sort.Strings(sign_params)
	signStr := strings.Join(sign_params, ",")

	sign_hash_bytes := sha256.Sum256([]byte(signStr))
	sign_hash := sign_hash_bytes[:]

	sign, err := rsa.SignPKCS1v15(rand.Reader, req.privateKey, crypto.SHA256, sign_hash)
	if err != nil {
		return "", err
	}

	sign_base64 := base64.StdEncoding.EncodeToString(sign)

	tupu_req_buffer := new(bytes.Buffer)
	multipart_writer := multipart.NewWriter(tupu_req_buffer)
	if err != nil {
		return "", err
	}
	//multipart_writer.WriteField("secretId", req.secretId)
	image_multipart_writer, err := multipart_writer.CreateFormFile("image", imgName)
	io.Copy(image_multipart_writer, imgBuf)
	multipart_writer.WriteField("timestamp", timestamp)
	multipart_writer.WriteField("nonce", nonce)
	multipart_writer.WriteField("signature", sign_base64)
	multipart_writer.Close()
	tupu_req, err := http.NewRequest("POST", req.url, tupu_req_buffer)
	if err != nil {
		return "", err
	}
	tupu_req.Header.Set("Content-Type", multipart_writer.FormDataContentType())
	tupu_req.Header.Set("Expect", "100-continue")
	tupu_req.Header.Set("Content-Length", fmt.Sprintf("%v", tupu_req_buffer.Len()))
	//var client http.Client
	client := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, connTimeout)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(readTimeout))
				return c, nil
			},
		},
	}
	tupu_resp, err := client.Do(tupu_req)
	if err != nil {
		return "", err
	}

	defer tupu_resp.Body.Close()
	if tupu_resp.StatusCode != http.StatusOK {
		return "", err
	}
	tupu_resp_bytes, err := ioutil.ReadAll(tupu_resp.Body)
	return string(tupu_resp_bytes), err
}
