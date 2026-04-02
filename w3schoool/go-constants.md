### Go Constants
If you want to have a fixed variable that cannot be change const keyword is the right way to do that. So using the const keyword you can creat a variable that is unchangeable and read-only. If you are using the const keyword it should be in upper case for easy identificaction and differentiation from variables, constants can be declared both inside and outside of a function

### Example:
```go
package main
import ("fmt")

const PI = 3.14

func main() {
  fmt.Println(PI)
}
```
