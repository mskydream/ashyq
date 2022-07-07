package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/skip2/go-qrcode"
)

func main() {
	// first example
	// err := qrcode.WriteFile("hello this is qrcode in golang", qrcode.Medium, 256, "myfirst_file.png")
	// if err != nil {
	// 	fmt.Printf("Sorry couldn't create qrcode: %v\n", err)
	// }

	// second example
	err := qrcode.WriteColorFile("this is colored", qrcode.High, 256, color.Transparent, color.RGBA{0, 0, 255, 1}, "experiment.png")
	if err != nil {
		fmt.Printf("Sorry couldn't create qrcode: %v\n", err)
	}
	fmt.Println(math.Pow(7089, 10))
}
