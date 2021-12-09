package main

import (
	"embed"
	"log"
	"net/http"
	"os"
)

//go:embed buy categories code css images posts tags toc  
//go:embed book_cover.jpg index.html index.xml sitemap.xml  
var siteData embed.FS

func main() {
	listenAddr := ":8080"
	if len(os.Getenv("LISTEN_ADDR")) != 0 {
		listenAddr = os.Getenv("LISTEN_ADDR")

	}
	mux := http.NewServeMux()
	staticFileServer := http.FileServer(http.FS(siteData))
	mux.Handle("/", staticFileServer)

	log.Fatal(http.ListenAndServe(listenAddr, mux))
}
