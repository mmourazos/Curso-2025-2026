# Go

## Learning Resources

What follows is a list of resources where you can learn Go.

* [Go by example](http://gobyexample.com).
* [Efective Go](https://go.dev/doc/effective_go).

## Code organization

Go programs are organized into **packages**. A package is a **collection of source files** in the **same directory** that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.

A **module** is a **collection of** related Go **packages** that are released together. A Go repository typically contains only one module, located at the root of the repository. A file named `go.mod` there declares the **module path**. This module path is the **prefix** we need to write to import any package of that module. The module contains the packages in the directory containing its `go.mod` file as well as subdirectories of that directory, up to the next subdirectory containing another `go.mod` file (if any).

```bash
mkdir first
go mod init example/hello
```

In the previous block we are creating a directory to organize de code of our module. The name of the módulo will be `example/hello`.

Inside the `first` directory we would write the code of our packages. Usually we will create a directory for each package.

## Running code

If we want to run our code we need to have an **entry point**. The entry point for a Go application will be the `main` function in the `main` package:

```Go
package "main"

import "fmt"

func main() {
  // Our code will be here.
  fmt.Print("Hello world!")
}
```

Once we have a file with the `main` package and a `main` function we have three options to run our code:
1. Using `go run` command with the name of the file that contains the `main` function: `go run main.go`.
3. Using `go build` command to create an executable file and then running that file: `go build .` will create an exacutable file with the name of the module (in this case `hello`), then we can run it with `./hello.exe`.
4. Using `go install` command to create an executable file in the `$GOPATH/bin` directory and then running that file. In  our case the path will be `C:/Users/<User>/go/bin/hello.exe`.

```bash
go run main.go
```
