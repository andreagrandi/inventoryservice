package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Product is a structure representing a product
type Product struct {
	ProductID      int    `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

var productList []Product

func init() {
	productsJSON := `[
		{
			"productId": 1,
			"manufacturer": "Johns-Jenkins",
			"sku": "g4k52kh54jk",
			"upc": "93724562387465879",
			"pricePerUnit": "497.45",
			"quantityOnHand": 9703,
			"productName": "sticky note"
		},
		{
			"productId": 2,
			"manufacturer": "Antani-Tarapia",
			"sku": "gk54khjgfsja",
			"upc": "57823658792368475",
			"pricePerUnit": "840.12",
			"quantityOnHand": 6345,
			"productName": "Sblinda"
		},
		{
			"productId": 3,
			"manufacturer": "Sassaroli",
			"sku": "kfkjs48fjs",
			"upc": "938958957289345",
			"pricePerUnit": "125.68",
			"quantityOnHand": 9894,
			"productName": "Afasol"
		}
	]`

	err := json.Unmarshal([]byte(productsJSON), &productList)
	if err != nil {
		log.Fatal(err)
	}
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productsJSON, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productsJSON)
	}
}

func main() {
	http.HandleFunc("/products", productsHandler)
	http.ListenAndServe(":8000", nil)
}
