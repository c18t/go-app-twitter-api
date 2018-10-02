# Gatapi - Go Application for Twitter API. (for study)

Gatapi(がたぴ /ɡä täˈpʲi̞/)は、Goの練習に作成したTwitter APIクライアントです。  
何が起きても知らないですけど、ご利用はお好きにどうぞ。

## Usage
```sh
$ gatapi -h
Gatapi - Go Application for Twitter API. (for study)

Usage:
  gatapi (say | tweet) <message> [--to <screen name>] [--id <tweet id>]
  gatapi (rt | retweet) (<tweet id> | --id <tweet id>)
  gatapi (qt | quote) <message> [--id <tweet id>]
  gatapi (f | like) (<tweet id> | --id <tweet id>)

  gatapi show (profile | home | tweet | favorite | friend | follower | list | listed ) [screen name] [-p | --page <n..m>] [-l | -limit <count>]
  gatapi (list-member | list-subscriber) <screen name>'/'<list-slug> [-p | --page <n..m>] [-l | -limit <count>]

  gatapi (-h | --help)
  gatapi (-v | --version)

Options:
  -h --help                       Show this screen.
  -v --version                    Show version.
```

## Install
```sh
$ go get -u github.com/c18t/go-app-twitter-api/cmd/gatapi
```

## License
[WTFPL](./LICENSE)

## Author
Uchi (/ɯ̹t͡ɕʲi/)
  - GitHub: [c18t](http://github.com/c18t)
  - Twitter: [@c18t](https://twitter.com/c18t)
