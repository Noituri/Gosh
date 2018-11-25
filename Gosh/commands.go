package Gosh

import "Gosh/Gosh/commands"

type Command interface {
	CommandName() 	 string
	CommandExecute(args...string)
}

var Commands = []Command{&commands.Cd{Name: "cd"}}

func GetCommand(cmdName string) Command {
	for _, v := range Commands {
		if v.CommandName() == cmdName {
			return v
		}
	}
	return nil
}
