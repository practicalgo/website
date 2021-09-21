package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed img sass posts index.html about
var siteData embed.FS

func main() {
	mux := http.NewServeMux()
	staticFileServer := http.FileServer(http.FS(siteData))
	mux.Handle("/", staticFileServer)

	log.Fatal(http.ListenAndServe(":80", mux))
}
