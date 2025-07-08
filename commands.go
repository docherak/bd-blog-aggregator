package main

type command struct {
	Name string
	Args []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if err := c.commands[cmd.Name](s, cmd); err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}
