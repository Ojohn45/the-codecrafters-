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

		if strings.ToLower(input) == "exit" {
			fmt.Println("\033[34;1m -Shutting down String Transformer. Goodbye.\033[0m")
			break
		}

		if strings.ToLower(input) == "help" {
			Help()
			continue
		}

		var history []string
		if input == "history" {
			fmt.Println("Operation History")
			for i, op := range history {
				fmt.Printf("[%d] %s\n", i+1, op)
			}
			if len(history) == 0 {
				fmt.Println("No history yet")
			} else {
				fmt.Println("--------")
				continue
			}
		}
		// split command + text
		part := strings.Fields(input)

		if len(part) <= 0 {
			fmt.Println("\033[31;1m -Please provide operation and text\033[0m")
			continue
		}

		conv := strings.ToLower(part[0])
		text := strings.Join(part[1:], " ")

		switch conv {
		case "upper":
			fmt.Println(toUpper(text))
			history = append(history, text)
			continue

		case "lower":
			fmt.Println(toLower(text))
			history = append(history, fmt.Sprintf(text))
			continue

		case "cap":
			fmt.Println(capFirstLetter(text))
			history = append(history, fmt.Sprintf(text))
			continue

		case "snakecase":
			fmt.Println(snakecase(text))
			history = append(history, fmt.Sprintf(text))
			continue

		case "reverse":
			result := text
			fmt.Println(reverse(text))
			history = append(history, fmt.Sprintf(text, result))
			continue

		default:
			fmt.Println("\033[31;1m -Invalid operation\033[0m")
		}
	}
}
