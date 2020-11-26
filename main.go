package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AutoGenerated struct {
	Rajaongkir Rajaongkir `json:"rajaongkir"`
}
type Query struct {
	Province string `json:"province"`
	ID       string `json:"id"`
}
type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}
type Results struct {
	CityID     string `json:"city_id"`
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}
type Rajaongkir struct {
	Query   Query   `json:"query"`
	Status  Status  `json:"status"`
	Results Results `json:"results"`
}

const (
	apiKey = "YOUR_API_KEY"
)

func httpHeader() http.Header {
	httpHeader := http.Header{}
	httpHeader.Set("Key", apiKey)

	return httpHeader
}

func getApi() ([]byte, error) {
	req, err := http.NewRequest("GET", "https://pro.rajaongkir.com/api/city?id=39", nil)
	if err != nil {
		return nil, err
	}

	req.Header = httpHeader()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 200 {
		return nil, errors.New(string(body))
	}

	return body, nil
}

type Data struct {
	Errors []Error `json:"errors"`
	*Ongkir
}

type Ongkir struct {
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
	CityID     string `json:"city_id"`
	CityName   string `json:"city"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func requestOngkir() (*AutoGenerated, error) {
	raw, err := getApi()
	if err != nil {
		return nil, err
	}

	result := new(AutoGenerated)
	if err = json.Unmarshal(raw, result); err != nil {
		return nil, err
	}

	return result, nil
}

func fromDataJSON(result *Ongkir, raw *AutoGenerated) {
	result.ProvinceID = raw.Rajaongkir.Results.ProvinceID
	result.Province = raw.Rajaongkir.Results.Province
	result.CityID = raw.Rajaongkir.Results.CityID
	result.CityName = raw.Rajaongkir.Results.CityName
}

func getOngkirService() (*Ongkir, error) {
	OngkirJSON, err := requestOngkir()
	if err != nil {
		return nil, err
	}

	ongkirData := new(Ongkir)
	fromDataJSON(ongkirData, OngkirJSON)

	return ongkirData, nil

}

func getOngkir(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data Data
	ongkirData, err := getOngkirService()
	if err != nil {
		json.NewEncoder(w).Encode(Error{
			Code:    500,
			Message: err.Error(),
		})

		return
	}

	data.Ongkir = ongkirData
	json.NewEncoder(w).Encode(data)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getOngkir", getOngkir).Methods("GET")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":4321", router))
}