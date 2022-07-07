package main

import (
	"fmt"

	"github.com/mskydream/qr-code/src/config"
)

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Не удается загрузить config:", err)
	} else {
		fmt.Println("Config working!")
	}

}
