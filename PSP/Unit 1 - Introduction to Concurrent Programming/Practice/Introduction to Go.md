# Go

Go is a programming language that was originally designed to solve various problems in large-scale software development in the real world, initially within Google. It addresses slow program construction, out-of-control dependency management, complex code, and difficult cross-language
construction.

Each language tries to address these problems in a different way, either by restricting the user or by being as flexible as possible. Go’s team chose to tackle these problems by targeting modern engineering: removing the constraint of dealing with memory and making it simple to run parallel pieces of code.

Go was built for concurrency and networked servers, which explains the fast adoption of the language in software companies of all sizes in the past few years. Go was designed to improve productivity during a time when multicore networked machines and large codebases were becoming the norm.

> Go started in September 2007 when Robert Griesemer, Ken Thompson, and I began
> discussing a new language to address the engineering challenges we and our colleagues at Google were facing in our daily work.
> —Rob Pike, coauthor of Go

Go was designed with simplicity and productivity in mind, making it a natural fit for backend services, APIs, and modern cloud computing needs. In this section, we’ll explore practical examples of how Go’s syntax, tools, and features come together to create maintainable and efficient applications.

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
func addition(a int, b int) int {
    return a + b
}

// We can call the function like this:
result := addition(3, 4)

fmt.Printf("The result is: %d", result)
```

We can also assign functions to variables (and pass them as arguments to other functions):

```Go

// We can assign a function to a variable.
add := addition

// And we can call the function using the variable.
result := add(3, 4)

fmt.Printf("The result is: %d", result)
```

Function in go can return multiple values:

```Go
value, err := strconv.ParseBool("true")
```

The previous example shows how to convert a string to a boolean. The `ParseBool` function returns two values: the parsed boolean value and an error value. If the conversion is successful, the error value will be `nil`. If the conversion fails, the error value will contain an error message.

If we want to ignore one of the returned values, we can use the blank identifier `_`:

```Go
value, _ := strconv.ParseBool("true")
```
