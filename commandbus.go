package evolve

type commandBus struct {
	command

	commands []command
}

func newCommandBus() *commandBus {
	return &commandBus{
		commands: []command{},
	}
}

func (c *commandBus) Execute(backEnd BackEnd) error {
	// Execute all the commands.
	for _, command := range c.commands {
		err := command.Execute(backEnd)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *commandBus) Add(command command) {
	c.commands = append(c.commands, command)
}
