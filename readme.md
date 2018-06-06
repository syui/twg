# twg

twitter client (golang)

## download 

[releases](https://github.com/syui/twg/releases)

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

## ref

https://github.com/mrjones/oauth/blob/master/examples/twitter/twitter.go
