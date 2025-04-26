# Functions in Go

```go
func someFunction() {

}
```

- This function cannot be exported even if it is an exportable package.

```go
func SomeFunction() {

}
```

- But this one is.
- This is because it starts with an upper case.
- So in order to export a function, it should start with an upper case.

```go
// Syntax
func name/Name (param1 type, param2 type) returnType {
    // Function body
}

// Example
func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(3,4)
    println(result)
}
```

## Multiple params with same type

```go
func sum(a, b, c int) int {
    return a + b + c
}

```

- Here every param on the left of `int` type will have `int` as a type.

## Returning multiple values

```go
func sumAndProduct(a, b int) (int, int) {
    return (a + b), (a * b)
}

func main() {
    sum, product := sumAndProduct(3, 5)
    println(sum, product)
}
```

- If we want to return multiple types, add them in `()`.
- Return the values using `,` separated `()`.
- If you want to use only one return value from the function, you can use the `_` to ignore the other values.

```go
func main() {
    sum, _ := sumAndProduct(3,5)
    println(sum)
}
```
