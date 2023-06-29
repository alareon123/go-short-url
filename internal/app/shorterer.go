package app

import "fmt"

var urls = make(map[string]string)

func ShortURL(url string, host string) string {
	randString := RandStringBytes(8)
	storeURL(url, randString)
	baseURL := "http://localhost:8080/"
	return baseURL + randString
}

func GetURLByID(shortURL string) string {
	return getURL(shortURL)
}

func storeURL(urlBase string, urlShort string) {
	fmt.Printf("url %s stored with id %s", urlBase, urlShort)
	urls[urlShort] = urlBase
	fmt.Println(urls)
}

func getURL(urlShort string) string {
	return urls[urlShort]
}
