package evolve

type CommandBus interface {
	command

	Add(command command)
}

func NewCommandBus() CommandBus {
	return &commandBus{
		commands: []command{},
	}
}

type commandBus struct {
	commands []command
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
