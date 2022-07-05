package main

import (
	"fmt"
	"image/color"

	"github.com/skip2/go-qrcode"
)

func main() {
	// first example
	// err := qrcode.WriteFile("hello this is qrcode in golang", qrcode.Medium, 256, "myfirst_file.png")
	// if err != nil {
	// 	fmt.Printf("Sorry couldn't create qrcode: %v\n", err)
	// }

	// second example
	err := qrcode.WriteColorFile("this is colored", qrcode.Medium, 256, color.Transparent, color.White, "exparimant.jpg")
	if err != nil {
		fmt.Printf("Sorry couldn't create qrcode: %v\n", err)
	}
}
