package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	HttpPort := flag.Int("p", 8000, "Port to bind to")
	HttpDir := flag.String("d", ".", "Directory to share")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*HttpDir)))
	fmt.Printf("Sharing %v on port %v\n", *HttpDir, *HttpPort)
	err := http.ListenAndServe(fmt.Sprintf(":%v", *HttpPort), logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
