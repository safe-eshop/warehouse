package products

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RossmannPromotion struct {
	Type        string `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	NavigateURL string `json:"navigateUrl,omitempty"`
}

type RossmannPicture struct {
	Large string `json:"large,omitempty"`
}

type RossmannProduct struct {
	ID             int                 `json:"id,omitempty"`
	Brand          string              `json:"brand,omitempty"`
	Caption        string              `json:"caption,omitempty"`
	Price          float64             `json:"price,omitempty"`
	OldPrice       float64             `json:"oldPrice,omitempty"`
	EanNumber      []string            `json:"eanNumber,omitempty"`
	CmpType        string              `json:"cmpType,omitempty"`
	CmpDescription string              `json:"cmpDescription,omitempty"`
	Availability   string              `json:"availability,omitempty"`
	AverageRating  float32             `json:"averageRating,omitempty"`
	Promotions     []RossmannPromotion `json:"promotions,omitempty"`
	Category       string              `json:"category,omitempty"`
	NavigateURL    string              `json:"navigateUrl,omitempty"`
	Pictures       []RossmannPicture   `json:"pictures,omitempty"`
}

type RossmannProductList struct {
	Products []RossmannProduct `json:"products,omitempty"`
}

type RossmannResult struct {
	Data RossmannProductList `json:"data,omitempty"`
}

func GetRossmannPricesFromApi(q int) (RossmannResult, error) {
	var products RossmannResult
	resp, err := http.Get(fmt.Sprintf("https://www.rossmann.pl/products/api/Products?PageSize=%d", q))
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return products, err
	}

	errr := json.Unmarshal(body, &products)
	if errr != nil {
		return products, errr
	}
	return products, nil
}
