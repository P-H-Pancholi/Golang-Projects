package main

import (
	"fmt"
	"os"

	"github.com/P-H-Pancholi/Golang-Projects/gator/internal/config"
)

type state struct {
	c *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	list_of_commands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.list_of_commands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.list_of_commands[cmd.name]
	if !ok {
		return fmt.Errorf("command does not exists, please register it")
	}
	if err := f(s, cmd); err != nil {
		return err
	}
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("expected one argument, got zero")
	}
	if err := s.c.SetUser(cmd.args[0]); err != nil {
		return err
	}
	fmt.Println("The user has been set")
	return nil
}

func main() {
	c, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	s := state{
		c: &c,
	}
	comm := commands{
		make(map[string]func(*state, command) error),
	}
	comm.register("login", handlerLogin)
	if len(os.Args) < 3 {
		fmt.Println("Expected more than 1 args")
		os.Exit(1)
	}
	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}
	if err := comm.run(&s, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
