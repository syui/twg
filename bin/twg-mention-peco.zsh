#!/bin/zsh
set -eu
d=${0:a:h}
if ! which twg;then
	exit
fi

if [ -n "$1" ];then
	o=$1
else
	o=t
fi

if [ -n "$2" ];then
	s=$2
else
	s=30
fi

tmp=`twg m $o $s | peco`
id=`echo $tmp | awk -F ' ' '{print $(NF -1)}'`
u=`echo $tmp | awk -F ' ' '{print $1}'`

rp=`echo $id|cut -d ' ' -f 1`
echo rp $rp
tw=`echo $id|cut -d ' ' -f 2`
echo tw $tw

echo message input:
vim $d/mes.txt
mes=`cat $d/mes.txt`

echo twg mm $tw "@${u} $mes[y]"
read key
if [ "$key" = "y" ];then
	twg mm $tw "@${u} $mes"
fi

if [ -f $d/mes.txt ];then
	rm $d/mes.txt
fi
