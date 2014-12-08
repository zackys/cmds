package main

import (
	"bufio"
	"flag"
	"fmt"
	"nkf/encode"
	"os"
)

func main() {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
//	var ic = fs.String("ic", "Shift_JIS", "Line Number from")
//	var oc = fs.String("oc", "UTF-8", "Line Number from")

	fs.Parse(os.Args[1:])

//	fmt.Println("ic:", *ic)
//	fmt.Println("oc:", *oc)
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
		t, _ := encode.SJIS_to_UTF8(line)
		fmt.Println(t)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
