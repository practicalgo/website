package main

import (
	"embed"
	"log"
	"net/http"
	"os"
)

//go:embed categories code css images posts tags  
//go:embed Caddyfile DEPLOY.md go.mod index.html index.xml practicalgowebsite.service server.go sitemap.xml  
var siteData embed.FS

func main() {
	listenAddr := ""
	if len(os.Getenv("LISTEN_ADDR")) != 0 {
		listenAddr = os.Getenv("LISTEN_ADDR")

	}
	mux := http.NewServeMux()
	staticFileServer := http.FileServer(http.FS(siteData))
	mux.Handle("/", staticFileServer)

	log.Fatal(http.ListenAndServe(listenAddr, mux))
}
