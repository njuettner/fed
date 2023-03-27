package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/njuettner/fed-data-collection/pkg/fed"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/data/{series_id}", GetData).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", r))
}

func GetData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fileName := params["series_id"]

	data, err := os.ReadFile(fmt.Sprintf("data/%s.json", fileName))
	if err != nil {
		panic(err)
	}

	fedData := fed.Data{}
	err = json.Unmarshal(data, &fedData)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fedData)
}
