package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("hello-go is listening")
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello-go again")
	fmt.Println("hello sent")
}
