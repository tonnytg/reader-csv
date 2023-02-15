package api

import (
	"net/http"
	"os"
)

func Start() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	loadHandlers()
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
