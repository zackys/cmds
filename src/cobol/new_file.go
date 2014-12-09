package cobol

import (
	"os"
	"bufio"
)

func DoFilter(in *os.File, out * os.File, filter func(string) string) {
	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)
	for scanner.Scan() {
		line := scanner.Text()
		writer.WriteString(filter(line))
	}
}
