package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	message   = flag.String("message", "This service is currently down for maintenance.", "Message to show on the landing page.")
	port      = flag.String("port", "8080", "Port to run on.")
	httpsport = flag.String("httpsport", "443", "Https port (requires cert and key)")
	certfile  = flag.String("cert", "", "Path to cert (if https)")
	keyfile   = flag.String("key", "", "Path to key (if https)")
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h2><center>Service is Down</center></h2></br><center>%s</center>", *message)
}

func hostHTTP() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), nil))
}

func hostHTTPS() {
	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%s", *httpsport), *certfile, *keyfile, nil))
}

func main() {
	flag.Parse()

	http.HandleFunc("/", handler)
	if *certfile != "" && *keyfile != "" {
		go hostHTTPS()
	}
	hostHTTP()
}
