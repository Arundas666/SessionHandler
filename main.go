package main

import (
	"net/http"
	"session-management/handlers"
)

func main() {

	http.HandleFunc("/set", handlers.SetSession)
	http.HandleFunc("/clear", handlers.ClearSession)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/dlt", handlers.DeleteSession)
	http.ListenAndServe(":8080", nil)

}
