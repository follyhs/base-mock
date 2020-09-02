#!/bin/bash
# Copyright (c) 2019 SHUMEI Inc. All Rights Reserved.
# Authors: zhanghaisong<zhanghaisong@ishumei.com>

set -e
BIN_DIR=$(which $0)
BIN_DIR=$(dirname $BIN_DIR)
BIN_DIR=$(cd $BIN_DIR; pwd)
ROOT_DIR=$(cd $BIN_DIR/../config; pwd)
TEMP_DIR=$BIN_DIR/tmp
MAKE_CONF_SHELL=$BIN_DIR/make_conf.sh


if [ $# -eq 1 ]; then
    RES_TAG=$1
else
    RES_TAG=$ENV_ROLE
fi
GLOBAL_RESOURCE_FILE=/mnt/engine/config-source/config-source/resources/$RES_TAG/global.conf
# GLOBAL_RESOURCE_FILE=/home/zhanghaisong/codes/union_pay/global.conf

if [ ! -f $GLOBAL_RESOURCE_FILE ]; then
    echo "$GLOBAL_RESOURCE_FILE not exist"
    exit -1
fi

# 创建临时目录
[ ! -d $TEMP_DIR ] && mkdir -p $TEMP_DIR
RESOURCE_FILE=$TEMP_DIR/resource.conf
echo "" > $RESOURCE_FILE

# 汇总配置
cat $GLOBAL_RESOURCE_FILE >> $RESOURCE_FILE

# 从下面开始自定义配置
CONF_FILE=$ROOT_DIR/config.json.temp
TARGET_CONF_FILE=$ROOT_DIR/config.json.$RES_TAG
echo "" > $TARGET_CONF_FILE

# 判断文件是否存在
if [ ! -f $CONF_FILE ]; then
    echo "$CONF_FILE not exist"
    exit -1
fi

sh $MAKE_CONF_SHELL $CONF_FILE $RESOURCE_FILE  >> $TARGET_CONF_FILE 

# 建立软链
cd $ROOT_DIR
ln -sf config.json.$RES_TAG config.json

# 清理数据
rm -rf $TEMP_DIR
# rm -rf $CONF_FILE

echo "init success"

