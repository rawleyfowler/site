#!/bin/sh
# Copyright 2022, Rawley Fowler
# This script is to be run by your init service, such as rc.d
# WIP: Do not use this at the moment.
if [ -e "rawleydotxyz" ]; then
	./rawleydotxyz >> /var/log/rawleydotxyz 2>&1
else
	go build "$1/main.go"
	./rawleydotxyz >> /var/log/rawleydotxyz 2>&1
fi
