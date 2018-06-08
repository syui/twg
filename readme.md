# twg

twitter client (golang)

## download 

[releases](https://github.com/syui/twg/releases)


バイナリをダウンロードして、パスが通っている場所に置きます。

```sh
# archlinux
$ curl -SLO https://github.com/syui/twg/releases/download/pre-release/linux_amd64_twg
$ mv linux_amd64_twg twg
$ chmod +x twg
$ echo $PATH
$ mv twg /usr/local/bin
```

## use

```sh
# 認証がない場合、ブラウザを開きます。
$ twg

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
```

## build

`releases`版を使わないと、`consumer_key`などが必要になります。

`config.json` : https://apps.twitter.com/

```sh
$ go get -u -v gitlab.com/syui/twg
$ cd $GOPATH/src/!$
$ cp ./config.json.example config.json
$ vim config.json
$ export CKEY=`cat ./config.json| jq -r ".consumer_key"`
$ export CSKEY=`cat ./config.json| jq -r ".consumer_secret"`
$ go build -ldflags="-X gitlab.com/syui/twg/oauth.ckey=$CKEY -X gitlab.com/syui/twg/oauth.cskey=$CSKEY"
```

## test func

テスト機能として、`iTerm`を使ってる場合は、アイコンを表示することができます。

```sh
$ twg t i
```

`tmux`を使用している場合は、うまく表示できないことがあります。もし端末が乱れた場合は、`reset`コマンドで修正できます。

## link

[https://github.com/syui/twg](https://github.com/syui/twg)

[https://gitlab.com/syui/twg](https://gitlab.com/syui/twg)

[https://aur.archlinux.org/packages/twg](https://aur.archlinux.org/packages/twg)

## ref

https://github.com/mrjones/oauth/blob/master/examples/twitter/twitter.go
