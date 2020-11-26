package scrapper

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	rawStruct "github.com/andynl/go-ongkir/services/scrapper/raw"
)

const (
	apiKey = "ed6c2bcfba37c2fe6cdc51bd1a6170a0"
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

func requestOngkir() (*rawStruct.DataFromDataJSON, error) {
	raw, err := getApi()
	if err != nil {
		return nil, err
	}

	result := new(rawStruct.DataFromDataJSON)
	if err = json.Unmarshal(raw, result); err != nil {
		return nil, err
	}

	return result, nil
}
