package Gosh

import (
	"Gosh/utils"
	"regexp"
)

type Highlighter struct {
	Command  string
	Quote    string
	Argument string
}

func (h *Highlighter) Highlight(line string, regex string, hType string) string {
	hReg, err := regexp.Compile(regex)

	if err != nil {
		return  line
	}

	switch hType {
	case "command":
		line = hReg.ReplaceAllString(line, utils.Colorize(hReg.FindString(line), h.Command))
	case "pipe":
		line = hReg.ReplaceAllString(line, "$1" + utils.Colorize("$2", h.Command))
	case "quote":
		line = hReg.ReplaceAllString(line, utils.Colorize("$1", h.Quote))
	case "argument":
		line = hReg.ReplaceAllString(line, utils.Colorize("$1", h.Argument))
	}
	return line
}

func (h *Highlighter) Paint(line []rune, _ int) []rune {
	strLine := string(line)

	strLine = h.Highlight(strLine, `^[A-Za-z0-9-\--_\+]*`, "command")
	strLine = h.Highlight(strLine, `(\|\s*)([A-Za-z0-9-]*)`, "pipe")
	strLine = h.Highlight(strLine, `(".*?")`, "quote")
	strLine = h.Highlight(strLine, `( --?[A-Za-z0-9-\--_]*)`, "argument")

	return []rune(strLine)
}