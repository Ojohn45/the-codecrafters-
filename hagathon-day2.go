package main

import (
	"fmt"
	"strconv"
)
   
func main() {

	for {
	var conver string
	fmt.Println("Enter convertion")
	fmt.Print("> ")
	fmt.Scan(&conver)

	if conver == "" {
		continue
	}

	if conver == "quit" {
		break
	}

	if len(conver) != 3 || conver[0] != "conver" {
		fmt.Println("Invalid operation use convert value, base")
		continue
	}
	  
	value := conver[1]
	base := conver[2]

	num, err := strconv.ParseInt(value, base, 64)
	if err != nil {
		fmt.Println("Invalid base number: ", base)
		continue
	}

	var base int
	switch base {
	case  "hex"

		fmt.Println()
	}
	}
}
