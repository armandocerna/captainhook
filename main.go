package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	proxyPort := flag.String("proxy-port", "9090", "Port that capitanhook proxy binds to")
	targetURL := flag.String("target-host", "google.com", "Target URL to proxy to")
	targetPort := flag.String("target-port", "80", "Target port to proxy to")
	flag.Parse()
	targetHost := fmt.Sprintf("%s:%s", targetURL, targetPort)
	listenString := fmt.Sprintf(":%f", proxyPort)
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   targetHost,
	})
	fmt.Printf("The capitan is listening %v", listenString)
	http.ListenAndServe(listenString, proxy)
}
