#!/bin/zsh

v=0.2
d=${0:a:h}
f=$d/PKGBUILD
a=$d/linux_386_twg
b=$d/linux_amd64_twg
url=https://github.com/syui/twg/releases/download
url=$url/$v

curl -SLO $url/linux_386_twg
curl -SLO $url/linux_amd64_twg

cp -rf $f $f.back

if [ "`cat $f | awk "NR==13" |cut -d = -f 1`" != "sha1sums" ];then
	exit
fi

sed -i '13d' $f.back
cat $f.back
as=`sha1sum $a|cut -d ' ' -f 1`
bs=`sha1sum $b|cut -d ' ' -f 1`
echo $bs
sed -i "12a sha1sums=('$as'\ '$bs')" $f.back
cat $f.back
echo e push
read k 

mv $f.back $f
makepkg --printsrcinfo > .SRCINFO

cd $d
git add .
git commit -m "up"
git push
