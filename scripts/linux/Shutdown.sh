#!/bin/sh

KillApp(){
	echo "查找 $1 "
	if [ $(ps -ef|grep sanfeng |grep ./$1|grep -v grep|awk '{print $2}') ]; then
		echo "kill $1"
		kill -9 $(ps -ef|grep sanfeng |grep ./$1|grep -v grep|awk '{print $2}')
	fi
}

KillApp logger
KillApp center
KillApp config
KillApp gateway
KillApp login
KillApp list
KillApp property
KillApp table
KillApp room
KillApp robot
