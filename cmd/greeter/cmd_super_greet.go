package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"github.com/rs/zerolog/log"
)

var _ subcommands.Command = (*superGreetCmd)(nil)

// superGreetCmd is just a simple command which will print super hello for specified username
type superGreetCmd struct {
	userName string
}

func (g *superGreetCmd) Name() string {
	return "super_greet"
}

func (g *superGreetCmd) Synopsis() string {
	return "greet will super great you, username!"
}

func (g *superGreetCmd) Usage() string {
	return "super_greet --user USERNAME"
}

func (g *superGreetCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&g.userName, "user", "ADMIN", "username to greet. Default username is 'ADMIN'. Example: super_greet --user ADMIN_USER")
}

func (g *superGreetCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	log.Info().Msgf("Super hello for %s!", g.userName)
	return subcommands.ExitSuccess
}

// NewSuperGreetCommand returns new instance of command which will print super hello to provided username.
func NewSuperGreetCommand() subcommands.Command {
	return &superGreetCmd{}
}
