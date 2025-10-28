package models



type CurrenciesType struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Countries struct {
	Name       string           `json:"name"`
	Capital    string           `json:"capital"`
	Region     string           `json:"region"`
	Population int              `json:"population"`
	Currencies []CurrenciesType `json:"currencies"`
	Flag       string           `json:"flag"`
}

type ExchangeRate struct {
	BaseCode string             `json:"base_code"`
	Rates    map[string]float64 `json:"rates"`
}

type DBData struct {
	Id                int       `db:"id" json:"id,omitempty"`
	Name              string    `db:"name" json:"name"`
	Capital           string    `db:"capital" json:"capital"`
	Region            string    `db:"region" json:"region"`
	Population        int       `db:"population" json:"population"`
	Currency_code     *string   `db:"currency_code" json:"currency_code,omitempty"`
	Exchange_rate     *float64   `db:"exchange_rate" json:"exchange_rate,omitempty"`
	Estimated_gdp     *float64   `db:"estimated_gdp" json:"estimated_gdp,omitempty"`
	Flag_url          string    `db:"flag_url" json:"flag_url"`
	Last_refreshed_at  string  `db:"last_refreshed_at" json:"last_refreshed_at,omitempty"`
}



type TopGDP struct {
	Name string `db:"name"`
	EstimatedGDP float64 `db:"estimated_gdp"`
}
