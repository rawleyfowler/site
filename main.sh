#!/bin/sh

if [ ./main & ]; then
	exit 0
else
	exit 1
fi
