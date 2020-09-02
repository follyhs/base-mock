package log

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/glog"
)

const (
	LT_CAPTCHA        = "CA"
	LT_REQUESTBEGIN   = "RB"
	LT_REQUESTEND     = "RE"
	LT_MYSQL          = "MY"
	LT_REDIS          = "RS"
	LT_AE             = "AE"
	LT_RG             = "RG"
	LT_LG             = "LG"
	LT_COMMON         = "CM"
	LT_UNEXCEPTED     = "UN"
	LT_TRACK          = "TK"
	LT_SMS_ANTI_FRAUD = "SMSANTIFRAUD"
	LT_BE_POST        = "BEPOST"
	LT_BE_REQ_DATA    = "REQ_DATA"
	LT_BE_PRO_START   = "START"
	LT_BE_PRO_END     = "END"
	LT_BE_GET_VALUE   = "GET_VALUE"
	LT_BE_INTERBYPATH = "GET_INTER_BY_PATH"
	LT_BE_DECODE      = "DECODE"
	LT_BE_PRODUCTER   = "PRODUCTER"
	LT_BE_RECORD      = "BERECORD"
)

const (
	LL_TRACE = "TRACE"
	LL_DEBUG = "DEBUG"
	LL_INFO  = "INFO"
	LL_WARN  = "WARN"
	LL_ERROR = "ERROR"
	LL_FATAL = "FATAL"
)

type CommonInfo struct {
	Level     string
	LogType   string
	RequestId string
}

func CreateCi(level string, ly string, ri string) CommonInfo {
	return CommonInfo{Level: level, LogType: ly, RequestId: ri}
}

func (this *CommonInfo) ToString() string {
	return fmt.Sprintf("logLev=[%s]\tobj=%s\treqId=%s\t", this.Level, this.LogType, this.RequestId)
}

type RequestBegin struct {
	Ci         CommonInfo
	RequestUri string
	RequestIp  string
	Params     string
}

func CreateRb(ci CommonInfo, ru string, ri string, params string) RequestBegin {
	return RequestBegin{Ci: ci, RequestUri: ru, RequestIp: ri, Params: params}
}

func (this *RequestBegin) ToString() string {
	return this.Ci.ToString() + fmt.Sprintf("reqUri=%s\treqIp=%s\treqParams=%s", this.RequestUri, this.RequestIp, this.Params)
}

type RequestEnd struct {
	Ci         CommonInfo
	RequestUri string
	RequestIp  string
	AccessKey  string
	Params     string
	ReturnCode int
	DataToUser string
	Cost       float64
}

func CreateRe(ci CommonInfo, ru string, ri string, ak string, params string, rc int, dtu string, cost float64) RequestEnd {
	return RequestEnd{Ci: ci, RequestUri: ru, RequestIp: ri, AccessKey: ak, Params: params, ReturnCode: rc, DataToUser: dtu, Cost: cost}
}

func (this *RequestEnd) ToString() string {
	return this.Ci.ToString() + fmt.Sprintf("reqUri=%s\treqIp=%s\taccKey=%s\treqParams=%s\tretData=%s\tretCode=%d\tcost=%.3f", this.RequestUri, this.RequestIp, this.AccessKey, this.Params, this.DataToUser, this.ReturnCode, this.Cost)
}

type MysqlSucc struct {
	Ci          CommonInfo
	Description string
	Sql         string
	Cost        float64
}

func CreateMs(ci CommonInfo, des string, sql string, cost float64) MysqlSucc {
	return MysqlSucc{Ci: ci, Description: des, Sql: sql, Cost: cost}
}

func (this *MysqlSucc) ToString() string {
	return this.Ci.ToString() + fmt.Sprintf("desc=%s\tsql=%s\tcost=%.3f", this.Description, this.Sql, this.Cost)
}

type MysqlFail struct {
	Ci          CommonInfo
	Description string
	ErrorInfo   string
	Cost        float64
}

func CreateMf(ci CommonInfo, des string, ei string, cost float64) MysqlFail {
	return MysqlFail{Ci: ci, Description: des, ErrorInfo: ei, Cost: cost}
}

func (this *MysqlFail) ToString() string {
	return this.Ci.ToString() + fmt.Sprintf("desc=%s\terr=%s\tcost=%.3f", this.Description, this.ErrorInfo, this.Cost)
}

type ThirdPartSucc struct {
	Ci          CommonInfo
	SubType     string
	ReturnValue string
	Cost        float64
}

func CreateTps(ci CommonInfo, st string, rv string, cost float64) ThirdPartSucc {
	return ThirdPartSucc{Ci: ci, SubType: st, ReturnValue: rv, Cost: cost}
}

func (this *ThirdPartSucc) ToString() string {
	return this.Ci.ToString() + fmt.Sprintf("subType=%s\tretData=%s\tcost=%.3f", this.SubType, this.ReturnValue, this.Cost)
}

type ThirdPartFail struct {
	Ci        CommonInfo
	SubType   string
	ErrorInfo string
	Cost      float64
}

func CreateTpf(ci CommonInfo, st string, ei string, cost float64) ThirdPartFail {
	return ThirdPartFail{Ci: ci, SubType: st, ErrorInfo: ei, Cost: cost}
}

