package app

var urls = make(map[string]string)

func ShortUrl(url string) string {
	randString := RandStringBytes(5)
	storeUrl(url, randString)
	baseUrl := "http://localhost:8080/"
	return baseUrl + randString
}

func GetUrlById(shortUrl string) string {
	return getUrl(shortUrl)
}

func storeUrl(urlBase string, urlShort string) {
	urls[urlShort] = urlBase
}

func getUrl(urlShort string) string {
	return urls[urlShort]
}
