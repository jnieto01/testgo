package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func main() {
	host := "192.168.0.55:80"
	router := mux.NewRouter()

	router.HandleFunc("/test", myTest)

	// -------------- start server ----------------------------------
	log.Fatal(http.ListenAndServe(host, router))
}

func myTest(w http.ResponseWriter, r *http.Request){
	fmt.Println("paso") 
}
