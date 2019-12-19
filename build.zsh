#!/bin/zsh

f=${0:a:h}/config.json

if [ ! -f $f ];then
	echo no file $f
	exit
fi

export CKEY=`cat $f| jq -r ".consumer_key"`
export CSKEY=`cat $f| jq -r ".consumer_secret"`
echo $CKEY
go build -ldflags="-X github.com/syui/twg/oauth.ckey=$CKEY -X github.com/syui/twg/oauth.cskey=$CSKEY"
