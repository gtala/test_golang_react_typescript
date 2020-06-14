package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/isalive", isAlive)

	spa := SpaHandler{StaticPath: "FRONTEND/build", IndexPath: "index.html"}
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		spa.ServeHTTP(w, r)
	})
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("FRONTEND/build"))))

	srv := &http.Server{
		Handler:      RequestLogger(router),
		Addr:         "127.0.0.1:5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println(srv.ListenAndServe())
}

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		targetMux.ServeHTTP(w, r)
		requesterIP := r.RemoteAddr
		println(
			"%s\t\t%s\t\t%s\t\t%v",
			r.Method,
			r.RequestURI,
			requesterIP,
			time.Since(start),
		)
	})
}

func isAlive(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ALIVE")
}

type SpaHandler struct {
	StaticPath string
	IndexPath  string
}

func (h SpaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(h.StaticPath, path)

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.StaticPath, h.IndexPath))
		return
	} else if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.StaticPath)).ServeHTTP(w, r)
}
