package cmd

import (
	"fmt"
	"os"
)

type AddLineNumber struct {
	*FilterCommand

	begin int
	step  int

	currentLn int

	outFileNm string
}

func (addLine *AddLineNumber) FilterImpl(line string) string {
	ret := fmt.Sprintf("%06d%s\n", addLine.currentLn, line)
	//fmt.Print(ret)
	addLine.currentLn += 1
	return ret
}

func NewAddLineNumber() *AddLineNumber {
	cmd := &AddLineNumber{
		FilterCommand: NewFilterCommand(),
	}
	//cmd.Filter = cmd.FilterImpl
	cmd.Filter = func(line string) string {
		ret := fmt.Sprintf("%06d%s\n", cmd.currentLn, line)
		//fmt.Print(ret)
		cmd.currentLn += 1
		return ret
	}

	fs := cmd.FlagSet

	fs.IntVar(&cmd.begin, "start", 100000, "Line Number from")
	fs.IntVar(&cmd.step, "step", 10, "Line Number step")

	fs.Parse(os.Args[1:])

	fmt.Println(cmd.begin)
	fmt.Println(cmd.step)
	return cmd
}
