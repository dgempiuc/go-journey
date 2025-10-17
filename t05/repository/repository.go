package repository

import (
	"t05/model"

	"gorm.io/gorm"
)

type WarRepository struct {
	DB *gorm.DB
}

func (wp WarRepository) GetAllWar() []model.War {
	var wars []model.War
	wp.DB.Find(&wars)
	return wars
}

func (wp WarRepository) AddWar(war model.War) {
	wp.DB.Create(&war)
}

func (wp WarRepository) GetWarByName(name string) model.War {
	var war model.War
	wp.DB.Where("name = ?", name).First(&war)
	return war
}
