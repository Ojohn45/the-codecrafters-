### Go Comments
Comments are text that are used for expalining a line of  code and also make it easier to read and aunderstand.
Go has teo type of comments single-line or multi-line comments.

### Go Single-line Comments:
Go single-line comments as it sounds it is only one one line, and it start with two forward slashes(//).
Any word after the slashes  will be ignored by the compiler (basicly it will not be executed).

### Example: 
// This is a comment
package main
import ("fmt")

func main() {
  fmt.Println("Hello World!") // This is a comment
}

### Go Multi-line Comments
Go Multi-line Comments as it sound you can commment on many line starting with /* and ends with */.
### Example:
package main
import ("fmt")

func main() {
  /* The code below will print Hello World
  to the screen, and it is amazing */
  fmt.Println("Hello World!")
}

### Comment to Prevent Code Execution
You can also used comment to stop a code frome execution