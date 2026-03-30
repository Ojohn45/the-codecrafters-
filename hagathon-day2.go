package main

import (
	"strconv"
	"strings"
	"bufio"
	"fmt"
	"os"
)
   
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	blue := "\033[34m"
	red := "\033[31m"
	green := "\033[32m"
	yellow := "\033[33m"
	reset := "\033[0m"

	for {

		fmt.Print(green + "Enter convertion " + reset)

		fmt.Print(yellow + "> " + reset)
        if !scanner.Scan() {
            break
        }
	
	input := strings.TrimSpace(scanner.Text())
	if input == "" {
		continue
	}

	if input == "quit" {
		break
	}

	conv := strings.Fields(input)
	if len(conv) != 3 || conv[0] != "convert" {
		fmt.Println(red + "Invalid operation use convert value, base" + reset)
		continue
	}
	  
	value := conv[1]
	Base := conv[2]

	var base int
	switch Base {
	case  "hex":
		base = 16
	case "bin":
    	base = 2
    case "dec":
        base = 10
	}

	no, err := strconv.ParseInt(value, base, 64)
	if err != nil {
		fmt.Println(red + "Invalid base number: " + reset, Base)
		continue
	}

        switch Base {
        case "dec":
            fmt.Println(blue + "+ Binary: " + reset, strconv.FormatInt(no, 2))
            fmt.Println(blue + "+ Hex: " + reset, strings.ToUpper(strconv.FormatInt(no, 16)))
        default:
            fmt.Println(blue + "+ Decimal:" + reset, no)
        }
	}
}
