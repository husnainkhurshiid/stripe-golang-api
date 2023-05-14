package model

type Payment struct {
	Name       string      `json:"name"`
	Email      string      `json:"email"`
	Amount     int64       `json:"amount"`
	Currency   string      `json:"currency"`
	Desc       string      `json:"desc"`
	Line1      string      `json:"line1"`
	Line2      string      `json:"line2"`
	City       string      `json:"city"`
	State      string      `json:"state"`
	PostalCode string      `json:"postal_code"`
	Country    string      `json:"country"`
	Card       CardDetails `json:"card"`
}

type CardDetails struct {
	Number   string `json:"number"`
	ExpMonth string `json:"exp_month"`
	ExpYear  string `json:"exp_year"`
	CVC      string `json:"cvc"`
}
