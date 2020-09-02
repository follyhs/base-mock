#!/bin/bash

set -e 

cnt=`ps aux | grep 'bin/be-record-run.sh' | grep -v grep | wc -l`

if [ $cnt -eq 1 ]; then
	echo "[OK] pid exists"
else
	echo "[ERROR] cnt=$cnt pid not exists"
	exit 1
fi

cnt=`ps aux | grep '/be-record -log_dir=../log -config_file=../config/config.json' | grep -v grep | wc -l`

if [ $cnt -eq 1 ]; then
	echo "[OK] pid exists"
else
	echo "[ERROR] cnt=$cnt pid not exists"
	exit 1
fi

# for i in 7750
# do
# curl -H 'm_request_id: 67c7bf09c8c0587439e61d6edc77ab18' -H 'm_remote_addr: 192.168.0.1' -H 'm_organization: IkzxwQ4vofwwvFqC8ir2' -d'{"accessKey": "S4AUW8LLJIdomF6uFYHP", "eventId": "login", "accessKey": "AEYAbiH38T44mo7ycsTZ", "data": {"eventName": "login", "ip": "132.45.89.32", "deviceId": "20160929120724e9dc7dcdbf909f898b1ebd1e12bfd39e090a1a6b4f5287340", "timestamp": 1526283969, "tokenId": "blacktokenId", "phone": "17612710729", "valid": 1, "userExist": 1}, "appId": "default"}' 'http://localhost:7750/v2/event' | grep 'code":1100'
# if [ $? -eq 0 ]; then
# 	echo "[OK] curl exec success"
# else
# 	echo "[ERROR] curl exec failure"
# 	exit 1
# fi
# done
