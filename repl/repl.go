package repl

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"

	"interpreter/lexer"
)

func Initiate() error {

	user, err := user.Current()
	if err != nil {
		return err
	}
	username := strings.ToUpper(string(user.Username[0])) + user.Username[1:]
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Welcome to Zen's interpreter %s. Please type an expression 🤟\n", username)

	for {
		fmt.Printf("➥  ")
		scanned := scanner.Scan()
		if !scanned {
			break
		}
		inputVal := scanner.Text()

		if inputVal == "exit" {
			fmt.Println("Exiting gracefully")
			break
		}
		l := lexer.NewLexer(inputVal)

		tokens, err := l.Tokenize()
		if err != nil {
			fmt.Println(err)
		}
		for i := range tokens {
			fmt.Println("✲ "+"Type: "+tokens[i].Type, ", ", "Value: "+tokens[i].Value)
		}

	}
	return nil
}
