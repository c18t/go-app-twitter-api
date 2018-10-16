# Project Making

Gatapiのプロジェクトを作成した際のコマンド実行ログです。

```sh
$ ghq get git@github.com:c18t/go-app-twitter-api.git
$ cd $(ghq root)/github.com/c18t/go-app-twitter-api
$ gitignore code node go gitbook windows macos
$ touch README.md
$ touch LICENSE
$ yarn init
$ yarn add -D gitbook-cli
$ mkdir -p ./doc/styles
$ touch ./doc/book.json
$ touch ./doc/styles/pdf.css
$ yarn doc:install
$ gitbook init ./doc
$ touch usage.md
$ touch change_log.md
$ touch project_making.md
$ mkdir -p ./cmd/gatapi
$ touch ./cmd/gatapi/main.go
$ cobra init github.com/c18t/go-app-twitter-api/gatapi
$ dep init
```

