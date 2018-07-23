package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	message = flag.String("message", "This service is currently down for maintenance.", "Message to show on the landing page.")
	port    = flag.String("port", "8080", "Port to run on.")
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h2><center>Service is Down</center></h2></br><center>%s</center>", *message)
}

func main() {
	flag.Parse()

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
