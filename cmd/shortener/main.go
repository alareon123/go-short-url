package main

import (
	"github.com/alareon123/go-short-url.git/internal/app"
	"io"
	"log"
	"net/http"
)

func urlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		shortURL := r.URL.String()[1:]
		longURL := app.GetUrlByID(shortURL)

		if longURL == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Location", longURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return

	} else {
		reqBodyBytes, _ := io.ReadAll(r.Body)
		shortURL := app.ShortURL(string(reqBodyBytes))

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "text/plain")
		_, err := w.Write([]byte(shortURL))
		if err != nil {
			log.Fatal("error while writing response")
		}
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", urlHandler)

	http.ListenAndServe(":8080", mux)
}
