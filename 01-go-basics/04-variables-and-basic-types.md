# Basic data types

- Intger (`int`)
- Float (`float32` or `float64`)
- Boolean (`bool`)
- String (`string`)

# Variables in Go

```go
var name string = "Parth"
fmt.Println("Hello", name)
```

- In Go, we need to always use the variable that we have declare as the compiler is pretty strict for unused variables.

## Using the format specifier

```go
var name string = "Parth Maru"
fmt.Printf("Hi %s\n", name);
```

- Here we use `Printf()` method to print.
- We use `%s` as a format specifier for string type.
- `Printf()` does not adds a new line, so we need to manually write `\n` character.

```go
age := 27 // Another way to create a variable (Shorthand operator)
fmt.Printf("My age is %d\n",age)
```

- Here `%d` is for integers.

```go
var city, country string = "Mumbai", "India"
fmt.Printf("I am from %s which is in %s\n", city, country)

// For different data types
var (
    isEmployed bool = true
    salary int = 50000
    position string = "Developer"
)

fmt.Printf("Are you employed? %t\n", isEmployed)
fmt.Printf("I'm a %s with a salary of %d\n", position, salary)
```

## Default or Zero values

```go
var defaultInt int
var defaultFloat float64
var defaultBool bool
var defaultString string

fmt.Println("Default values: ", defaultInt, defaultFloat, defaultBool, defaultString) // 0 0 false "empty string"
```
