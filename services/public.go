package scrapper

func getOngkir() (*OngkirData, error) {
	OngkirJSON, err := requestOngkir()
	if err != nil {
		return nil, err
	}

	ongkirData := new(OngkirData)
	fromDataJSON(ongkirData, OngkirJSON)

	return ongkirData, nil

}
