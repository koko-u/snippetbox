package main

import (
	"log"
	"net/http"
)

// "Hello from Snippetbox" と返却する home ハンドラを定義する
func home(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello from Snippetbox"))
	if err != nil {
		log.Fatalf("error occurred in home handler: %w", err)
	}
}

func main() {
	// 新しい servemux を作成して、/ に対して home ハンドラを割り当てます
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on :4000")
	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Fatalf("cannot start servemux: %w", err)
	}
}