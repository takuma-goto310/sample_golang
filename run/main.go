package main

import (
	"net/http"

	"../app/controller"
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
