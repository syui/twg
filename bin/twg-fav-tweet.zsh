#!/bin/zsh
set -eu
d=${0:a:h}
if ! which twg;then
	exit
fi

tmp=`twg m t 100 | peco|awk -F ' ' '{print $(NF -1)}'`
n=`echo "$tmp"|wc -l`
for ((i=1;i<=$n;i++))
do
		t=`echo "$tmp"|awk "NR==$i"`
		twg f $t
done
