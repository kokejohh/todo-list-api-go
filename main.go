package main

import (
	"net/http"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	});

	log.Println("Server starting on port " + port)

	err := http.ListenAndServe(":" + port, nil);
	if err != nil {
		log.Fatal(err)
	}
}
