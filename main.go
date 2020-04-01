package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
)

// This will get called for each HTML element found
func processElement(index int, element *goquery.Selection) {
	// See if the href attribute exists on the element
	fmt.Println("=================================")
	fmt.Println(element.Text())
	href, exists := element.Attr("href")
	if exists {
		fmt.Println(href)
	}
}

func getCoronaData(w http.ResponseWriter, r *http.Request) {

	// Make request
	response, err := http.Get("https://en.wikipedia.org/wiki/2020_coronavirus_pandemic_in_Kerala")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find all links and process them with the function
	// defined earlier
	document.Find("#mw-content-text > div > table.infobox > tbody > tr:nth-child(8) > td").Each(processElement)

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

	router.HandleFunc("/api/get-vehicle-data", getCoronaData).Methods("GET")

	http.ListenAndServe(":8001", router)

}
