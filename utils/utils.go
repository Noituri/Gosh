package utils

import (
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"os"
	"strings"
)

func Colorize(text string, color string) string {
	c, err := colorful.Hex(color)

	if err != nil {
		fmt.Println(err)
		fmt.Printf(`Color "%s" does not exist. Using #FFFFFF`, color)
		fmt.Println("\nFix that!")
		return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", 255, 255, 255, text)
	}
	return fmt.Sprintf(`\x1b[38;2;%d;%d;%dm%s\x1b[0m`, int(c.R * 255), int(c.G * 255), int(c.B * 255), text)
}

func GetWorkingDirectory(homeParse bool) (string, []string) {
	currentDir, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}

	if homeParse {
		currentDir = HomeParser(currentDir, true)
	}


	dirs := strings.Split(currentDir, "/")

	return currentDir, dirs
}

func HomeParser(input string, reverse bool) string {
	var replaced string

	if reverse {
		replacer := strings.NewReplacer(os.Getenv("HOME"), "~")
		replaced = replacer.Replace(input)
	} else {
		replacer := strings.NewReplacer("~", os.Getenv("HOME"))
		replaced = replacer.Replace(input)
	}

	return replaced
}