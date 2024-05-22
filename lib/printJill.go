package lib

import (
	"bufio"
	"os"
)

func printJill() {
	file, err := os.Open("jill.txt")
	if err != nil {
		os.Stdout.WriteString("Error: problem opening file")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		os.Stdout.WriteString(scanner.Text() + "\n")
	}

}
