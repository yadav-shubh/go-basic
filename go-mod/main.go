package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")
	router := mux.NewRouter()
	router.HandleFunc("/", GreetGoodMorning).Methods(http.MethodGet)

	// create server and server by router
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}

func GreetGoodMorning(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("<h1>Hello World</h1>"))
	checkNullError(err)
}
