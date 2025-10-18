package model

import (
	"time"
)

type War struct {
	Name      string    `json:"war-name" gorm:"primary_key"`
	DateBegin time.Time `json:"begin-date" gorm:"not null"`
	Duration  int       `json:"total-day" gorm:"not null"`
}

func (w War) getAllWars() []War {
	return InMemoryWarData
}

var InMemoryWarData = []War{
	{Name: "Miryokefalon", DateBegin: time.Date(1176, time.September, 17, 0, 0, 0, 0, time.UTC), Duration: 1},
	{Name: "Yassıçemen", DateBegin: time.Date(1230, time.August, 10, 0, 0, 0, 0, time.UTC), Duration: 2},
}
