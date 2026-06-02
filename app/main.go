package main

import (
	f "fmt"
	"log"
	"net/http"
)

func main() {

	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileserver)

	f.Printf("servidor rodando em > http://localhost:8081/\n")

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
