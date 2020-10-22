#!/usr/bin/env bash

SERVICE=./lessismore
PIDFILE=./lessismore.pid

case "${1}" in
start)
    rm -f ./nohup.out "$PIDFILE"
    nohup "$SERVICE" &
    echo "${!}" > "$PIDFILE"
    sleep 5
    cat ./nohup.out
    ;;
stop)
    if [ -f "$PIDFILE" ]; then
        kill -15 $(cat "$PIDFILE") &> /dev/null
        sleep 8
        if test -r /proc/$(cat "$PIDFILE"); then
            echo -e "\033[1;31m${SERVICE} failed to shutdown\033[0m"
            exit -1
        fi
        rm -f "$PIDFILE"
    else
        echo -e "\033[1;31m${PIDFILE} Not Found\033[0m"
        exit -1
    fi
    ;;
restart)
    ${0} stop
    ${0} start
    ;;
status)
    if [ -f "$PIDFILE" ] && [ -r /proc/$(cat "$PIDFILE") ]; then
        echo -e "\033[1;32m${SERVICE} is running, pid = $(cat $PIDFILE)\033[0m"
        echo -e "\033[1;32mcwd    : $(readlink -q /proc/$(cat $PIDFILE)/cwd)\033[0m"
        echo -e "\033[1;32mexe    : $(readlink -q /proc/$(cat $PIDFILE)/exe)\033[0m"
        echo -e "\033[1;32mcmdline: $(cat /proc/$(cat $PIDFILE)/cmdline | xargs -0 echo)\033[0m"
    else
        echo -e "\033[1;32m${SERVICE} is NOT running\033[0m"
    fi
    ;;
*)
   echo "Usage: $0 {start|stop|status|restart}"
esac

exit 0
