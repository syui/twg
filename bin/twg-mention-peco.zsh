#!/bin/zsh
set -eu
d=${0:a:h}
if ! which twg;then
	exit
fi

tmp=`twg m t | peco`
id=`echo $tmp | awk -F ' ' '{print $(NF -1)}'`
u=`echo $tmp | awk -F ' ' '{print $1}'`

rp=`echo $id|cut -d ' ' -f 1`
echo rp $rp
tw=`echo $id|cut -d ' ' -f 2`
echo tw $tw

echo message input:
read mes
echo twg mm $tw "@${u} $mes"
read key
twg mm $tw "@${u} $mes"
