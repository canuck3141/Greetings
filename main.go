package main

import (
	"fmt"

	"github.com/canuck3141/greetings/bye"
	"github.com/canuck3141/greetings/hello"
)

func main() {
	fmt.Printf("hello: %s\n", hello.Greet())
	fmt.Printf("bye: %s\n", bye.Greet())
}
