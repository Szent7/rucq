package webserver

import (
	"log"
	"net/http"
)

var HubMain *HubStruct = newHub()

func serveChat(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	http.ServeFile(w, r, "templates/chat.html")
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	http.ServeFile(w, r, "templates/index.html")
}

func serveLogin(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	http.ServeFile(w, r, "templates/login.html")
}

func serveReg(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	http.ServeFile(w, r, "templates/registration.html")
}

func serveJS(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	w.Header().Set("Content-Type", "application/javascript")
	http.ServeFile(w, r, "templates/script.js")
}

func StartWebserver(socket string) {
	go HubMain.run()
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/login", serveLogin)
	http.HandleFunc("/registration", serveReg)
	http.HandleFunc("/chat", serveChat)
	http.HandleFunc("/script.js", serveJS)

	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(HubMain, w, r)
	})
	err := http.ListenAndServe(socket, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
