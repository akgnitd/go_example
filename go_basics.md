# Note:

main package is the starting point to run the program
func main() - main function where the program execution begins here

# Keywords

## reserved words in Go & not to be used as constant or variable or any other identifier names
break	default	func	interface	select
case	defer	Go	map	Struct
chan	else	Goto	package	Switch
const	fallthrough	if	range	Type
continue	for	import	return	

# Variables

#### var a = "Hello World"
#### var a string = "Hello"
#### var m,n string = "Hello", "World"
#### var t string
#### t = "Hello World"
#### u := "Hello World"

### Which style should be used?

There is some constraint enforced by the language here
not possible to use short variable declaration outside a function, but otherwise, anything goes.

# Setting up the development environment
### Create a root folder for our workspace

mkdir ~/goworkspace
cd goworkspace
export GOPATH=`pwd`


### Create folders and files

mkdir -p src/github.com/callistaenterprise
cd src/github.com/callistaenterprise
mkdir -p goblog/accountservice
cd goblog/accountservice
touch main.go
mkdir service

# Root Folder - $GOPATH/src/github.com/callistaenterprise/goblog



