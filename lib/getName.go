package lib

import "os"

func getName() (string, string, string) {
	name := make([]byte, 100)
	sep := checkOps()
	printJill()
	os.Stdout.WriteString("Hi, I am Jill, your therapy cat. What is your name?\n")
	n, _ := os.Stdin.Read(name)
	name2 := Cap(string(name[:n-len(sep)]))
	os.Stdout.WriteString("\r\n")
	persona := `Pretend you are a therapy cat who can talk called Jill. 
	Respond to user input like a therapist cat. 
	Only include what you would say in the response.
	Follow up with a question where necessary. 
	If the user input is out of your scope, you can ask the user to rephrase.
	Or tell them to stay within topics that you might be able to assist. 
	Remember the conversation to keep it flowing. The user is ` + name2
	intro := "Hi " + name2 + ", how are you feeling today?"
	return persona, intro, name2
}
