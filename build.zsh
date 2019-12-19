#!/bin/zsh

f=${0:a:h}/config.json

if [ ! -f $f ];then
	echo no file $f
	exit
fi

export CKEY=`cat $f| jq -r ".consumer_key"`
export CSKEY=`cat $f| jq -r ".consumer_secret"`
go build -ldflags="-X gitlab.com/syui/twg/oauth.ckey=$CKEY -X gitlab.com/syui/twg/oauth.cskey=$CSKEY"
