package main

import (
	"Gosh/Gosh"
	"os"
)

func main() {
	switch Gosh.Start() {
	case "exit":
		os.Exit(0)
	case "reload":
		Gosh.Start()
	}
}
