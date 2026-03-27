package main

import "fmt"

func main() {
	var oper int
	var cal1 int
	var cal2 int

	for {
        start:
		fmt.Println("Enter calculation: ")
		_, err1 := fmt.Scan(&cal1)


		fmt.Println("Enter calculation: ")
        _, err2 := fmt.Scan(&cal2)

        if err1 != nil || err2 != nil {
            fmt.Println("Input a valid number:")
            goto start
        }

		fmt.Println("Which operation would you like Add[1], Sub[2], Mul[3], Div[4], Quit[5]: ")
		fmt.Scan(&oper)

		a := cal1
		b := cal2

		if oper == 1 {
			fmt.Println("Adding: ", fmt.Sprint(a+b))
		} else if oper == 2 {
			fmt.Println("Subtracting:", fmt.Sprint(a-b))
		} else if oper == 3 {
			fmt.Println("Multiplying:", fmt.Sprint(a*b))
		} else if oper == 4 {
			if b == 0 {
				fmt.Println("CAN'T do any Divition")
				continue
			}
			fmt.Println("Dividing:", fmt.Sprint(a/b))
		} else if oper == 5 {
			fmt.Println("===Thanks for your time===")
			break
		} else {
			fmt.Println("Help: Not a operation choose from 1-5 for any disired operation")
		}
	}
}
