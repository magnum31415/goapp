package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/goji/httpauth"
)

func main() {

	var hostandport string
	var hostArg, portArg, crtArg, keyArg string

	flag.StringVar(&hostArg, "host", "localhost", "a string")
	flag.StringVar(&portArg, "port", "8443", "a string")
	flag.StringVar(&crtArg, "crt", "server.crt", "a string")
	flag.StringVar(&keyArg, "key", "server.key", "a string")
	flag.Parse()

	hostandport = hostArg + ":" + portArg

	http.Handle("/", httpauth.SimpleBasicAuth("userguest", "passwdguest")(http.FileServer(http.Dir("./public_html"))))
	//http.ListenAndServe(":8080", nil)
	//openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 10000 -nodes -subj '/CN=localhost'
	log.Fatal(http.ListenAndServeTLS(hostandport, crtArg, keyArg, nil))

}
