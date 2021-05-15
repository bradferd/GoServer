package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const Port = ":5500"

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", rootPage)
	router.HandleFunc("/products/{fetchCountPercentage}", products).Methods("GET")

	fmt.Println("Serving @ http://127.0.0.1:" + Port)
	log.Fatal(http.ListenAndServe(Port, router))

}

func rootPage(W http.ResponseWriter, r *http.Request) {

	W.Write([]byte("This is the root page"))
}

func products(w http.ResponseWriter, r *http.Request) {

	fetchCountPercentage, errInput := strconv.ParseFloat(mux.Vars(r)["fetchCountPercentage"], 64)

	fetchCount := 0

	if errInput != nil {
		fmt.Println(errInput.Error())
	} else {
		fetchCount = int(float64(len(productList)) * fetchCountPercentage / 100)

		if fetchCount > len(productList) {
			fetchCount = len(productList)
		}
	}

	// write to response
	jsonList, err := json.Marshal(productList[0:fetchCount])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	} else {
		w.Header().Set("content-type", "application/json")
		w.Write(jsonList)
	}
}

type product struct {
	Name  string
	Price float64
	Count int
}

var productList = []product{

	product{"p1", 25.0, 30},
	product{"p2", 20.0, 10},
	product{"p3", 250.0, 320},
	product{"p4", 22.0, 50},
	product{"p5", 23.0, 30},
	product{"p6", 258.0, 30},
	product{"p7", 2.0, 3},
	product{"p8", 5.0, 30},
	product{"p9", 205.0, 50},
	product{"p10", 225.0, 40},
	product{"p11", 253.0, 88},
	product{"p12", 54.0, 78},
	product{"p13", 29.0, 67},
}
