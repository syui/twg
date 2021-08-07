# twg

twitter client (golang)

- 0.3 change key, cannot use the 0.1 ~ 0.2 twitter key.

## download 

[releases](https://github.com/syui/twg/releases)

Download the binary and place it where the path passes.

```sh
$ go get -u -v github.com/syui/twg

# archlinux
$ curl -SLO https://github.com/syui/twg/releases/download/pre-release/linux_amd64_twg
$ mv linux_amd64_twg twg
$ chmod +x twg
$ echo $GOPATH
$ mv twg $GOPATH/bin

$ which twg
```

## use

```sh
# If authentication does not exist, open the browser.
$ twg o

# timeline
$ twg

# help
$ twg h

# post
$ twg p "send tweet"

# user timeline
$ twg u syui__
$ twg u

# stream
$ twg s

# oauth
$ twg o

# search
$ twg / "#twitter" 2
```

## build

If you do not use the `releases` version, you will need `consumer_key` etc.

`config.json` : https://apps.twitter.com/

```sh
# go get -u -v github.com/syui/twg
$ go get -v github.com/syui/twg
$ cd $GOPATH/src/!$
$ cp ./config.json.example config.json
$ vim config.json
$ export CKEY=`cat ./config.json| jq -r ".consumer_key"`
$ export CSKEY=`cat ./config.json| jq -r ".consumer_secret"`
$ go build -ldflags="-X github.com/syui/twg/oauth.ckey=$CKEY -X github.com/syui/twg/oauth.cskey=$CSKEY"

```

## help

```sh
# command sub help
$ twg stream help
```

## test func 2

Create a setting file and load the settings as needed. This function is currently being tested.

Setting is the setting of icon display effective only with `mac: iTerm`. Below, we will show you how to switch ON / OFF function and how to use it.

```sh
# iterm-mode : enable
$ twg set help
$ twg set true
$ twg t

# iterm-mode : disable
$ twg set false
$ twg t

# Check ON / OFF of function
$ twg set c
iterm-mode/check false

Delete `~ /.config/twg/img` if the picture icon is wrong. It will be reacquired when it becomes necessary
$ twg set d
```


If you use `tmux`, you may not be able to display it properly. If the terminal is disturbed, you can modify it with `reset` command.

## test func 3

This is a test function.

A command related to acquisition of user information was added.

```sh
# Acquire three latest notifications
$ twg n 3
```

## test func 4

```sh
# Acquire three latest user timeline
$ twg u syui__ 3
```

## config file

`~/.config/twg/verify.json` : `twg_icon`

[https://api.twitter.com/1.1/account/verify_credentials.json](https://api.twitter.com/1.1/account/verify_credentials.json)

## type json

```sh
$ cat ~/.config/twg/verify.json | jq . | gojson
```

## v 0.4.0

```
# mention, reply
$ twg m | peco | awk -F ' ' '{print $NF}'
123456789
# warning : replies to others need to be marked with @user.
$ twg mm 123456789 "@user $message"
or
$ twg m | peco | awk -F ' ' '{print $NF}' | xargs -I {} twg mm {} "$message"
```

## v 0.4.1

```sh
## fav
$ twg m
123456789
$ twg f $tweet_id
or
$ twg m | peco | awk -F ' ' '{print $(NF -1)}' | xargs -I {} twg f {}

## retweet
$ twg r $tweet_id
```

## v 0.4.2

```sh
# tweet id user post
$ twg m u

# tweet id notify(mention)
$ twg m n

# tweet id timeline 
$ twg m t
```

## v 0.4.3

```sh
$ twg m t 100

$ ./bin/twg-mention-peco.zsh 100
```

## v 0.4.4

```sh
# delete tweet
$ twg u syui__ 100 | awk -f ' ' '{print $(nf -1)}'
123456789

$ twg d 123456789
```

if there is a line break, use `C+space` in peco to select it.

## v 0.4.5

```sh
$ sudo pacman -S libsixel
$ img2sixel ~/file/img.jpg

# image tweet
$ twg i ~/file/img.jpg "$message"
```

## link

[https://github.com/syui/twg](https://github.com/syui/twg)

[https://aur.archlinux.org/packages/twg](https://aur.archlinux.org/packages/twg)

```sh
$ sha1sum ./linux_386_twg ./linux_amd64_twg >> PKGBUILD
$ vim PKGBUILD
$ makepkg --printsrcinfo > .SRCINFO
```

## ref

https://github.com/mrjones/oauth/blob/master/examples/twitter/twitter.go
