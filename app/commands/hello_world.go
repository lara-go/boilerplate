package commands

import (
	"github.com/lara-go/larago/logger"

	"github.com/urfave/cli"
)

// HelloWorld console command.
type HelloWorld struct {
	Logger *logger.Logger
}

// GetCommand for the cli to register.
func (c *HelloWorld) GetCommand() cli.Command {
	return cli.Command{
		Name:  "hello",
		Usage: "Example Hello World command",
	}
}

// Handle command.
func (c *HelloWorld) Handle(args cli.Args) error {
	c.Logger.Success("Hello World!")

	return nil
}
