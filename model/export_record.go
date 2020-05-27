package model

import "time"

type ExportRecord struct {
	ID        uint
	Amount    uint
	CreatedAt time.Time
}
