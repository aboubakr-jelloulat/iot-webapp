package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"inception-of-things/app1/frontend"
)

type identity struct {
	AppName  string
	Hostname string
	PodID    string
}

func main() {
	info, err := readIdentity()
	if err != nil {
		log.Fatalf("read hostname: %v", err)
	}

	server := &http.Server{
		Addr:              ":8080",
		Handler:           newHandler(info),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("%s listening on %s", info.AppName, server.Addr)
	log.Fatal(server.ListenAndServe())
}

func readIdentity() (identity, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return identity{}, err
	}

	appName, _, _ := strings.Cut(hostname, "-")
	return identity{AppName: appName, Hostname: hostname, PodID: os.Getenv("POD_UID")}, nil
}

func newHandler(info identity) http.Handler {
	page := template.Must(template.ParseFS(frontend.Assets, "index.html"))
	mux := http.NewServeMux()
	mux.Handle("/style.css", http.FileServer(http.FS(frontend.Assets)))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			w.Header().Set("Allow", "GET, HEAD")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", "no-store")
		if err := page.Execute(w, info); err != nil {
			log.Printf("render page: %v", err)
		}
	})

	return mux
}
