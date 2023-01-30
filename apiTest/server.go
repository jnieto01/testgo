package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func main() {
	host := "localhost:80"
	router := mux.NewRouter()

	router.HandleFunc("/test", myTest)

	// -------------- start server ----------------------------------
	log.Fatal(http.ListenAndServe(host, router))
}

func myTest(w http.ResponseWriter, r *http.Request){
	fmt.Println("paso") 
}
