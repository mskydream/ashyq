package main

import (
	"fmt"

	"github.com/mskydream/ashyq/api/config"
	"github.com/mskydream/ashyq/api/db"
	"github.com/mskydream/ashyq/api/handler"
	"github.com/mskydream/ashyq/api/repository"
	"github.com/mskydream/ashyq/api/service"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Не удается загрузить config:", err)
	}
	db := new(db.DB).InitDatabase(&config.Database)
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	router := handlers.SetupRouter()

	fmt.Printf("Сервер, работающий на порту:%v\n", config.Port)
	router.Run()
}
