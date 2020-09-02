package zookeeper

import (
	"encoding/json"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"shumei/log"
	"strconv"
	"strings"
	"time"
)

var zkservers string
var conn *zk.Conn
var logger *log.Log

type zklistener func()

// 首先调用此方法设置zk地址
// 初始化zk server
func SetZkServers(servers string, log *log.Log, listeners ...zklistener) {
	if log != nil {
		logger = log
	}
	if zkservers != servers {
		zkservers = servers
		initConn()
		for _, listener := range listeners {
			listener()
		}
	}
}
func Register(path, port string) {
	ip := GetLocalIp()
	fmt.Println(ip)
	if ip == "" {
		panic("get local ip error")
	}
	zkKey := ip + ":" + port
	value := make(map[string]interface{})
	value["host"] = ip
	p, _ := strconv.Atoi(port)
	value["port"] = p
	zkVal, _ := json.Marshal(value)
	go func() {
		for {
			var exist bool
			var err error
			registerPath := path + "/" + zkKey
			exist, _, err = conn.Exists(registerPath)
			if err != nil {
				time.Sleep(time.Second)
				continue
			}
			if !exist {
				_, err := conn.Create(registerPath, zkVal, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
				if err != nil {
					fmt.Println(err)
					time.Sleep(time.Second)
					continue
				}

			}
			time.Sleep(10 * time.Second)
		}
	}()
}
func initConn() {
	servers := strings.Split(zkservers, ",")
	conn_, _, err := zk.Connect(servers, time.Second)
	if err != nil {
		panic(err)
	}
	if conn != nil {
		conn.Close()
	}
	conn = conn_
}
func GetConf(path string) (string, error) {
	ret, _, err1 := conn.Get(path)
	if err1 != nil {
		return "", err1
	}
	return string(ret), nil
}
func GetBatchConf(path string) (map[string]string, error) {
	ret := make(map[string]string)
	keys, _, err1 := conn.Children(path)
	if err1 != nil {
		return ret, err1
	}

	for _, key := range keys {
		newPath := path + "/" + key
		val, _, err := conn.Get(newPath)
		if err != nil {
			continue
		}

		ret[key] = string(val)
	}
	return ret, nil
}
