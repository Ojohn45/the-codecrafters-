package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func toUpper(conv string) string {
	return strings.ToUpper(conv)
}

func toLower(conv string) string {
	return strings.ToLower(conv)
}

func capFirstLetter(conv string) string {
	w := strings.Fields(conv)
	for i := range w {
		if len(w[i]) > 0 {
			w[i] = strings.ToUpper(w[i][:1]) + w[i][1:]
		}
	}
	return strings.Join(w, " ")
}

func snakecase(word string) string {
	return strings.ReplaceAll(word, " ", "_")
}

func reverse(word string) string {
	if len(word) == 0 {
		return ""
	}
	return reverse(word[1:]) + word[:1]
}

func Help() {
	fmt.Println("\033[33;1m -Welcome to the help station\033[0m")
	fmt.Println("\033[33;1m -Type: upper/lower/cap followed by your text\033[0m")
	fmt.Println("\033[33;1m Example: upper fine boy\033[0m")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var history []string

	for {
		fmt.Println("\033[32;1m -Welcome To Operation Save The World\033[0m")
 		fmt.Println("\033[32;1m -If you need Help please type help for assistance\033[0m")
 		fmt.Println("\033[32;1m -If you do not wish to continue please type exit\033[0m")
 		fmt.Println("\033[32;1m -If you do not have any problem please continue\033[0m")

		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			continue
		}

		switch strings.ToLower(input) {
		case "exit":
			fmt.Println("\033[34;1m -Shutting down String Transformer. Goodbye.\033[0m")
			return

		case "help":
			Help()
			continue

		case "history":
			fmt.Println("Operation History:")
			if len(history) == 0 {
				fmt.Println("No history yet")
			} else {
				for i, op := range history {
					fmt.Printf("[%d] %s\n", i+1, op)
				}
			}
			continue
		}

		part := strings.Fields(input)

		if len(part) < 2 {
			fmt.Println("\033[31;1m -Please provide operation and text\033[0m")
			continue
		}

		conv := strings.ToLower(part[0])
		text := strings.Join(part[1:], " ")

		var result string

		switch conv {
		case "upper":
			result = toUpper(text)

		case "lower":
			result = toLower(text)

		case "cap":
			result = capFirstLetter(text)

		case "snakecase":
			result = snakecase(text)

		case "reverse":
			result = reverse(text)

		default:
			fmt.Println("\033[31;1m -Invalid operation\033[0m")
			continue
		}

		fmt.Println(result)

		history = append(history, fmt.Sprintf("%s -> %s", input, result))
	}
}
