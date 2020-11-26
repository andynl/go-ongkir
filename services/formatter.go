package scrapper

import (
	rawStruct "github.com/andynl/go-ongkir/services/scrapper/raw"
)

type OngkirData struct {
	ProvinceID string `json:"provinceId"`
	Province   string `json:"provinceName"`
	CityID     string `json:"cityId"`
	CityName   string `json:"cityName"`
}

func fromDataJSON(result *OngkirData, raw *rawStruct.DataFromDataJSON) {
	result.ProvinceID = raw.Rajaongkir.Results.ProvinceID
	result.Province = raw.Rajaongkir.Results.Province
	result.CityID = raw.Rajaongkir.Results.CityID
	result.CityName = raw.Rajaongkir.Results.CityName
}
