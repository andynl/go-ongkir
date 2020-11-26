package controller

import (
	"encoding/json"
	"net/http"

	scrapper "github.com/andynl/go-ongkir/services/scrapper"
)

type Data struct {
	Errors []Error `json:"errors"`
	*scrapper.OngkirData
}
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func GetOngkir(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data Data
	ongkirData, err := scrapper.GetOngkir()
	if err != nil {
		json.NewEncoder(w).Encode(Error{
			Code:    500,
			Message: err.Error(),
		})

		return
	}

	data.OngkirData = ongkirData
	json.NewEncoder(w).Encode(data)
}
