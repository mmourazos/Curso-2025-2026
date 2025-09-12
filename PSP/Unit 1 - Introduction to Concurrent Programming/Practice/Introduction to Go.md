# Go

## Basic syntax

## Variables

### Data types

### Data structures

### Objects in Go

## Basic statements

### Conditionals

The conditional statements in Go are similar to those in other programming languages. The primary conditional statements are `if`, `else if`, and `else`.

```Go
// Code before...

x = 10

if x > 0 {
    fmt.Println("x is positive")
} else if x < 0 {
    fmt.Println("x is negative")
} else {
    fmt.Println("x is zero")
}

// Code after...
```

The second conditional statement is `switch`, which is used to select one of many code blocks to be executed.

```Go
import (
    "fmt"
    "runtime"
)
// Code before...

l := "linux"

switch os := runtime.GOOS; os { // We can initialize a variable in the switch statement.
case "darwin":
    fmt.Println("OS X.")
    // break is implicit.
case l: // We can use variables as case values.
    fmt.Println("Linux.")
    // break is implicit.
default:
    // freebsd, openbsd,
    // plan9, windows...
    fmt.Printf("%s.\n", os)
}

// Code after...
```

### Loops

Go has only one looping construct, the `for` loop. This `for` loop can be used in several ways.

#### C-style

```Go
for i := 0; i < 10; ++i {
    fmt.Println(i)
}
```

#### While-style

```Go
condition := true
i := 0
for condition {
    fmt.Println(i)
    condition = i < 10
    ++i
}
```

#### Loop over a collection

```Go
for key, value := range collection {
    fmt.Println(key, value)
}
```

## Functions

To define a function in Go, we use the `func` keyword followed by the function name, parameters, and return type.

```Go
func addtion(a int, b int) int {
    return a + b
}

// We can assign a function to a variable.
add := addition
```
