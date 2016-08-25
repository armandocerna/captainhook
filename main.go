package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	var (
		proxyPort  string
		targetURL  string
		targetPort string
	)

	flag.StringVar(&proxyPort, "proxy-port", "9090", "Port that capitanhook proxy binds to")
	flag.StringVar(&targetURL, "target-url", "google.com", "Target URL to proxy to")
	flag.StringVar(&targetPort, "target-port", "80", "Target port to proxy to")
	flag.Parse()
	targetHost := fmt.Sprintf("%s:%s", targetURL, targetPort)
	listenString := fmt.Sprintf(":%v", proxyPort)

	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   targetHost,
	})
	fmt.Printf("The capitan is listening on port: %s, proxying traffic to %s", proxyPort, targetHost)
	http.ListenAndServe(listenString, proxy)
}
