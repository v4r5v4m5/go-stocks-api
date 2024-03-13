package models

type Stock struct {
	StockID int64  `json:"stockid"` // stock id
	Name    string `json:"name"`    // stock name
	Price   int64  `json:"price"`   // stock price
	Company string `json:"company"` // stock company
}
