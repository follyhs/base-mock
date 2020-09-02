// Copyright (c) 2019 SHUMEI Inc. All rights reserved.
// Authors: ybwang <wangyanbo@ishumei.com>.

package main_test

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"os"
	"shumei/recordService/prediction"
	"testing"
	"time"
)

var requestId *string
var serviceId *string
var organization *string
var appId *string
var tokenId *string
var data *string
var stype *string
var timestamp *int64

func init() {
	setup()
}
func setup() {
	requestId = flag.String("request_id", "callfromzhanghaisong", "requestId")
	serviceId = flag.String("service_id", "VERIFY_CAPTCHA", "serviceId")
	stype = flag.String("type", "POST_IMG", "type")
	organization = flag.String("organization", "RlokQwRlVjUrTUlkIqOg", "organization")
	appId = flag.String("app_id", "default", "appId")
	tokenId = flag.String("token_id", "zhanghaisong", "tokenId")
	timestamp = flag.Int64("timestamp", 123353, "timestamp")
	data = flag.String("data", `{"result":{"code":1100,"message":"成功","requestId":"b4d6c39691684f8f35d8a2d8c517f714","riskLevel":"PASS","score":0,"detail":{"token":{"groupSize":0,"riskGrade":"","riskReason":"","riskType":"","score":0,"groupId":""},"description":"正常","hits":[],"model":"M1000"}},"features":{"appId":"shuabao","accessKey":"Lnqwf0ROLrvikIvWi9UQ","eventId":"like","data":{"apdid_token":"","appVersion":"1.850","contentCategory":"","contentLength":0,"contentType":"","deviceId":"20191114140418ad4741bebfca64da314cc57b5cf497c501eec687cc4b8938","eventId":"like","eventName":"video_praise","idfa":"","imei":"","ip":"183.215.51.48","os":"android","referId":"","timestamp":"1576575479945","tokenId":"790309071","watchLength":0}}}`, "request data")
	flag.Parse()
}

func TestPredict(t *testing.T) {
	timestamp := int64(time.Now().UnixNano() / 1000000)

	t.Log("----- PREDICTION START. -----")

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "8868"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport := transportFactory.GetTransport(transport)
	client := prediction.NewPredictorClientFactory(useTransport, protocolFactory)

	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:8868", " ", err)
		os.Exit(1)
	}

	defer transport.Close()

	request := prediction.NewPredictRequest()
	request.RequestId = requestId
	request.ServiceId = serviceId
	request.AppId = appId
	request.Organization = organization
	request.TokenId = tokenId
	request.Timestamp = &timestamp
	request.Data = data
	request.Type = stype
	result, err := client.Predict(request)

	//test
	num := 1
	time.Sleep(time.Duration(1 * time.Second))
	for i := 0; i < num; i++ {
		result, _ := client.Predict(request)
		fmt.Println(result)
	}
	//test end
	// fmt.Println("result detail value:", result.GetDetail())

	t.Log(fmt.Sprintf("requestId=%v, result=%v", *requestId, *result))
}
