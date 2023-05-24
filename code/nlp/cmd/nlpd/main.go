package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/leo1994/nlp"
)

func main() {
	http.HandleFunc("/tokenize", tokenizeHandler)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}
func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	defer r.Body.Close()
	rdr := io.LimitReader(r.Body, 1_000_000)
	b, err := io.ReadAll(rdr)

	if err != nil {
		log.Fatalln(err)
		http.Error(w, "can't read", http.StatusBadRequest)
		return
	}

	if len(b) == 0 {
		http.Error(w, "missing data", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")

	tokens := nlp.Tokenize(string(b))
	resp := map[string]any{
		"tokens": tokens,
	}
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "can't enconde", http.StatusInternalServerError)
		return
	}

	w.Write(data)

	/* tokens := struct {
		Tokens []string `json:"tokens,omitempty"`
	}{
		Tokens: nlp.Tokenize(string(b)),
	}

	json.NewEncoder(w).Encode(tokens) */
}
