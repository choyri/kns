package model

import "time"

type ImportRecord struct {
	ID        uint
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
