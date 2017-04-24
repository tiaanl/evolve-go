package evolve

type CommandBus interface {
	Command

	Add(command Command)
}

func NewCommandBus() CommandBus {
	return &commandBus{
		commands: []Command{},
	}
}

type commandBus struct {
	commands []Command
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

func (c *commandBus) Add(command Command) {
	c.commands = append(c.commands, command)
}
