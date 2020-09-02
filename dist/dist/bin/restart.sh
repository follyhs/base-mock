#!/bin/bash

BIN_DIR=$(dirname $0)
BIN_DIR=$(cd $BIN_DIR; pwd)
ROOT_DIR=$(cd $BIN_DIR/..; pwd)
cd $ROOT_DIR

sh $ROOT_DIR/bin/stop.sh
sh $ROOT_DIR/bin/start.sh >/dev/null 2>&1
sleep 8
sh $ROOT_DIR/bin/check.sh
if [[ $? -ne 0 ]]; then
        exit -1
fi
echo "[OK] restart success"
