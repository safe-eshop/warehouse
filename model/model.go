package model

type Product struct {
	ID int `json:"id,omitempty"`
	Brand string `json:"brand,omitempty"`
	Caption string `json:"caption,omitempty"`
}

type Price struct {
	ID int `json:"id,omitempty"`
	Product *Product
	Value float64 `json:"value,omitempty"`
	Date Ti
}

type Promotion struct {
	ID int `json:"id,omitempty"`
	Product *Price
	Value float64 `json:"value,omitempty"`
}