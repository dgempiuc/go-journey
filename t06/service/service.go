package service

import (
	"t06/model"
	"t06/repository"
)

type WarService struct {
	wp repository.WarRepository
}

func NewWarService(wp repository.WarRepository) WarService {
	return WarService{wp}
}

func (ws WarService) GetAllWar() []model.War {
	return ws.wp.GetAllWar()
}

func (ws WarService) AddWar(war model.War) {
	ws.wp.AddWar(war)
}

func (ws WarService) GetWarByName(name string) model.War {
	return ws.wp.GetWarByName(name)
}
