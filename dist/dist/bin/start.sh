#!/bin/bash

set -e

BIN_DIR=$(dirname $0)
BIN_DIR=$(cd $BIN_DIR; pwd)
ROOT_DIR=$(cd $BIN_DIR/..; pwd)
cd $ROOT_DIR

HTTPPORT=8588
if [ x$1 != x ]
then
       HTTPPORT=$1
fi

THRIFTPORT=8868
if [ x$2 != x ]
then
       THRIFTPORT=$1
fi
export GOGC=800
nohup $ROOT_DIR/bin/be-record-run.sh $HTTPPORT $THRIFTPORT >$ROOT_DIR/log/nohup.out 2>&1 &
