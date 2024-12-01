package files

import (
	"bufio"
	"os"
	"strings"
)

func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	linescan := bufio.NewScanner(file)
	linescan.Split(bufio.ScanLines)

	lines := []string{}
	for linescan.Scan() {
		line := linescan.Text()
		if len(strings.TrimSpace(line)) > 0 {
			lines = append(lines, line)
		}
	}

	return lines
}
