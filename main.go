package main

import (
	"fmt"

	"github.com/mskydream/qr-code/cmd/config"
	"github.com/mskydream/qr-code/cmd/db"
	"github.com/mskydream/qr-code/cmd/handler"
	"github.com/mskydream/qr-code/cmd/repository"
	"github.com/mskydream/qr-code/cmd/service"
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
