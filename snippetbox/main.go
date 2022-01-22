package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from c3n7"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Cache-Control", "public, max-age=2343242")
	w.Header().Add("Cache-Control", "public")
	w.Header().Add("Cache-Control", "max-age=234324")
	// w.Header().Del("Cache-Control")
	log.Println(w.Header().Get("Cache-Control"))
	w.Write([]byte("Display a specific snippet..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet"))
}

func demoJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header()["X-XSS-Protection"] = []string{"1; mode=block"}
	// w.Header().Del("Date") // Wont work
	w.Header()["Date"] = nil
	w.Write([]byte(`{"name":"Alex"}`))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	mux.HandleFunc("/json", demoJson)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
