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

func setup() {
	requestId = flag.String("request_id", "callfromzhanghaisong", "requestId")
	serviceId = flag.String("service_id", "POST_IMG", "serviceId")
	stype = flag.String("type", "POST_IMG", "type")
	organization = flag.String("organization", "RlokQwRlVjUrTUlkIqOg", "organization")
	appId = flag.String("app_id", "default", "appId")
	tokenId = flag.String("token_id", "zhanghaisong", "tokenId")
	timestamp = flag.Int64("timestamp", 123353, "timestamp")
	data = flag.String("data", `{"appId":"","data":{"appId":"com.xmeng.mrddz","class":"action","data":"","eventId":"captchaFverify","features":{"act":{"c":1925,"cs":0,"d":0.7179020182291667,"h":150,"m":[[0,0,1],[33.936187744140625,0.970062255859375,126],[127.2840576171875,5.723480224609375,210],[190.75701904296875,13.9110107421875,309],[233.36181640625,13.666671752929688,405],[243,12.666671752929688,505],[243,12.666671752929688,603],[242.66668701171875,12.666671752929688,707],[241.5,11.333343505859375,808],[240.83334350585938,11,908],[227.82989501953125,19.251708984375,1008],[216.2347412109375,22.04931640625,1108],[211.39657592773438,22.936767578125,1202],[211,22.666671752929688,1308],[211.33334350585938,21.916793823242188,1402],[211.6666717529297,21.333343505859375,1509],[212,21,1601],[212,21,1702],[213,20.666671752929688,1803],[213.6666717529297,20.666671752929688,1903]],"mh":20,"os":"android","sm":-1,"w":300,"wd":0},"actPass":1,"actTrace":{"count":20,"dupcount":2,"velocityvar":1206.6424158813074,"yvar":48.35377276116982},"c_exists":1,"channel":"DEFAULT","duration":4029,"fverify_version":2,"ip":"163.204.16.110","model":"slide","os":"","os_exists":1,"ostype":"web","protocol":2,"rid":"20200116230612461f1f56bf77166fac","ridPass":1,"sm_exists":1,"stage":"FIRST_VERIFY","weapp_flag":0,"xgbFeatures":{"a_cv":1.782959942486894,"a_end":0.00003769826480434921,"a_inter_range":0.0014208124917200085,"a_max":9223372036854776000,"a_mean":0.0013785290578712122,"a_min":0.000012564580025603953,"a_q1":0.00004471851090079677,"a_q2":0.000347278659042374,"a_q3":0.0014655310026208052,"a_range":9223372036854776000,"a_start":0.01001337641311946,"a_std":0.002457862089738569,"a_x_cv":1.8430816398534389,"a_x_end":0.00003234272569713026,"a_x_inter_range":0.0012013364735216198,"a_x_max":9223372036854776000,"a_x_mean":0.0013394836981327502,"a_x_min":0.0000020752641273878516,"a_x_q1":0.00003234272569713026,"a_x_q2":0.00014082946777343758,"a_x_q3":0.0012336791992187501,"a_x_range":9223372036854776000,"a_x_start":0.009997555514311847,"a_x_std":0.002468777810911458,"a_y_cv":1.2358548500787205,"a_y_end":0.0000330027967396349,"a_y_inter_range":0.00023077517607450032,"a_y_max":9223372036854776000,"a_y_mean":0.00021601405615190265,"a_y_min":0,"a_y_q1":0.0000330027967396349,"a_y_q2":0.00009867834714379641,"a_y_q3":0.0002637779728141352,"a_y_range":9223372036854776000,"a_y_start":0.0005812834655346513,"a_y_std":0.00026696201898050595,"count":20,"diff_cv":1.6937257543974669,"diff_end":0.6666717529296875,"diff_inter_range":14.920844607346337,"diff_max":9223372036854776000,"diff_mean":14.893446887398659,"diff_min":0,"diff_q1":0.47981686219464714,"diff_q2":1.0540909449829114,"diff_q3":15.400661469540985,"diff_range":9223372036854776000,"diff_start":33.95004947545443,"diff_std":25.225414564937896,"diff_t_cv":0.07743892318171774,"diff_t_end":100,"diff_t_inter_range":5,"diff_t_max":9223372036854776000,"diff_t_mean":100.10526315789474,"diff_t_min":84,"diff_t_q1":96,"diff_t_q2":100,"diff_t_q3":101,"diff_t_range":9223372036854776000,"diff_t_start":125,"diff_t_std":7.75204378376985,"diff_x_cv":1.725053862556191,"diff_x_end":0.6666717529296875,"diff_x_inter_range":12.670120239257812,"diff_x_max":9223372036854776000,"diff_x_mean":14.614035355417352,"diff_x_min":0,"diff_x_q1":0.3333282470703125,"diff_x_q2":1,"diff_x_q3":13.003448486328125,"diff_x_range":9223372036854776000,"diff_x_start":33.936187744140625,"diff_x_std":25.20999813739544,"diff_y_cv":1.5436676757517322,"diff_y_end":0,"diff_y_inter_range":1.0889892578125,"diff_y_max":9223372036854776000,"diff_y_mean":1.6330992046155428,"diff_y_min":0,"diff_y_q1":0.2443389892578125,"diff_y_q2":0.5834503173828125,"diff_y_q3":1.3333282470703125,"diff_y_range":9223372036854776000,"diff_y_start":0.970062255859375,"diff_y_std":2.5209624534608777,"t_cv":0.6004431455325837,"t_end":1903,"t_inter_range":1004,"t_max":9223372036854776000,"t_mean":956.4,"t_min":1,"t_q1":505,"t_q2":1008,"t_q3":1509,"t_range":9223372036854776000,"t_start":1,"t_std":574.263824387363,"v_cv":1.8077707736505617,"v_diff_cv":1.6777114174298964,"v_diff_end":0.0037698264804349206,"v_diff_inter_range":0.1423495602374056,"v_diff_max":9223372036854776000,"v_diff_mean":0.1280361368471944,"v_diff_min":0.0011559413623555637,"v_diff_q1":0.004203540024674896,"v_diff_q2":0.034727865904237404,"v_diff_q3":0.14655310026208052,"v_diff_range":9223372036854776000,"v_diff_start":0.8411236187020346,"v_diff_std":0.2148076886321547,"v_diff_x_cv":1.7395069695464616,"v_diff_x_end":0.0032342725697130257,"v_diff_x_inter_range":0.12013364735216199,"v_diff_x_max":9223372036854776000,"v_diff_x_mean":0.12417047087162832,"v_diff_x_min":0.00019507482797445803,"v_diff_x_q1":0.0032342725697130257,"v_diff_x_q2":0.014082946777343758,"v_diff_x_q3":0.12336791992187501,"v_diff_x_range":9223372036854776000,"v_diff_x_start":0.8397946632021951,"v_diff_x_std":0.21599539949306334,"v_diff_y_cv":1.2309959769200136,"v_diff_y_end":0.00330027967396349,"v_diff_y_inter_range":0.022813739634635895,"v_diff_y_max":9223372036854776000,"v_diff_y_mean":0.020821270758398246,"v_diff_y_min":0,"v_diff_y_q1":0.00330027967396349,"v_diff_y_q2":0.009867834714379641,"v_diff_y_q3":0.026114019308599384,"v_diff_y_range":9223372036854776000,"v_diff_y_start":0.04882781110491071,"v_diff_y_std":0.025630900537950564,"v_end":0.006666717529296875,"v_inter_range":0.14888261341882744,"v_max":9223372036854776000,"v_mean":0.1561611443941502,"v_min":0,"v_q1":0.0051240012765824,"v_q2":0.010436544009731796,"v_q3":0.15400661469540985,"v_range":9223372036854776000,"v_start":0.27160039580363543,"v_std":0.28230355281556996,"v_x_cv":1.8399627339361886,"v_x_end":0.006666717529296875,"v_x_inter_range":0.12648827735413898,"v_x_max":9223372036854776000,"v_x_mean":0.15334572465265553,"v_x_min":0,"v_x_q1":0.0035462075091422874,"v_x_q2":0.009900990099009901,"v_x_q3":0.13003448486328126,"v_x_range":9223372036854776000,"v_x_start":0.271489501953125,"v_x_std":0.28215041876932606,"v_y_cv":1.548750042665071,"v_y_end":0,"v_y_inter_range":0.010656071968204511,"v_y_max":9223372036854776000,"v_y_mean":0.016787739801745298,"v_y_min":0,"v_y_q1":0.0025451978047688804,"v_y_q2":0.005452806704512266,"v_y_q3":0.013201269772973392,"v_y_range":9223372036854776000,"v_y_start":0.007760498046875,"v_y_std":0.026000012734203142,"x_cv":0.33115622163258485,"x_end":213.6666717529297,"x_inter_range":29.5,"x_max":9223372036854776000,"x_mean":196.82335052490234,"x_min":0,"x_q1":211.33334350585938,"x_q2":213,"x_q3":240.83334350585938,"x_range":9223372036854776000,"x_start":0,"x_std":65.17927708889249,"y_cv":0.4514025526918007,"y_end":20.666671752929688,"y_inter_range":8.666671752929688,"y_max":9223372036854776000,"y_mean":15.40462646484375,"y_min":0,"y_q1":12.666671752929688,"y_q2":19.251708984375,"y_q3":21.333343505859375,"y_range":9223372036854776000,"y_start":0,"y_std":6.953687709494138}},"get_data":null,"header":null,"method":"","operation":null,"organization":"HhDjIEZD8S9MaDDsYVZZ","post_data":"","requestId":"55288def05f3736c88351ec83d6416af","result":{"code":1100,"detail":{"description":"正常","model":"M1000"},"message":"success","requestId":"55288def05f3736c88351ec83d6416af","riskLevel":"PASS","score":0},"serviceId":"VERIFY_CAPTCHA","tags":null,"timestamp":1579187176199,"tokenId":"","type":"","uri":""},"eventId":"captchaFverify","organization":"HhDjIEZD8S9MaDDsYVZZ","requestId":"55288def05f3736c88351ec83d6416af","serviceId":"VERIFY_CAPTCHA","timestamp":0,"tokenId":"","type":""}`, "request data")
	flag.Parse()
}

func TestPredict(t *testing.T) {
	setup()
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
	for {
		time.Sleep(time.Duration(1 * time.Second))
		for i := 0; i < num; i++ {
			result, _ := client.Predict(request)
			fmt.Println(result)
		}
	}
	//test end
	// fmt.Println("result detail value:", result.GetDetail())

	t.Log(fmt.Sprintf("requestId=%v, result=%v", *requestId, *result))
}
