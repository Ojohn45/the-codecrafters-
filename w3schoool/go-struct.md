A struct in Go is used to group different types of data into one variable.

While arrays store many values of the same type, a struct stores different types together (like name, age, job, etc.).

You create a struct using type and struct, and define its fields (members), like:

name (string)
age (int)
job (string)
salary (int)

To use a struct:

Create a variable of that struct
Assign values to each field using the dot (.) operator

Example:

person.name = "John"
person.age = 30

You can also:

Access and print the values using the dot operator
Pass a struct into a function, just like any other variable
