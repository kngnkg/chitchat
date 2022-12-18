package main

import (
	"net/http"

	"github.com/kwtryo/chitchat/data"
)

func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		generateHTML(writer, threads, "layout", "public.navbar", "index")
	}
}

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
}
