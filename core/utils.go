package core

import (
	"bufio"
	"os"
)

func Cap(s string) string {
	result := ""
	if string(s[0]) >= "a" && string(s[0]) <= "z" {
		result += string(s[0] - 32)
	}
	return result + s[1:]
}

func toLower(s string) string {
	result := ""
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			result += string(ch)
		} else {
			result += string(ch + 32)
		}
	}
	return result
}

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
