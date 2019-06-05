package main

import (
	"github.com/stiei-iss/isDrive/app/init/config"
	"net/http"
)

func main() {
	RunWebServer()
}
func RunWebServer(){
	http.HandleFunc("/init", config.InitConfigHandler)
	http.ListenAndServe(":8219", nil)
}
//
//func ConnectDB(url string) {
//	conn.Connect(url)
//}