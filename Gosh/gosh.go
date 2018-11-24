package Gosh

import (
	"bytes"
	"fmt"
	"github.com/chzyer/readline"
	"io"
	"os"
	"os/exec"
	"strings"
)

func Execute(args []string) {
	cmd := exec.Command("sh", append([]string{"-c"}, strings.Join(args, " "))...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Run()

	if stderr.String() != "" {
		fmt.Println(stderr.String())
	}

}

func Start() string {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "\033[31m»\033[0m ",
		HistoryFile:     fmt.Sprintf(`%s/history.gosh`, os.Getenv("HOME")),
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
		Painter:		 &Highlighter{
			command:  "#FF00FF",
			quote: 	  "#0F42BB",
			argument: "#44FA12",
		},
	})

	if err != nil {
		panic(err)
	}

	defer rl.Close()
	
	for {
		input, err := rl.Readline()
		if err == readline.ErrInterrupt {
			if len(input) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		input = strings.TrimSpace(input)

		args := strings.Split(input, " ")
		command := args[0]

		switch command {
		case "exit":
			fmt.Println("See you soon!")
			goto exit
		case "reload":
			fmt.Println("Reloading shell")
			goto reload
		}

		Execute(args)
	}

	exit:
		return "exit"
	reload:
		return "reload"
}
