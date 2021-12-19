package main

import (
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/slayer/autorestart"
)

//go:embed buy categories code css images posts support tags toc
//go:embed book_cover.jpg go.mod index.html index.xml sitemap.xml
var siteData embed.FS

func main() {
	listenAddr := ":8080"
	if len(os.Getenv("LISTEN_ADDR")) != 0 {
		listenAddr = os.Getenv("LISTEN_ADDR")

	}
	mux := http.NewServeMux()
	staticFileServer := http.FileServer(http.FS(siteData))
	mux.Handle("/", staticFileServer)

	// Notifier
	restart := autorestart.GetNotifier()
	go func() {
		<-restart
		log.Printf("Detected change in binary. Restarting.")
	}()

	autorestart.StartWatcher()

	log.Fatal(http.ListenAndServe(listenAddr, mux))
}
