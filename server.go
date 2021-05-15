package main

import (

	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

const Port = ":5500"

func main(){

	router:= mux.NewRouter()
	router.HandleFunc("/", rootPage)

	fmt.Println("Serving @ http://127.0.0.1:" + Port)
	log.Fatal(http.ListenAndServe(Port, router))


}

func rootPage(W http.ResponseWriter, r *http.Request) {

	W.Write([]byte("This is the root page"))	
}