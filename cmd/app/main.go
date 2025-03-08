package main

import (
	"java-code-task/internal/routes"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	routes.Init(mux)
	err := http.ListenAndServe(":3333", mux)
	if err != nil {
		panic(err)
	}
}
