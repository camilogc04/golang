package webserver

import (
	"net/http"
)

func MiServidorWeb() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)
}

// Interfaces de net/http (Writer y puntero de Request)
func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webserver/index.html")
}