func (this *ThirdPartFail) ToString() string {
	return this.Ci.ToString() + fmt.Sprintf("subType=%s\terr=%s\tcost=%.3f", this.SubType, this.ErrorInfo, this.Cost)
}

type Log struct{}

func LogInit() {
	flag.Parse()
}

func (this *Log) LogRequestBeginP(level string, ly string, reqId string, ru string, ri string, params string) {
	ci := CreateCi(level, ly, reqId)
	rb := CreateRb(ci, ru, ri, params)
	go glog.Info(rb.ToString())
}

func (this *Log) LogRequestEndP(level string, ly string, reqId string, ru string, ri string, ak string, params string, rc int, dtu string, cost float64) {
	ci := CreateCi(level, ly, reqId)
	re := CreateRe(ci, ru, ri, ak, params, rc, dtu, cost)
	go glog.Info(re.ToString())
}

func (this *Log) LogMysqlSuccP(level string, ly string, reqId string, des string, sql string, cost float64) {
	ci := CreateCi(level, ly, reqId)
	ms := CreateMs(ci, des, sql, cost)
	go glog.Info(ms.ToString())
}

func (this *Log) LogMysqlFailP(level string, ly string, reqId string, des string, err string, cost float64) {
	ci := CreateCi(level, ly, reqId)
	mf := CreateMf(ci, des, err, cost)
	go glog.Info(mf.ToString())
}

func (this *Log) LogThirdPartSuccP(level string, ly string, reqId string, st string, rv string, cost float64) {
	ci := CreateCi(level, ly, reqId)
	tps := CreateTps(ci, st, rv, cost)
	go glog.Info(tps.ToString())
}

func (this *Log) LogThirdPartFailP(level string, ly string, reqId string, st string, err string, cost float64) {
	ci := CreateCi(level, ly, reqId)
	tpf := CreateTpf(ci, st, err, cost)
	go glog.Info(tpf.ToString())
}

func (this *Log) LogRequestBegin(rb RequestBegin) {
	go glog.Info(rb.ToString())
}
func (this *Log) LogRequestEnd(re RequestEnd) {
	go glog.Info(re.ToString())
}
func (this *Log) LogMysqlSucc(ms MysqlSucc) {
	go glog.Info(ms.ToString())
}
func (this *Log) LogMysqlFail(mf MysqlFail) {
	go glog.Info(mf.ToString())
}
func (this *Log) LogThirdPartSucc(tps ThirdPartSucc) {
	go glog.Info(tps.ToString())
}
func (this *Log) LogThirdPartFail(tpf ThirdPartFail) {
	go glog.Info(tpf.ToString())
}

type CommonEngineBegin struct {
	RequestId string
	Params    string
	HostPort  string
}

type CommonEngineError struct {
	RequestId string
	ErrorInfo string
	Cost      float64
}

type CommonEngineEnd struct {
	RequestId   string
	ReturnValue string
	Cost        float64
}

type AEError struct {
	Cee CommonEngineError
}

type AEInvokeEnd struct {
	Cee CommonEngineEnd
}

type AEFatal struct {
	Cees [5]AEError
}

type PIInvokeBegin struct {
	Ceb CommonEngineBegin
}

type PIError struct {
	Cee CommonEngineError
}

type PIInvokeEnd struct {
	Cee CommonEngineEnd
}

type PIFatal struct {
	Cees [5]PIError
}

type RTSInvokeBegin struct {
	Ceb CommonEngineBegin
}

type RTSError struct {
	Cee CommonEngineError
}

type RTSInvokeEnd struct {
	Cee CommonEngineEnd
}

type RTSFatal struct {
	Cees [5]RTSError
}

func (this *Log) LogConfigLoadSuccess(file string) {
	this.LogThirdPartSuccP(LL_INFO, LT_COMMON, "loadconfigsuccess", "", "", 0)
}

func (this *Log) LogConfigLoadError(errInfo string) {
	go glog.Error(errInfo)
	this.LogThirdPartFailP(LL_WARN, LT_COMMON, "loadconfigfailed", "", "", 0)
}

func (this *Log) LogRequestHeader(requestId string, header map[string][]string) {
	headerStr, err := json.Marshal(&header)
	if err == nil {
		info := "RequestId:" + requestId + "\tRequestHeader:" + string(headerStr)
		go glog.Info(info)
	}
}

func (this *Log) Fatal(args ...interface{}) {
	ci := CreateCi(LL_FATAL, LT_COMMON, "0")
	logInfo := fmt.Sprintf("%sinfo=%v", ci.ToString(), args)
	glog.Infoln(logInfo)
	glog.Flush()
}

func (this *Log) Debug(rid string, args ...interface{}) {
	ci := CreateCi(LL_DEBUG, LT_COMMON, "0")
	logInfo := fmt.Sprintf("%sinfo=%v", ci.ToString(), args)
	glog.Infoln("requestId: "+rid, logInfo)
	glog.Flush()
}

func (this *Log) DebugM(rid string, inf string, m map[string]interface{}) {
	ci := CreateCi(LL_DEBUG, LT_COMMON, "0")
	mjson, _ := json.Marshal(m)
	mString := "requestId:" + rid + "  " + inf + " : " + string(mjson)
	logInfo := fmt.Sprintf("%sinfo=%v", ci.ToString(), mString)
	glog.Infoln(logInfo)
	glog.Flush()
}
