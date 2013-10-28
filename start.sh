#!/bin/sh

MY_DIR=$(cd "$(dirname "$0")"; pwd)
#PID_FILE="${0%.*}.pid"
PID_FILE="$MY_DIR/gusgen_siren_bot.pid"

if [[ -e "$PID_FILE" ]]; then
  if kill -0 "$PID_FILE" >/dev/null 2>&1; then
    echo "already running"
    exit 1
  else
    rm "$PID_FILE"
  fi
fi

cd "$MY_DIR"
nohup bundle exec ruby gusgen_siren_bot.rb </dev/null >/dev/null 2>&1 &
echo -n $! >"$PID_FILE"
