package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	u, err := url.Parse("https://127.0.0.1:8443")
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}

	reverseProxy := httputil.NewSingleHostReverseProxy(u)
	reverseProxy.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	server := &http.Server{
		Addr:    ":8081",
		Handler: reverseProxy,
	}

	fmt.Println(server.ListenAndServe())
}
