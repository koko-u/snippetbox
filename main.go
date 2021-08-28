package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// "Hello from Snippetbox" と返却する home ハンドラを定義する
func home(w http.ResponseWriter, r *http.Request) {
	// ルートURL以外のリクエストがきたら 404 Not Found とする
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if _, err := w.Write([]byte("Hello from Snippetbox")); err != nil {
		log.Fatalf("error occurred in home handler: %w", err)
	}
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}
	if _, err := fmt.Fprintf(w, "Display a specific snippet with ID:%d...", id); err != nil {
		http.Error(w, "Error has occurred.", http.StatusInternalServerError)
	}
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Add("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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