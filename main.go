package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Product struct {
	ProductId      int     `json:"productId"`
	Manufacturer   string  `json:"manufacturer"`
	Sku            string  `json:"sku"`
	Upc            string  `json:"upc"`
	PricePerUnit   float64 `json:"pricePerUnit"`
	QuantityOnHand int     `json:"quantityOnHand"`
	ProductName    string  `json:"productName"`
}

var productsList []Product

func init() {
	productsJson := ` [
{
    "productId": 1,
    "manufacturer": "Johns-Jenkins",
    "sku": "p5z343vdS",
    "upc": "939581000000",
    "pricePerUnit": 497.45,
    "quantityOnHand": 9703,
    "productName": "sticky note"
  },
  {
    "productId": 2,
    "manufacturer": "Hessel, Schimmel and Feeney",
    "sku": "i7v300kmx",
    "upc": "740979000000",
    "pricePerUnit": 282.29,
    "quantityOnHand": 9217,
    "productName": "leg warmers"
  },
  {
    "productId": 3,
    "manufacturer": "Swaniawski, Bartoletti and Bruen",
    "sku": "q0L657ys7",
    "upc": "111730000000",
    "pricePerUnit": 436.26,
    "quantityOnHand": 5905,
    "productName": "lamp shade"
  }
]
`
	err := json.Unmarshal([]byte(productsJson), &productsList)
	if err != nil {
		log.Fatal(err)
	}
}

func getNextId() int {
	highestId := -1
	for _, product := range productsList {
		if highestId < product.ProductId {
			highestId = product.ProductId
		}
	}

	return highestId + 1

}

func findProductById(productId int) (*Product, int) {
	for i, product := range productsList {
		if product.ProductId == productId {
			return &productsList[i], i
		}
	}

	return nil, -1
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "products/")
	productId, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	product, listItemIndex := findProductById(productId)
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		productJson, err := json.Marshal(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productJson)

	case http.MethodPut:
		var updatedProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(bodyBytes, &updatedProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if updatedProduct.ProductId != productId {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		product = &updatedProduct
		productsList[listItemIndex] = updatedProduct
		w.WriteHeader(http.StatusOK)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)

	}

}

func productListHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productsJson, err := json.Marshal(productsList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(productsJson)

	case http.MethodPost:
		var newProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(bodyBytes, &newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if newProduct.ProductId != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newProduct.ProductId = getNextId()
		productsList = append(productsList, newProduct)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func timeConsumptionMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware start")
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Println("middleware end in %s", time.Since(start))
	})
}

func main() {
	pListHandler := http.HandlerFunc(productListHandler)
	pHandler := http.HandlerFunc(productHandler)

	http.Handle("/products", timeConsumptionMiddleware(pListHandler))
	http.Handle("/products/", timeConsumptionMiddleware(pHandler))
	http.ListenAndServe(":8080", nil)
}
