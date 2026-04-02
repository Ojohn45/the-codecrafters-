### Go switch Statement
The switch statement is used to select a block of code to executable.
The value of the switch expression is compared with the values of each case. If there is a match, the associated block of code is executed
The default keyword specifies the code if there is no case to match

Example of switch cases
Single-Case switch Example:
The example below uses a weekday number to calculate the weekday name:
```go
package main
import ("fmt")

func main() {
  day := 4

  switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  case 3:
    fmt.Println("Wednesday")
  case 4:
    fmt.Println("Thursday")
  case 5:
    fmt.Println("Friday")
  case 6:
    fmt.Println("Saturday")
  case 7:
    fmt.Println("Sunday")
  }
}
```
