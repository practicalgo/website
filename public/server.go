package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed about categories css images posts
//go:embed index.html index.xml sitemap.xml
var siteData embed.FS

func main() {
	mux := http.NewServeMux()
	staticFileServer := http.FileServer(http.FS(siteData))
	mux.Handle("/", staticFileServer)

	log.Fatal(http.ListenAndServe(":80", mux))
}
