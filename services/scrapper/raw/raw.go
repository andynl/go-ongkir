package raw

type DataFromDataJSON struct {
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
