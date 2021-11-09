package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"github.com/rs/zerolog/log"
)

var _ subcommands.Command = (*greetCmd)(nil)

// greetCmd is just a simple command which will print hello for specified username
type greetCmd struct {
	userName string
}

func (g *greetCmd) Name() string {
	return "greet"
}

func (g *greetCmd) Synopsis() string {
	return "greet will great you, username!"
}

func (g *greetCmd) Usage() string {
	return "greet --user USERNAME"
}

func (g *greetCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&g.userName, "user", "", "username to greet. Example: greet --user ADMIN")
}

func (g *greetCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	// user is mandatory parameter
	if g.userName == "" {
		log.Error().Msgf("mandatory parameter 'username' is not provided. Execute `greeter greet --help` to check command syntax.")
		return subcommands.ExitUsageError
	}

	log.Info().Msgf("Hello %s!", g.userName)
	return subcommands.ExitSuccess
}

// NewGreetCommand returns new instance of command which will print hello to provided username.
func NewGreetCommand() subcommands.Command {
	return &greetCmd{}
}
