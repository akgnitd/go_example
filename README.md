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

main package is the starting point to run the program
func main() - main function where the program execution begins

