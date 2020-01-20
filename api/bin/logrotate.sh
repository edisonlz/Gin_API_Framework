#!/bin/bash

TORNADO_LOG_DIR=/data/logs/tornado
YESTERDAY_LOG_DIR=$TORNADO_LOG_DIR/`date -d yesterday +%Y%m%d`

mkdir $YESTERDAY_LOG_DIR
cp $TORNADO_LOG_DIR/900* $YESTERDAY_LOG_DIR/

echo > $TORNADO_LOG_DIR/9000.log
echo > $TORNADO_LOG_DIR/9001.log
echo > $TORNADO_LOG_DIR/9002.log
echo > $TORNADO_LOG_DIR/9003.log

rm -f $TORNADO_LOG_DIR/9000.log.*
rm -f $TORNADO_LOG_DIR/9001.log.*
rm -f $TORNADO_LOG_DIR/9002.log.*
rm -f $TORNADO_LOG_DIR/9003.log.*

rm -rf /data/datika/*