package main

import (
	"fmt"

	"github.com/mskydream/ashyq/cmd/config"
	"github.com/mskydream/ashyq/cmd/db"
	"github.com/mskydream/ashyq/cmd/handler"
	"github.com/mskydream/ashyq/cmd/repository"
	"github.com/mskydream/ashyq/cmd/service"
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
