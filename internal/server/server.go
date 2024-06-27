package server

import (
	"imggenerator/configs"
	"log"
	"net/http"
)

func rend(w http.ResponseWriter, msg string) {
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Println(err)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "favicon")
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "img")
}
func pingHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "PONG")
}
func robotsHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "favicon")
}

func Run(conf configs.ConfI) {
	http.HandleFunc("/", imgHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/robots.txt", robotsHandler)
	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":"+conf.GetPort(), nil); err != nil {
		log.Fatal(err)
	}

}
