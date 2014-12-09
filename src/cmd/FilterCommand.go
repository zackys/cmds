package cmd

import (
	"bufio"
	"os"
)

type FilterCommandIF interface {
	filter(string) string

	Args() []string

	outFileNm() string
}

type FilterCommand struct {
	*CommandBase

	outFileNm string

	Filter func(string)string
}

func NewFilterCommand() *FilterCommand {
	cmd := &FilterCommand {
		CommandBase: NewCommandBase(),
	}

	fs := cmd.FlagSet
	fs.StringVar(&cmd.outFileNm, "o", "", "Output File name")

	return cmd
}

func doFilter(in *os.File, out *os.File, filter func(string) string) {
	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)
	for scanner.Scan() {
		line := scanner.Text()
		writer.WriteString(filter(line))
	}
	writer.Flush()
}

func (cmd FilterCommand) Run() {

	outFileNm := cmd.outFileNm
	args := cmd.Args()

	var in *os.File
	var err error

	if len(args) < 1 {
		in = os.Stdin
	} else {
		in, err = os.Open(args[0])
		if err != nil {
			panic(err)
		}
		defer in.Close()
	}

	var out *os.File
	if len(outFileNm) < 1 {
		out = os.Stdout
	} else {
		out, err = os.Open(outFileNm)
		if err != nil {
			panic(err)
		}
		defer out.Close()
	}

	doFilter(in, out, cmd.Filter)
}
