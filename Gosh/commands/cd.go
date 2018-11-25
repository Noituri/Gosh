package commands

import (
	"fmt"
	"os"
)

type Cd struct {
	Name string
}

func (c *Cd) CommandName() string {
	return c.Name
}

func (c *Cd) CommandExecute(args...string) {
	if args != nil {
		err := os.Chdir(args[0])

		if err != nil {
			fmt.Println(err)
		}
	}
}