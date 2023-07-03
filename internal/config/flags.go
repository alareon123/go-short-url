package config

import "flag"

var AppServerURL string
var BaseAddressURL string

func Init() {
	flag.StringVar(&AppServerURL, "a", ":8080", "address and port to run server")
	flag.StringVar(&BaseAddressURL, "b", "http://localhost:8080/", "base address for short urls")
	flag.Parse()
}
