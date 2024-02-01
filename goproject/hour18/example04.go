package main

import "net/http"

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Header.Get("Accept") {
		case "application/json" :
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Write([]byte(`{"message" : "Hello rld"}`))
		case "application/xml" :
			w.Header().Set("Content-Type", "application/xml; charset=utf-8")
			w.Write([]byte(`<?xml version="1.0" encoding="utf-8"?><message>Hello wo</message>`))
		default :
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte("helld\n"))
	}
}

func main () {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8000", nil)
}