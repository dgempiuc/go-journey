package main

import (
	"fmt"
	"t05/config"
	"t05/handler"
	"t05/model"
	"t05/repository"
	"t05/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	var cfg config.DBConfig = config.LoadDBConfig()
	db, err := config.DatabaseConnection(cfg)
	if err != nil {
		fmt.Println("db connection sırasında hata olustu. %v", err)
	}
	err = db.AutoMigrate(model.War{})
	if err != nil {
		fmt.Println("tabloları otomatik olusturma sırasında hata olustu. %v", err)
	}
	var warRepo repository.WarRepository = repository.WarRepository{db}
	var warService service.WarService = service.WarService{warRepo}
	var warHandler handler.WarHandler = handler.WarHandler{warService}

	initDBData(warService)

	router := gin.Default()
	router.GET("/wars", warHandler.GetWars)
	router.POST("/wars", warHandler.AddWar)
	router.GET("/wars/:name", warHandler.GetWarByName)

	router.Run("localhost:8080")
}

var InMemoryWarData = []model.War{
	{Name: "Miryokefalon", DateBegin: time.Date(1176, time.September, 17, 0, 0, 0, 0, time.UTC), Duration: 1},
	{Name: "Yassıçemen", DateBegin: time.Date(1230, time.August, 10, 0, 0, 0, 0, time.UTC), Duration: 2},
}

func initDBData(ws service.WarService) {
	fmt.Println(cap(ws.GetAllWar()))
	if len(ws.GetAllWar()) == 0 {
		for _, data := range InMemoryWarData {
			ws.AddWar(data)
		}
	}
}
