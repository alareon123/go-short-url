package main

import (
	"github.com/alareon123/go-short-url.git/internal/app"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
)

func urlShortHandler(w http.ResponseWriter, r *http.Request) {
	reqBodyBytes, _ := io.ReadAll(r.Body)

	if len(reqBodyBytes) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	shortURL := app.ShortURL(string(reqBodyBytes), r.Host)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte(shortURL))
	if err != nil {
		log.Fatal("error while writing response")
	}
}

func getUrlHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	longURL := app.GetURLByID(id)

	if longURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", longURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func main() {
	r := chi.NewRouter()

	r.Post("/", urlShortHandler)
	r.Get("/{id}", getUrlHandler)

	http.ListenAndServe(":8080", r)
}
