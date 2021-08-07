#!/bin/zsh
set -eu
d=${0:a:h}

if [ -z "$1" ];then
	exit
else
	img=$1
fi

if [ -n "$2" ];then
	mes=$2
fi

if ! which twg;then
	exit
fi

if ! which img2sixel;then
	exit
fi

img2sixel $img

echo twg i $img "$mes"
echo "[y]"
read key

if [ "$key" = "y" ];then
	twg i $img "$mes"
fi

