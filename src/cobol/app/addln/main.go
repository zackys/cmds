package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func DoFilter(in *os.File, out *os.File, filter func(string) string) {
	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)
	for scanner.Scan() {
		line := scanner.Text()
		writer.WriteString(filter(line))
	}
	writer.Flush()
}

type AddLine struct {
	*Command

	begin int
	step  int

	currentLn int

	outFileNm string
}

func (addLine *AddLine) AddLineFilter(line string) string {
	ret := fmt.Sprintf("%06d%s\n", addLine.currentLn, line)
	//fmt.Print(ret)
	addLine.currentLn += 1
	return ret
}

func NewAddLine() *AddLine {
	cmd := &AddLine{
		Command: NewCommand(),
	}

	fs := cmd.fs

	fs.IntVar(&cmd.begin, "start", 100000, "Line Number from")
	fs.IntVar(&cmd.step, "step", 10, "Line Number step")
	fs.StringVar(&cmd.outFileNm, "o", "", "Output File name")

	fs.Parse(os.Args[1:])

	fmt.Println(cmd.begin)
	fmt.Println(cmd.step)
	return cmd
}

//var fs = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

type Command struct {
	fs *flag.FlagSet
}

func NewCommand() *Command {
	cmd := new(Command)
	cmd.fs = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	return cmd
}

func (cmd Command) Args() []string {
	return cmd.fs.Args()
}

func main() {

	var cmd *AddLine = NewAddLine()

	outFileNm := cmd.outFileNm
	args := cmd.Args()

	var in *os.File
	var err error

	fmt.Println("1")
	if len(args) < 1 {
		in = os.Stdin
	} else {
		in, err = os.Open(args[0])
		if err != nil {
			panic(err)
		}
		defer in.Close()
	}

	fmt.Println("2")
	var out *os.File
	if len(outFileNm) < 1 {
		out = os.Stdout
		fmt.Println("3")
	} else {
		out, err = os.Open(outFileNm)
		if err != nil {
			panic(err)
		}
		fmt.Println("4")
		defer out.Close()
	}


	DoFilter(in, out, cmd.AddLineFilter)

}
