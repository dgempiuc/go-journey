package service

import (
	"t05/model"
	"t05/repository"
)

type WarService struct {
	WP repository.WarRepository
}

func (ws WarService) GetAllWar() []model.War {
	return ws.WP.GetAllWar()
}

func (ws WarService) AddWar(war model.War) {
	ws.WP.AddWar(war)
}

func (ws WarService) GetWarByName(name string) model.War {
	return ws.WP.GetWarByName(name)
}
