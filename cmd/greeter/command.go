package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).Level(zerolog.InfoLevel)

	subcommands.Register(subcommands.HelpCommand(), "help")
	subcommands.Register(subcommands.FlagsCommand(), "help")
	subcommands.Register(subcommands.CommandsCommand(), "help")
	subcommands.Register(NewGreetCommand(), "")
	subcommands.Register(NewSuperGreetCommand(), "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
