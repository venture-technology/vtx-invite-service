package models

type Response struct {
	Payout *Payout `json:"payout"`
}

type Payout struct {
	Driver *string `json:"driver"`
	School *string `json:"school"`
}
