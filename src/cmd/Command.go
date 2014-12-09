package cmd

import (
	"flag"
	"os"
)

type Command interface {
	Run()
	Args() []string
}

type CommandBase struct {
	FlagSet *flag.FlagSet
}

func NewCommandBase() *CommandBase {
	cmd := new(CommandBase)
	cmd.FlagSet = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	return cmd
}

func (cmd CommandBase) Args() []string {
	return cmd.FlagSet.Args()
}