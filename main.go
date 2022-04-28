package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Server listening on port 3000")
	http.ListenAndServe(":3000", http.FileServer(http.Dir(os.Args[1])))
}
