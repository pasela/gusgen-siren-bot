#!/bin/sh

MY_DIR=$(cd $(dirname $0); pwd)
#PID_FILE="${0%.*}.pid"
PID_FILE=$MY_DIR/gusgen_siren_bot.pid

if [[ -e $PID_FILE ]]; then
  kill $(cat $PID_FILE)
  rm $PID_FILE
else
  echo "not running"
  exit 1
fi
