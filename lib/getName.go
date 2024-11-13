package lib

import (
	"fmt"
	"os"
	"strings"

	cfg "chatty/config"
)

func validateName(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", fmt.Errorf("name cannot be empty")
	}
	if len(name) > 100 {
		return "", fmt.Errorf("name cannot be longer than 100 characters")
	}
	if strings.ContainsAny(name, "!@#$%^&*()_+=-{}[]|:;<>,.?/~`") {
		return "", fmt.Errorf("name contains invalid characters")
	}
	return name, nil
}
func getName() (string, string, string) {
	var name string
	var err error
	for {
		printJill()
		os.Stdout.WriteString("Hi, I am Jill, your therapy cat. What is your name?\n")
		
		// Read user input and trim whitespace and newline characters
		input := make([]byte, 100)
		n, _ := os.Stdin.Read(input)
		inputName := string(name[:n-1])

		// Check if name is valid
		name, err = validateName(inputName)
		if err == nil {
			break
		}
		// If name is not valid, print error message and ask again
		fmt.Printf("Oops! %s. Please try again.\n\n", err)
	}
	// sep := checkOps()
	name2 := Cap(name)
	os.Stdout.WriteString("\r\n")
	persona := fmt.Sprintf(cfg.Persona, name2)
	intro := fmt.Sprint(cfg.DefaultIntro, name2)
	return persona, intro, name2
}
