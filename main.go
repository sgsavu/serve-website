package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "8080", "Port to start the server on")
	flag.Parse()
	path := flag.Arg(0)

	if path == "" {
		log.Fatal("directory needs to be provided")
	}

	fs := http.FileServer(http.Dir(path))

	http.Handle("/", fs)

	fmt.Printf("Server is running on http://localhost:%s\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", *port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
