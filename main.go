package main

import (
	"log"
	"net/http"
)

// "Hello from Snippetbox" と返却する home ハンドラを定義する
func home(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Hello from Snippetbox")); err != nil {
		log.Fatalf("error occurred in home handler: %w", err)
	}
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Display a specific snippet...")); err != nil {
		log.Fatalf("error occurred in showSnippet handler: %w", err)
	}
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Create a new snippet...")); err != nil {
		log.Fatalf("error occurred in createSnippet handler: %w", err)
	}
}

func main() {
	// 新しい servemux を作成して、/ に対して home ハンドラを割り当てます
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")
	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Fatalf("cannot start servemux: %w", err)
	}
}