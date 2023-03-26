package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/njuettner/fed-data-collection/pkg/moneysupply"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/supply/{money}", GetMoneySupply).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", r))
}

func GetMoneySupply(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fileName := params["money"]
	fmt.Println(fileName)
	data, err := os.ReadFile(fmt.Sprintf("server/%s.json", fileName))
	if err != nil {
		panic(err)
	}
	money := moneysupply.MoneySupply{}
	err = json.Unmarshal(data, &money)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(money)
}
