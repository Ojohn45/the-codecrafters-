package main

import (
	"fmt"
	"strconv"
)

func main() {

	blue := "\033[34m"
	red := "\033[31m"
	green := "\033[32m"
	yellow := "\033[33m"
	reset := "\033[0m"

	var oper string
	var a string
	var b string
	var Help string

	for {
		
		fmt.Println(green + "Welcome to my smart calculator" + reset)
		fmt.Println(green + "If you are having any trouble please type help for more assitance" + reset)
		fmt.Println(green + "If you don't have any type cont to continue" + reset)
		fmt.Println(green + "And if don't want to continue operation type quit" + reset)
		fmt.Scan(&Help)

		start:
		if Help == "help" {
			fmt.Println(yellow + "First enter a number" + reset)
			fmt.Println("Then from the option 1-5 for any operation")
			goto start2
		} else if Help == "quit" {
			fmt.Println(green + "===Thanks for your time===" + reset)
			break
		} else if Help == "cont" {
			fmt.Println(yellow + "Let start calculating" + reset)
			goto start2
		} else {
			fmt.Println(red + "This is not an operation choose a proper operation" + reset)
			goto start
		}

		start2:
		fmt.Println(blue + "Enter Fist no: " + reset)
		fmt.Scan(&a)

		fmt.Println(blue + "Enter second no: " + reset)
        fmt.Scan(&b)
		a1 , err1 := strconv.Atoi(a)
		b1 , err2 := strconv.Atoi(b)

        if err1 != nil || err2 != nil {
            fmt.Println(red + "Input A Valid Number " + reset)
			fmt.Println(red + "Try Again" + reset)
            goto start
        }

		if a == "quit" || b == "quit" {
			fmt.Println(green + "===Thanks for your time===" + reset)
			break
		}

		fmt.Println(blue + "Which operation would you like Add[1], Sub[2], Mul[3], Div[4]: " + reset)
		fmt.Scan(&oper)

		if oper == "1" {
			fmt.Println(green + "Added:" + reset, fmt.Sprint(a1+b1))
		} else if oper == "2" {
			fmt.Println(green + "Subtracted:" + reset, fmt.Sprint(a1-b1))
		} else if oper == "3" {
			fmt.Println(green + "Multiplied:" + reset, fmt.Sprint(a1*b1))
		} else if oper == "4" {
			if b1 == 0 {
				fmt.Println(red + "CAN'T be divided by 0" + reset)
				fmt.Println(red + "Lets start again" + reset)
				goto start2
			}
			fmt.Println(green + "Divided:" + reset, fmt.Sprint(a1/b1))
		} else if oper == "quit" {
			fmt.Println(green + "===Thanks for your time===" + reset)
			break
		} else {
			goto start
		}
	}
}