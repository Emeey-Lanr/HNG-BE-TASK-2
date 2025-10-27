package models

type CurrenciesType struct {
	Code string `json:"code"`
	Name string  `json:"name"`
}

type Countries struct {
   Name string `json:"name"`
   Capital string `json:"capital"`
   Region string `json:"region"`
   Population int `json:"population"`
   Currencies []CurrenciesType `json:"currencies"`
   Flag string `json:"flag"`

}


type ExchangeRate struct {
	BaseCode string `json:"base_code"`
	Rates map[string]float64 `json:"rates"`

}



