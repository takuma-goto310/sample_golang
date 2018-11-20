package main

import (
	"net/http"

	// "github.com/takuma-goto310/sample_golang/app/controller" // for docker
	"../app/controller" // for local
)

func main() {
	setRoute()
	// 指定したポートをListen
	http.ListenAndServe(":8080", nil)
}

// ルーティング設定
func setRoute() {
	// 静的ファイルのルーティング
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	// TOP画面
	http.HandleFunc("/index", controller.Index)
}
