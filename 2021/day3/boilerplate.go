package main

import (
	"bufio"
	"os"

	. "helpers"
)

type Something struct {
}

func (self *Something) GetInput(s string) {
}

func main() {
	input := OpenInputFile()
	defer CloseFile(input)

	scanner := bufio.NewScanner(input)

	something := Something{}

	for scanner.Scan() {
		something.GetInput(scanner.Text())
	}

	os.Exit(0)
}
