# Constants

```go
const pi = 3.14
pi = 1234 // Not allowed, compiler throws an error

const name
// Weirdly compiler doesn't throw any error for not using the constant name.
// The reason is something to do with memory allocation as the compiler wants to 100%
// know that if we are allocating it some memory and using it or not.

// Creating multiple constants
const (
    Monday = 1
    Tuesday = 2
    Wednesday = 3
)

const typedAge int = 25 // Constant with a datatype
const untypedAge = 27 // Constant without a datatype
```

# Enums

- Go natively does not support enums.

```go
const (
    Jan int = iota + 1 // 0 + 1
    Feb
    Mar
    Apr
)
```

- Here it is going to enumerator from 1 and so one.
- Go’s `iota` identifier is used in const declarations to simplify definitions of incrementing numbers.
- Because it can be used in expressions, it provides a generality beyond that of simple enumerations.
- The value of iota is reset to 0 whenever the reserved word const appears in the source (i.e. each const block) and incremented by one after each `ConstSpec` e.g. each Line.
- This can be combined with the constant shorthand (leaving out everything after the constant name) to very concisely define related constants.
