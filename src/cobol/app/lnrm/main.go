package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
)

func main() {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	var trimCol = fs.Int("col", 6, "Line Number from")

	fs.Parse(os.Args[1:])

//	fmt.Println("Begin:", *trimCol)
//	fmt.Println("File", fs.Args())

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
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line[*trimCol:])
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
