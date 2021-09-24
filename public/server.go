package main

import (
	"embed"
	"log"
	"net/http"
	"os"
)

//go:embed posts code
//go:embed index.html index.xml sitemap.xml
//go:embed categories css images
var siteData embed.FS

func main() {
	listenAddr := ":80"
	if len(os.Getenv("LISTEN_ADDR")) != 0 {
		listenAddr = os.Getenv("LISTEN_ADDR")

	}
	mux := http.NewServeMux()
	staticFileServer := http.FileServer(http.FS(siteData))
	mux.Handle("/", staticFileServer)

	log.Fatal(http.ListenAndServe(listenAddr, mux))
}
