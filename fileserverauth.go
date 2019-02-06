package main

import(
    "net/http"
    "github.com/goji/httpauth"
)


func main() {
    http.Handle("/", httpauth.SimpleBasicAuth("someuser", "somepassword")(http.FileServer(http.Dir("./public_html"))))
    http.ListenAndServe(":8080", nil)
}


