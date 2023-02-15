package api

import "net/http"

func ListAllAuthors(w http.ResponseWriter, r *http.Request) {

}

func loadHandlers() {
	http.HandleFunc("/", ListAllAuthors)
}
