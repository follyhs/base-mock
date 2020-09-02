// Copyright (c) 2019 SHUMEI Inc. All rights reserved.
// Authors: zhanghaisong<zhanghaisong@ishumei.com>.

package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net/http"
	"shumei/mockService"
	"shumei/mockService/prediction"
	"shumei/zookeeper"
)

var (
	conf           *config.Config
	logger         *log.Log
	configFilePath *string
	httpPort       *string
	thriftPort     *string
	server         *thrift.TSimpleServer
)

func loadConfig() {
	defer func() {
		if err := recover(); err != nil {
			logger.LogThirdPartFailP(log.LL_WARN, log.LT_UNEXCEPTED, "", "CONFIGLOAD", fmt.Sprintf("%v", err), 0)
		}
	}()

	if str, err := conf.LoadConfig(*configFilePath); err != nil {
		logger.LogConfigLoadError(str + "," + err.Error())
	} else {
		logger.LogConfigLoadSuccess(*configFilePath)
	}
}

func commandLine() {
	configFilePath = flag.String("config_file", "", "config file")
	httpPort = flag.String("http_port", "80", "http server port")
	thriftPort = flag.String("thrift_port", "80", "thrift server port")
	flag.Parse()
}

func main() {
	commandLine()

	//init log
	log.LogInit()
	logger = new(log.Log)

	conf = new(config.Config)
	loadConfig()

	go func() {
		http.ListenAndServe("0.0.0.0:8080", nil)
	}()
	go thriftServer()
	time.Sleep(time.Second)
	zookeeper.SetZkServers(conf.ConfigMap.BasicC.ZkServers, logger)
	zookeeper.Register("/public/record", *thriftPort)
	logger.LogThirdPartFailP(log.LL_FATAL, log.LT_COMMON, "", "app start or restart, please check.", "", 0)

	logger.LogThirdPartFailP(log.LL_FATAL, log.LT_COMMON, "", "stop ok", "", 0)
}

func thriftServer() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocalFactory := thrift.NewTBinaryProtocolFactory(true, true)
	serverAddr := "0.0.0.0:" + *thriftPort
	serverTransport, err := thrift.NewTServerSocket(serverAddr)
	if err != nil {
		logger.LogConfigLoadError(err.Error())
		return
	}
	handler := &recordService.RecordService{Conf: conf, Logger: logger, Storager: rStorager, PdStorager: pdStorager, CsStorager: csStorager}
	processor := prediction.NewPredictorProcessor(handler)
	server = thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocalFactory)
	if err := server.Serve(); err != nil {
		logger.LogConfigLoadError(err.Error())
		server = nil
	}
}
