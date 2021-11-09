#!/bin/bash

NAME="busybox"
BIN="${BASH_SOURCE-$0}"
BIN="$(dirname "${BIN}")"
BINDIR="$(cd "${BIN}"; pwd)"
BINFILE="$BINDIR/$NAME"
BASEDIR="$(dirname "${BINDIR}")"
LOGSDIR="$BASEDIR/logs"
SENSORDIR="$BASEDIR/sensor"

export PATH=$PATH:$BASEDIR/bin

if [ ! -d "$LOGSDIR" ]; then
    mkdir -p "$LOGSDIR"
fi
LOGFILE="$LOGSDIR/out.log"

if [ ! -d "$SENSORDIR" ]; then
    mkdir -p "$SENSORDIR"
fi
PIDFILE="$SENSORDIR/busybox.pid"


start() {
  if [ -f "$PIDFILE" ] && kill -0 $(cat "$PIDFILE"); then
    echo 'Service is already running' >&2
    return 1
  fi
  # Disable core dumps
    ulimit -c 0

    # Omits the goroutine stack traces entirely.
    export GOTRACEBACK=none
  echo 'Starting service…' >&2
  cd $BASEDIR
  nohup "$BINFILE" >> "$LOGFILE" 2>&1 < /dev/null &
  cd - 2>/dev/null
  if [ $? -eq 0 ]
    then
      pid=$!
      if [ $? -eq 0 ];
      then
        sleep 1
        /bin/echo -n "${pid}" > "$PIDFILE"
        if ps -p "${pid}" > /dev/null 2>&1; then
          echo STARTED
        else
          echo FAILED TO START
          exit 1
        fi
      else
        echo FAILED TO WRITE PID
        exit 1
      fi
    else
      echo SERVER DID NOT START
      exit 1
    fi
}

stop() {
  if [ ! -f "$PIDFILE" ] || ! kill -0 $(cat "$PIDFILE"); then
    echo 'Service is not running' >&2
    return 1
  fi
  echo 'Stopping service…' >&2
  kill $(cat "$PIDFILE") && rm -f "$PIDFILE"
  echo 'Service stopped' >&2
}

restart() {
  stop
  sleep 2
  start
}

status() {
  if [ -f "$PIDFILE" ] && kill -0 $(cat "$PIDFILE"); then
    echo "Service is running with pid = $(cat "$PIDFILE")" >&2
    return 0
  fi
  if [ ! -f "$PIDFILE" ] || ! kill -0 $(cat "$PIDFILE"); then
    echo 'Service stopped' >&2
    return 0
  fi
}

version() {
  "$BINFILE" version
}

case "$1" in
  start)
    start
    ;;
  stop)
    stop
    ;;
  status)
    status
    ;;
  restart)
    restart
    ;;
  version)
    version
    ;;
  *)
    echo "Usage: $0 {start|stop|restart|status|version}"
esac