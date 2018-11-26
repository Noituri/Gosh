package main

import (
	"Gosh/Gosh"
	"Gosh/utils"
	"fmt"
	"os"
	"testing"
)

/*

	Gosh File

 */

//GOSH.GO
func TestExecute(t *testing.T) {
	args := []string{"echo", "gosh"}
	if Gosh.Execute(args) != "" {
		t.Fail()
	}
}

//HIGHLIGHTING.GO
func TestHighlight(t *testing.T) {
	h := Gosh.Highlighter {
		Command:  "#375eab",
		Quote: 	  "#1060FF",
		Argument: "#5579BF",
	}

	testCmd := `cmd test "test" --test | test test`

	testCmd = h.Highlight(testCmd, `^[A-Za-z0-9-\--_\+]*`, "command")
	testCmd = h.Highlight(testCmd, `(\|\s*)([A-Za-z0-9-]*)`, "pipe")
	testCmd = h.Highlight(testCmd, `(".*?")`, "quote")
	testCmd = h.Highlight(testCmd, `( --?[A-Za-z0-9-\--_]*)`, "argument")
	if testCmd != `\x1b[38;2;55;94;171mcmd\x1b[0m test \x1b[38;2;16;96;255m"test"\x1b[0m\x1b[38;2;85;121;191m --test\x1b[0m | \x1b[38;2;55;94;171mtest\x1b[0m test` {
		t.Fail()
	}
}

//COMMANDS.GO
func TestGetCommand(t *testing.T) {
	if Gosh.GetCommand("cd") != Gosh.Commands[0]  {
		t.Fail()
	}
}

/*

	Utils File

*/

//UTILS.GO
func TestColorize(t *testing.T) {
	if utils.Colorize("gosh", "#0000FF") != `\x1b[38;2;0;0;255mgosh\x1b[0m` {
		t.Fail()
	}
}

func TestHomeParser(t *testing.T)  {
	if utils.HomeParser("~/Documents", false) != fmt.Sprintf("%s/Documents", os.Getenv("HOME")) {
		t.Fail()
	}

	if utils.HomeParser(fmt.Sprintf("%s/Documents", os.Getenv("HOME")), true) != "~/Documents" {
		t.Fail()
	}
}