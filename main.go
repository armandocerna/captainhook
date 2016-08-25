package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	targetHost := fmt.Sprintf("%s:%s", os.Args[2], os.Args[3])
	proxyPort := os.Args[1]
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   targetHost,
	})
	fmt.Printf("The capitan is listening")
	http.ListenAndServe(":"+proxyPort, proxy)
}
