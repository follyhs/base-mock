#!/bin/bash

set -e

source ~/.bashrc
export GOPATH=pwd
BIN_DIR=$(dirname $0)
BIN_DIR=$(cd $BIN_DIR; pwd)
ROOT_DIR=$(cd $BIN_DIR/..; pwd)
cd $ROOT_DIR
export GOPATH=$ROOT_DIR/..
export GOBIN=$GOPATH/bin

# package the service
if [ ! -d $ROOT_DIR/dist/bin ]; then
    mkdir -p $ROOT_DIR/dist/bin
fi

if [ ! -d $ROOT_DIR/dist/log ]; then
    mkdir -p $ROOT_DIR/dist/log
fi
if [[ -d /home/compile/makepkg_go/go1.9 ]]; then
    export GOROOT=/home/compile/makepkg_go/go1.9
fi
export PATH=$GOROOT/bin:$PATH

cd ../src/
thrift -out shumei/mockService --gen go ../prediction.thrift
go install shumei/main/base-mock.go 
if [ $? -eq 0 ]; then
	echo "[OK] go compile success"
else
	echo "[ERROR] go compile failure !"
	exit 1
fi

cd $ROOT_DIR

cp -rf ../src/shumei/config $ROOT_DIR/dist/

cp -f ../bin/base-mock $ROOT_DIR/dist/bin/

for lib in `ldd ../bin/base-mock | awk '{if (index($3, "/") == 1) print $3}'`; do
	    cp ${lib} ${ROOT_DIR}/dist/bin
done

cp -f $ROOT_DIR/bin/start.sh $ROOT_DIR/dist/bin/
cp -f $ROOT_DIR/bin/stop.sh $ROOT_DIR/dist/bin/
cp -f $ROOT_DIR/bin/check.sh $ROOT_DIR/dist/bin/
cp -f $ROOT_DIR/bin/restart.sh $ROOT_DIR/dist/bin/
cp -f $ROOT_DIR/bin/be-record-run.sh $ROOT_DIR/dist/bin/
cp -f $ROOT_DIR/bin/init.sh $ROOT_DIR/dist/bin/
cp -f $ROOT_DIR/bin/make_conf.sh $ROOT_DIR/dist/bin/

cp -f ../src/VERSION $ROOT_DIR/dist/
