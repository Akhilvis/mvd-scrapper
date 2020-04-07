package greet

// func GetSelectorsList() map[string]string {
// 	var Selectorsmap = make(map[string]string)
// 	Selectorsmap["coxxnfirmedCases"] = "dsj"
// 	return Selectorsmap
// }

var Selectorsmap = make(map[string]string)

func init() {
	Selectorsmap["confirmedCases"] = "#mw-content-text > div > table.infobox > tbody > tr:nth-child(8) > td"
	Selectorsmap["activeCases"] = "#mw-content-text > div > table.infobox > tbody > tr:nth-child(9) > td"
	Selectorsmap["recovered"] = "#mw-content-text > div > table.infobox > tbody > tr:nth-child(10) > td"
	Selectorsmap["death"] = "#mw-content-text > div > table.infobox > tbody > tr:nth-child(11) > td"

}
