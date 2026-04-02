### Go Functions
A function is a block of code that performs a task. It helps you reuse code instead of writing it many times. A function does not run automatically. It only runs when you call it.

### Create a Function
You create a function using the func keyword.
You give the function a name.
The code it runs is written inside { }.
Example: func greet() { ... }
To use it, you call it like: greet()

### Call a Function
A function can be called many times.
This makes your code clean and organized.
Function names must start with a letter.
They can contain letters, numbers, and underscores only.
Names are case-sensitive (small and capital letters matter).
Always give your function a clear name that explains what it does.

### Go Function Parameters and Arguments
Parameters are variables used in a function. They allow you to pass information into a function. Parameters are written inside the function’s parentheses. 

### Multiple Parameters
Each parameter must have a name and a type. Example: func greet(name string), when calling a function, you pass arguments. Arguments are the actual values you send in.

Example: greet("John") → "John" is the argument. Inside the function, the parameter uses that value, you can have one or many parameters and you can separate multiple parameters with a comma.
Example: func info(name string, age int)

The order of arguments must match the parameters. Parameters make functions more flexible and reusable.
