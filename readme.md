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

## test func 2

設定ファイルを作成して、適時、設定を読み込みます。この機能は現在テスト中です。

設定は、上記にある`mac:iTerm`でのみ有効なアイコン表示の設定になります。以下、機能のON/OFFの切替方法と使い方を紹介します。有効にするとオプションの数を少なくできます。

現時点では、作成するユーザーファイルに追記する形で運用しています。ただし、オプションを実行しない限り新しい設定は追記されません。このファイルはOAuth認証の際に、認証が通っているか確かめるついでにAPIから取得するユーザーのプロファイルです。

`~/.config/twg/verify.json` : `twg_icon`

[https://api.twitter.com/1.1/account/verify_credentials.json](https://api.twitter.com/1.1/account/verify_credentials.json)

```sh
# 有効にする
$ twg set true
$ twg t
# タイムラインを表示する際にユーザーアイコンも表示される(mac:iTerm限定)

# 無効にする
$ twg set false
$ twg t

# 機能のON/OFFをチェックする
$ twg set c
iterm-mode/check false
```

`tmux`を使用している場合は、うまく表示できないことがあります。もし端末が乱れた場合は、`reset`コマンドで修正できます。

## link

[https://github.com/syui/twg](https://github.com/syui/twg)

[https://gitlab.com/syui/twg](https://gitlab.com/syui/twg)

[https://aur.archlinux.org/packages/twg](https://aur.archlinux.org/packages/twg)

## ref

https://github.com/mrjones/oauth/blob/master/examples/twitter/twitter.go
