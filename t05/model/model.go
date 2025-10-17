package model

import "time"

type War struct {
	Name      string    `json:"war-name" gorm:"primary_key"`
	DateBegin time.Time `json:"begin-date" gorm:"not null"`
	Duration  int       `json:"total-day" gorm:"not null"`
}
