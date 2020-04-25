package model

import "time"

type ImportRecord struct {
	ID        uint
	StartDate string
	EndDate   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
