package model

// License describes the Xpack license state.
type License struct {
	IsValid     bool   `json:"is_valid"`
	Edition     string `json:"edition"`
	ExpiredDate string `json:"expired_date"`
	Count       int    `json:"count"`
	Corporation string `json:"corporation"`
}
