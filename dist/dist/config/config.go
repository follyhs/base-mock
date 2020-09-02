package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	ConfigMap *StoreConfig
}

type MysqlConf struct {
	Host         string
	Port         int
	User         string
	Password     string
	DbName       string
	ConnTimeout  int
	ReadTimeout  int
	WriteTimeout int
	MaxOpenConn  int
	MaxIdleConn  int
}
type BasicConf struct {
	ZkServers string
	IdcName   string
}

type StoreConfig struct {
	MysqlC        MysqlConf `json:"MysqlC"`
	MysqlHistoryC MysqlConf `json:"MysqlHistoryC"`
	BasicC        BasicConf `json:"BasicC"`
	StoragerC     string    `json:"StoragerC"`
}

func (this *Config) LoadConfig(filePath string) (string, error) {
	defer func() error {
		if err := recover(); err != nil {
			return errors.New(fmt.Sprintf("%v", err))
		}
		return nil
	}()

	fd, openErr := os.Open(filePath)
	if openErr != nil {
		return "open file", openErr
	}
	defer fd.Close()

	content, readErr := ioutil.ReadAll(fd)
	if readErr != nil {
		return "read file", readErr
	}

	this.ConfigMap = new(StoreConfig)
	if err := json.Unmarshal(content, this.ConfigMap); err != nil {
		return "new config", err
	}
	return "", nil
}
