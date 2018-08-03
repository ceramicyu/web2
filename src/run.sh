#!/bin/sh
go build main.go

kill -9 $(ps |grep main)
kill -9 $(ps |grep main)
./main &
kill -9 $(ps |grep run.sh)