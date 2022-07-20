package main

import "fmt"
import "rsc.io/quote"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Println(fmt.Sprintf("Hi, %v. Welcome!", Hello("Atul")))
	fmt.Println(Hello("Atul"))
}

func Hello(name string) string {
	return name
}
