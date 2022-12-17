package main

import (
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.threads()
	if err == nil {
		generateHTML(writer, threads, "layout", "public.navbar", "index")
	}
}
