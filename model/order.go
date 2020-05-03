package model

import (
	"reflect"
	"strings"
	"time"
)

type Order struct {
	ID                    uint      `json:"id"`
	ImportRecordID        uint      `json:"-"`
	CustomerName          string    `json:"customer_name"`
	Salesman              string    `json:"-"`
	CustomerOrderNumber   string    `json:"customer_order_number"`
	Brand                 string    `json:"brand"`
	OrderNumber           string    `json:"order_number"`
	SerialNumber          uint64    `json:"serial_number"`
	ProductNameCode       string    `json:"product_name_code"`
	ProductNameChinese    string    `json:"product_name_chinese"`
	ProductNameEnglish    string    `json:"product_name_english"`
	Ingredient            string    `json:"ingredient"`
	Specification         string    `json:"specification"`
	Color                 string    `json:"color"`
	ColorNumber           string    `json:"color_number"`
	CustomerVersionNumber string    `json:"customer_version_number"`
	CreatedAt             time.Time `json:"-"`
	UpdatedAt             time.Time `json:"-"`

	Keyword string `json:"-" gorm:"-"`
}

func (o *Order) TrimSpace() {
	refValue := reflect.ValueOf(o).Elem()

	for i, j := 0, refValue.NumField(); i < j; i++ {
		v := refValue.Field(i)
		if v.Type().Kind() != reflect.String {
			continue
		}
		v.SetString(strings.TrimSpace(v.String()))
	}
}

type Orders []Order

func (o Orders) Len() int {
	return len(o)
}

func (o Orders) Less(i, j int) bool {
	return o[i].ID > o[j].ID
}

func (o Orders) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}
