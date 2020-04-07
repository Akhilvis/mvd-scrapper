package main

import (
	"fmt"
	"net/http"

	scrapdata "./src/scrapper"
	"github.com/gorilla/mux"
)

func getCoronaData(w http.ResponseWriter, r *http.Request) {
	scrapdata.CoronaData(w, r)
}

func main() {
	fmt.Println(".............")
	router := mux.NewRouter()

	router.HandleFunc("/api/get-vehicle-data", getCoronaData).Methods("GET")

	http.ListenAndServe(":8001", router)

}
