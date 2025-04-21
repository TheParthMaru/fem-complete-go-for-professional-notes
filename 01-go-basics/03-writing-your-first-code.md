# Hello World in Go

- Create a file called `main.go`.
- Go compiler is very opinionated and wants us to follow a set of rules.
- Any go project starts with a `main()` function in the `main` package.

```go
package main
```

- The compiler is going to always look for the `main` package as we want to compile it into an executable binary.

## Create the `main()` function

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

- We create a function using `func` keyword.
- We are importing a package called `fmt` that has methods such as `Println` to print something on the console.- Use `gopls` as LSP for Go.
- As for the go compiler, if an import is unused, it will throw an error.
- The `main` package cannot export anything but can always import other packages.

## Running Go code

```bash
go run <filename>

go run main.go
-> Hello World
```

## Creating a binary

```bash
go build -o <executable_name> <src_file>

go build -o main main.go

# Executing the above created binary
./main
```

- If we change anything in the code and save it without creating a binary for it, the `./main` file won't be updated.
- We will have to re-create the binary in order to reflect the changes.
