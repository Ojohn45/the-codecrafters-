### Go Output Functions
Here are three function:
* Print()
* Println()
* Printf()

### Go Formatting Verbs
Go has many formatting verbs that can be used with the Printf() function. like Integer Formatting Verbs, String Formatting Verbs, Boolean Formatting Verbs, Float Formatting Verbs
### Example of formating verb
package main
import ("fmt")

func main() {
  var i = 15.5
  var txt = "Hello World!"

  fmt.Printf("%v\n", i)
  fmt.Printf("%#v\n", i)
  fmt.Printf("%v%%\n", i)
  fmt.Printf("%T\n", i)

  fmt.Printf("%v\n", txt)
  fmt.Printf("%#v\n", txt)
  fmt.Printf("%T\n", txt)
}