# go-local-server

goで立てるローカルサーバー

## インストール

```shell
$ go get github.com/K1-Style/go-local-server
```

## 実行例

`/path/to/dir`ディレクトリを http://localhost:8080 でアクセスしたい場合

```shell
$ go-local-server -d /path/to/dir -p 8080
```

## オプション

* `--dir`, `-d`
    * localhostでhttpアクセスしたい対象ディレクトリ(デフォルト:コマンド実行したカレントディレクトリ)
* `--port`, `-p`
    * localhostのポート番号(デフォルト:9090)
