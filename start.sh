#!/bin/sh

MY_DIR=$(cd $(dirname $0); pwd)
#PID_FILE="${0%.*}.pid"
PID_FILE=$MY_DIR/gusgen_siren_bot.pid

if [[ -e $PID_FILE ]]; then
  echo "already running"
  exit 1
fi

nohup bundle exec ruby gusgen_siren_bot.rb </dev/null >/dev/null 2>&1 &
echo -n $! >$PID_FILE
