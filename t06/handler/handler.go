package handler

import (
	"net/http"
	"t06/model"
	"t06/service"

	"github.com/gin-gonic/gin"
)

type WarHandler struct {
	ws service.WarService
}

func NewWarHandler(ws service.WarService) WarHandler {
	return WarHandler{ws}
}

func (wh WarHandler) GetWars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, wh.ws.GetAllWar())
}

func (wh WarHandler) AddWar(c *gin.Context) {
	var newWar model.War
	c.BindJSON(&newWar)
	wh.ws.AddWar(newWar)
	c.IndentedJSON(http.StatusCreated, newWar)
}

func (wh WarHandler) GetWarByName(c *gin.Context) {
	name := c.Param("name")
	c.IndentedJSON(http.StatusOK, wh.ws.GetWarByName(name))
}
