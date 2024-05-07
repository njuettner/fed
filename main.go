package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	url := "https://api.stlouisfed.org/fred/series/observations"
	apikey := os.Getenv("FED_API_KEY")

	var seriesId = map[string]string{
		"m1":                 "M1SL",
		"m2":                 "M2SL",
		"debt":               "GFDEGDQ188S",
		"interest_payments":  "A091RC1Q027SBEA",
		"totalassets":        "RESPPANWW",
		"cpi":                "CPALTT01USM657N",
		"cpi_yoy":            "CPALTT01USM659N",
		"cpi_raw":            "CPIAUCNS",
		"bonds10y":           "IRLTLT01USM156N",
		"unfilled_vacancies": "LMJVTTUVUSM647S",
		"gdp":                "GDP",
		"consumer":           "PCE",
		"unemployment":       "UNRATE",
		"fedfunds":           "FEDFUNDS",
		"treasury":           "DGS1",
		"bitcoin":            "CBBTCUSD",
		"ethereum":           "CBETHUSD",
		"industrial":         "INDPRO",
		"retail":             "RSAFS",
		"rdpi":               "DSPIC96",
		"eu":                 "DEXUSEU",
		"oil":                "DCOILWTICO",
	}

	// Fetch data from Fed St. Louis API
	for k, v := range seriesId {
		params := "?series_id=" + v + "&api_key=" + apikey + "&file_type=json"
		req, err := http.NewRequest("GET", url+params, nil)
		if err != nil {
			fmt.Println(err)
		}
		// Make the API request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		ioutil.WriteFile(fmt.Sprintf("data/%s.json", k), data, 0644)
	}
}
