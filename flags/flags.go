package flags

import (
	"fmt"
)

type Color string

const (
	ColorBlack        = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset  Color = "\u001b[0m"
)

func Colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

// Main
// useColor := flag.Bool("color", false, "display colorized output")
// flag.Parse()

// if *useColor {
// 	colorize(ColorBlue, "Hello, DigitalOcean!")
// 	return
// }
// fmt.Println("Hello, DigitalOcean!")
