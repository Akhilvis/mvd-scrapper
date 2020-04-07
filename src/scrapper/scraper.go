package scrapdata

import (
	"log"
	"net/http"
	"strings"

	greet "../tagselectors"
	"github.com/PuerkitoBio/goquery"
	"github.com/bitly/go-simplejson"
)

func CoronaData(w http.ResponseWriter, r *http.Request) {

	json := simplejson.New()

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

	Selectorslist := greet.Selectorsmap

	confirmedCasesFinal := strings.Split(document.Find(Selectorslist["confirmedCases"]).Text(), "[")[0]
	activeCases := strings.Split(document.Find(Selectorslist["activeCases"]).Text(), "[")[0]
	recovered := strings.Split(document.Find(Selectorslist["recovered"]).Text(), "[")[0]
	death := strings.Split(document.Find(Selectorslist["death"]).Text(), "[")[0]

	json.Set("confirmedCases", confirmedCasesFinal)
	json.Set("activeCases", activeCases)
	json.Set("recovered", recovered)
	json.Set("death", death)

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

// This will get called for each HTML element found
// func processElement(index int, element *goquery.Selection) string {
// 	// See if the href attribute exists on the element
// 	fmt.Println("=================================")
// 	fmt.Println(element.Text())
// 	return element.Text()
// }
