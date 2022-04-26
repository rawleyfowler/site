#!/bin/sh
# This script is to be run by your init service, such as rc.d
# WIP: Do not use this at the moment.
if [ -e "main" ]; then
	$1/main >> /var/log/rawleyxyz/log 2>&1
else
	go build "$1/main.go"
	$1/main >> /var/log/rawleyxyz/log 2>&1
fi
