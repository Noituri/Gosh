package Gosh

import (
	"Gosh/utils"
	"bytes"
	"fmt"
	"github.com/chzyer/readline"
	"io"
	"os"
	"os/exec"
	"strings"
)

func Execute(args []string) string {
	cmd := exec.Command("sh", append([]string{"-c"}, strings.Join(args, " "))...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Run()

	if stderr.String() != "" {
		fmt.Println(stderr.String())
	}

	return stderr.String()
}

func filterInput(r rune) (rune, bool) {
	switch r {
	case readline.CharInterrupt:
		fmt.Println(fmt.Sprintf("%s: type \"%s\".",
			utils.Colorize("Gosh", "#2874a6"),
			utils.Colorize("exit", "#c0392b")),
		)
	}
	return r, true
}

func Start() string {

	_, dirs := utils.GetWorkingDirectory(true)

	rl, err := readline.NewEx(&readline.Config{
		Prompt:          utils.Colorize(fmt.Sprintf("%s » ", utils.Colorize(dirs[len(dirs) - 1], "#01A0CD")), "#5579BF"),
		HistoryFile:     fmt.Sprintf(`%s/history.gosh`, os.Getenv("HOME")),
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
		FuncFilterInputRune: filterInput,
		Painter: &Highlighter {
			Command:  "#375eab",
			Quote: 	  "#1060FF",
			Argument: "#5579BF",
		},
	})

	if err != nil {
		panic(err)
	}

	defer rl.Close()
	
	for {
		input, err := rl.Readline()
		if err == readline.ErrInterrupt {
			continue
		} else if err == io.EOF {
			break
		}

		input = strings.TrimSpace(input)
		input = utils.HomeParser(input, false)

		args := strings.Split(input, " ")
		command := args[0]

		switch command {
		case "exit":
			fmt.Println("See you soon!")
			goto exit
		case "reload":
			fmt.Println("Reloading shell")
			goto reload
		default:
			cmd := GetCommand(command)

			if cmd != nil {
				if len(args) > 1 {
					cmd.CommandExecute(args[1])
				}
			} else {
				Execute(args)
			}
		}

		_, dirs := utils.GetWorkingDirectory(true)
		rl.SetPrompt(utils.Colorize(fmt.Sprintf("%s » ", utils.Colorize(dirs[len(dirs) - 1], "#01A0CD")), "#5579BF"))
	}

	exit:
		return "exit"
	reload:
		return "reload"
}
