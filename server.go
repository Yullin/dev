package main

import (
    "os"
    "strings"
    "fmt"
    "net/http"
)

func HealthResponse(w http.ResponseWriter, r *http.Request) {
	remote_ip := strings.Split(r.RemoteAddr, ":")[0]
	fmt.Println(remote_ip, "200", r.RequestURI)
    fmt.Fprintln(w, "")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	for k,v := range r.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	w.Header().Set("Version", os.Getenv("VERSION"))
	remote_ip := strings.Split(r.RemoteAddr, ":")[0]
	fmt.Println(remote_ip, "200", r.RequestURI)
    fmt.Fprintln(w, "hello there, You are from", r.RemoteAddr)
}

func main() {
    http.HandleFunc("/", IndexHandler)
    http.HandleFunc("/healthz", HealthResponse)
    http.ListenAndServe(":8000", nil)
}
