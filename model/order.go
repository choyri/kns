package model

import "time"

type Order struct {
	ID                    uint
	ImportRecordID        uint
	CustomerName          string
	Salesman              string
	CustomerOrderNumber   string
	Brand                 string
	OrderNumber           string
	SerialNumber          uint64
	ProductNameCode       string
	ProductNameChinese    string
	ProductNameEnglish    string
	Ingredient            string
	Specification         string
	Color                 string
	ColorNumber           string
	CustomerVersionNumber string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
