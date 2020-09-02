#!/bin/bash

BIN_DIR=$(dirname $0)
BIN_DIR=$(cd $BIN_DIR; pwd)
ROOT_DIR=$(cd $BIN_DIR/..; pwd)
cd $ROOT_DIR

if [ ! -d $ROOT_DIR/log ]; then
    mkdir $ROOT_DIR/log
fi

HTTPPORT=8588
THRIFTPORT=8868
if [ x$1 != x ]
then
       HTTPPORT=$1
fi
if [ x$2 != x ]
then
       THRIFTPORT=$2
fi
while [ true ]; do
    cd $ROOT_DIR/bin/
    export LD_LIBRARY_PATH=.:${LD_LIBRARY_PATH}
    ./be-record -log_dir=../log -config_file=../config/config.json -http_port=$HTTPPORT -thrift_port=$THRIFTPORT
    echo "$? `date +%Y%m%d%H%M%S`" >> $ROOT_DIR/log/exit_codes
    sleep 3
done
