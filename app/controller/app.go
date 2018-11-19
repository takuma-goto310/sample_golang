package controller

import (
	"html/template"
	"log"
	"net/http"
	)

// Index is method to render Top page.
func Index(rw http.ResponseWriter, request *http.Request) {

	log.Println("call Index")

	err := parseTemplate().ExecuteTemplate(rw, "index.html", "")
	if err != nil {
		outputErrorLog("HTML 描画 エラー", err)
		log.Fatalln("エラーのため強制終了")
	}
}

// parse HTML
func parseTemplate() *template.Template {
	tmpl, err := template.ParseGlob("./../app/view/*.html")
	if err != nil {
		outputErrorLog("HTML パース 失敗", err)
	}
	return tmpl
}

// output error log
func outputErrorLog(message string, err error) {
	log.Println(message)
	log.Println(err)
}