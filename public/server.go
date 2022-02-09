package main

import (
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/slayer/autorestart"
)

//go:embed buy categories code css images posts support tags toc
//go:embed book_cover.jpg go.mod index.html index.xml sitemap.xml
var siteData embed.FS

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Lshortfile)
	listenAddr := ":8080"
	if len(os.Getenv("LISTEN_ADDR")) != 0 {
		listenAddr = os.Getenv("LISTEN_ADDR")

	}
	mux := http.NewServeMux()
	staticFileServer := http.FileServer(http.FS(siteData))
	mux.Handle("/", staticFileServer)

	srv := http.Server{
		Addr:    listenAddr,
		Handler: mux,
	}

	autorestart.RestartFunc = autorestart.SendSIGUSR2
	restart := autorestart.GetNotifier()
	go func() {
		<-restart
		logger.Printf("Detected change in binary. Restarting.")
	}()

	autorestart.StartWatcher()

	gracehttp.SetLogger(logger)
	logger.Fatalf("Server terminating: %v", gracehttp.Serve(&srv))
}
