package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// 指定したディレクトリに対してローカルでHttpサーバーを立ち上げるスクリプト
//
// 例) go run main.go -d /path/to/dir -p 8080
// ↑http://localhost:8080 にアクセスすると /path/to/dir 内の内容を確認できる
func main() {
	// プログラム実行したカレントディレクトリをデフォルトとする
	currentDir, _ := os.Getwd()
	var (
		// サーバー立てる対象のディレクトリ
		dir = flag.String("d", currentDir, "directory for launch server")
		// ポート番号
		port = flag.String("p", "9090", "port number")
	)

	flag.Parse()

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(*dir))))
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
