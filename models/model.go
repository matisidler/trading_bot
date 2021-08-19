package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type CreateOrderResponse struct {
	gorm.Model
	Symbol                   string `gorm:"type:varchar(70)"`
	OrderID                  int    `gorm:"type:int;not null"`
	Side                     string `gorm:"type:varchar(70)"`
	Type                     string `gorm:"type:varchar(70)"`
	Price                    string `gorm:"type:varchar(70)"`
	ExecutedQuantity         string `gorm:"type:varchar(70)"`
	CummulativeQuoteQuantity string `gorm:"type:varchar(70)"`
}

type FinnhubResponse struct {
	CurrentPrice       json.Number `json:"c"`
	Change             json.Number `json:"d"`
	PercentChange      json.Number `json:"dp"`
	HighPriceOfDay     json.Number `json:"h"`
	LowPriceOfDay      json.Number `json:"l"`
	OpenPriceOfDay     json.Number `json:"o"`
	PreviousPriceClose json.Number `json:"pc"`
}
