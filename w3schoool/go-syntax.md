### Go Syntax
So when we talk about syntax it means the set of rules that define the correct structure and format for writing code in a specific language.
### Now a Go file consists of these following parts:
* Package declaration
* Import packages
* Functions
* Statements and expressions

### For Example:
package  
// So in Go file or program is part of a package we define this using the package keyword.
import ("fmt") 
// This line helps us import files like fmt package. 

func main() { 
// This is a function. Any well writen code inside its curly brackets {} will be executed.
  fmt.Println("Hello World!") 
// This is a function made available from the fmt package. This is used to print out text like "Hello World!" in our example. 
}

### Go Statements
In go something like fmt.Println("Hello World!") is a statement.

### Go Compact Code
You can also write more compact code, like shown below (this is not adviseable because it makes the code looks very difficult to read):

### Example:
 package main; import ("fmt"); func main() { fmt.Println("Hello World!");}