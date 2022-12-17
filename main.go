package main

import (
	"net/http"
	"time"
)

func main() {
	p("ChitChat", config.Version, "started at", config.Address)

	// handle static assets
	mux := http.NewServeMux()
	// ディレクトリpublicを起点とする
	files := http.FileServer(http.Dir(config.Static))
	// /static/で始まるリクエストURLについて、/static/を削除
	// 残った文字列を名前にもつファイルを探す
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// ルートハンドラ関数は別ファイル

	// index
	mux.HandleFunc("/", index)

	// starting up the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
	// server.ListenAndServeTLS()
}
