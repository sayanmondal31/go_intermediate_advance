package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling incoming orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling incoming users")
	})
	PORT := 3000

	// load the TLS cert and key
	cert := "cert.pem"
	key := "key.pem"

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// create a custom server
	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", PORT),
		Handler:   nil,
		TLSConfig: tlsConfig,
	}

	// enable http2
	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("server is running on port:", PORT)
	if err := server.ListenAndServeTLS(cert, key); err != nil {
		log.Fatal("Couldn't start server", err)
	}

	// http 1.1 server
	// if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil); err != nil {
	// 	log.Fatal("Couldn't start server", err)
	// }

}
