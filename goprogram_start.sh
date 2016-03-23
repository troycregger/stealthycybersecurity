#!/bin/bash
case $1 in
    start)
        echo "Starting hello web app."
        /var/www/Golang/src/stealthycybersecurity/main &
        ;;
    stop)
        echo "Stopping hello web app."
        sudo kill $(sudo lsof -t -i:8080)
        ;;
    *)
        echo "Hello web app service."
        echo $"Usage $0 {start|stop}"
        exit 1
esac
exit 0
