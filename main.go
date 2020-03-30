package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
)

func getVehicleData(w http.ResponseWriter, r *http.Request) {
	json := simplejson.New()
	json.Set("foo", "bar")

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
func main() {
	fmt.Println(".............")
	router := mux.NewRouter()

	router.HandleFunc("/api/get-vehicle-data", getVehicleData).Methods("GET")

	http.ListenAndServe(":8001", router)

}
