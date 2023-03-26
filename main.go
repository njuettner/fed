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
		"m1": "M1SL",
		"m2": "M2SL",
	}

	// Fetch data from Fed St. Louis API
	for k, v := range seriesId {

		params := "?series_id=" + v + "&api_key=" + apikey + "&file_type=json"
		req, err := http.NewRequest("GET", url+params, nil)
		if err != nil {
			panic(err)
		}
		// Make the API request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		ioutil.WriteFile(fmt.Sprintf("server/%s.json", k), data, 0644)
	}
}
