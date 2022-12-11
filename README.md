# go play

The Go Playground is a web service that runs on go.dev's servers. The service receives a Go program, vets, compiles, links, and runs the program inside a sandbox, then returns the output.


https://go.dev/play/

# Basic Commands

go version: print Go version

go build: compile packages and dependencies

go run: compile and run Go program

go clean: remove object files and cached files

go env: print Go environment information

go help: go help 

# Famous Packages

https://pkg.go.dev/search?q=quote

fmt

quote

# Add new module requirements and sums.

Go will add the quote module as a requirement, as well as a go.sum file for use in authenticating the module.

$ go mod tidy
go: finding module for package rsc.io/quote
go: found rsc.io/quote in rsc.io/quote v1.5.2

When you ran go mod tidy, it located and downloaded the rsc.io/quote module that contains the package you imported. By default, it downloaded the latest version -- v1.5.2.

# Run your code

$ go run .
$ go run hello.go
$ /.hello

# Note:

# Packages Convention
main package is the starting point to run the program
func main() - main function where the program execution begins

Go programs are made up of packages
All Go source is part of a package.
Every file begins with a package statement.
Programs start in package main.

## package main
```

import "fmt"

func main() {
fmt.Println("Hello, world!")
}
```

For very small programs, main is the only package you need to write.

The hello world program imports package fmt.

The function Println is defined in the fmt package.

## An example package: fmt
```
// Package fmt implements formatted I/O.
package fmt

// Println formats using the default formats for its
// operands and writes to standard output.
func Println(a ...interface{}) (n int, err error) {
...
}

func newPrinter() *pp {
...
}
```

The Println function is exported. It starts with an upper case
letter, which means other packages are allowed to call it.

The newPrinter function is unexported. It starts with a lower
case letter, so it can only be used inside the fmt package.

## The shape of a package
Packages collect related code.

They can be big or small,
and may be spread across multiple files.

All the files in a package live in a single directory.

The net/http package exports more than 100 names. (18 files)
The errors package exports just one. (1 file)

## The name of a package
Keep package names short and meaningful.
Don't use underscores, they make package names long.

** io/ioutil not io/util
** suffixarray not suffix_array
Don't overgeneralize. A util package could be anything.

The name of a package is part of its type and function names.
On its own, type Buffer is ambiguous. But users see:

```buf := new(bytes.Buffer) ```
Choose package names carefully.

Choose good names for users.

## The testing of a package
Tests are distinguished by file name. Test files end in _test.go.

```package fmt

import "testing"

var fmtTests = []fmtTest{
{"%d", 12345, "12345"},
{"%v", 12345, "12345"},
{"%t", true, "true"},
}

func TestSprintf(t *testing.T) {
for _, tt := range fmtTests {
if s := Sprintf(tt.fmt, tt.val); s != tt.out {
t.Errorf("...")
}
}
}
```

Test well.


# Code organization
