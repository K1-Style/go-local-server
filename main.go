package main

import (
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

// 指定したディレクトリに対してローカルでhttpサーバーを立ち上げるスクリプト
//
// 例) go run main.go -d /path/to/dir -p 8080
// ↑http://localhost:8080 にアクセスすると /path/to/dir 内の内容を確認できる
func main() {
	// プログラム実行したカレントディレクトリをデフォルトとする
	currentDir, _ := os.Getwd()

	app := cli.NewApp()
	app.Name = "go-local-server"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "dir, d",
			Value: currentDir,
			Usage: "directory for launch server",
		},
		cli.StringFlag{
			Name:  "port, p",
			Value: "9090",
			Usage: "port number",
		},
	}

	app.Action = func(c *cli.Context) error {
		http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(c.String("dir")))))
		return http.ListenAndServe(":"+c.String("port"), nil)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
