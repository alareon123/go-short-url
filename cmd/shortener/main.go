package main

import (
	"github.com/alareon123/go-short-url.git/internal/app"
	"io"
	"log"
	"net/http"
)

func urlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		shortUrl := r.URL.String()[1:]
		longUrl := app.GetUrlById(shortUrl)

		if longUrl == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusTemporaryRedirect)
		w.Header().Set("Location", longUrl)

	} else {
		reqBodyBytes, _ := io.ReadAll(r.Body)
		shortUrl := app.ShortUrl(string(reqBodyBytes))

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "text/plain")
		_, err := w.Write([]byte(shortUrl))
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
