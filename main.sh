#!/bin/sh
# WIP: Do not use this at the moment
if [ -e "main" ]; then
	./main &> /var/log/rawleyxyz/log
else
	go build "main.go" && ./main &> /var/log/rawleyxyz/log
fi
