package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("dir: %s\n", currentDir)
	// todo make this path relative to main.go file?
	buildDir := path.Join(currentDir, "frontend/build")

	staticFileHandler := http.FileServer(http.Dir(buildDir))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("path: %s\n", r.URL.Path)
		if r.URL.Path == "/" {
			http.ServeFile(w, r, path.Join(buildDir, "index.html"))
		} else {
			staticFileHandler.ServeHTTP(w, r)
		}
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("listening on :%s\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
