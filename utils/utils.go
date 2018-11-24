package utils

import (
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
)

func Colorize(text string, color string) string {
	c, err := colorful.Hex(color)

	if err != nil {
		fmt.Println(err)
		fmt.Printf(`Color "%s" does not exist. Using #FFFFFF`, color)
		fmt.Println("Fix that!")
		return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", 255, 255, 255, text)
	}
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[38;2;0m", int(c.R * 255), int(c.G * 255), int(c.B * 255), text)
}
