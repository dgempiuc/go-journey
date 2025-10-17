package handler

import (
	"net/http"
	"t05/model"
	"t05/service"

	"github.com/gin-gonic/gin"
)

type WarHandler struct {
	WS service.WarService
}

func (wh WarHandler) GetWars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, wh.WS.GetAllWar())
}

func (wh WarHandler) AddWar(c *gin.Context) {
	var newWar model.War
	c.BindJSON(&newWar)
	wh.WS.AddWar(newWar)
	c.IndentedJSON(http.StatusCreated, newWar)
}

func (wh WarHandler) GetWarByName(c *gin.Context) {
	name := c.Param("name")
	c.IndentedJSON(http.StatusOK, wh.WS.GetWarByName(name))
}
