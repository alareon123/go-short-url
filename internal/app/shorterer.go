package app

import "fmt"

var urls = make(map[string]string)

func ShortURL(url string, host string) string {
	randString := RandStringBytes(8)
	storeURL(url, randString)
	return host + randString
}

func GetUrlByID(shortURL string) string {
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

func GetAll() {
	for s, s2 := range urls {
		fmt.Println(s)
		fmt.Println(s2)
	}
}
