package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
)

func main() {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	var beginLn = fs.Int("b", 100000, "Line Number from")
	var stepLn = fs.Int("s", 10, "Line Number step")

	fs.Parse(os.Args[1:])

	fmt.Println("Begin:", *beginLn)
	fmt.Println("Step:", *stepLn)
	fmt.Println("File", fs.Args())

	var fp *os.File
	var err error

	if len(fs.Args()) < 1 {
		fp = os.Stdin
	} else {
		fp, err = os.Open(fs.Args()[0])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	scanner := bufio.NewScanner(fp)
	var ln int = *beginLn
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Print(fmt.Sprintf("%06d %s\n", ln, line))
		ln += *stepLn
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
