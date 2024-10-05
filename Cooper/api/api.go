package api

import (
	"net/http"

	handlers "example.com/coomper/api/handlers"
)

func InitServer(port string) {
	setHandlers()

	err := http.ListenAndServeTLS(port, "server.crt", "server.key", nil)
	print(err.Error())
}

func setHandlers() {
	http.HandleFunc("/test", handlers.Test)

	http.HandleFunc("/init", handlers.Initial)

	http.HandleFunc("/command", handlers.Command)
}
