package repository

import (
	"t06/model"
	"time"
)

type WarRepository struct{}

func NewWarRepository() WarRepository {
	return WarRepository{}
}

func (wp WarRepository) GetAllWar() []model.War {
	return InMemoryWarData
}

func (wp WarRepository) AddWar(war model.War) {
	InMemoryWarData = append(InMemoryWarData, war)
}

func (wp WarRepository) GetWarByName(name string) model.War {
	for _, war := range InMemoryWarData {
		if war.Name == name {
			return war
		}
	}
	return model.War{}
}

var InMemoryWarData = []model.War{
	{Name: "Miryokefalon", DateBegin: time.Date(1176, time.September, 17, 0, 0, 0, 0, time.UTC), Duration: 1},
	{Name: "Yassıçemen", DateBegin: time.Date(1230, time.August, 10, 0, 0, 0, 0, time.UTC), Duration: 2},
}
