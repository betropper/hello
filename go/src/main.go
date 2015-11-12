package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	"os"
)

func plain() {
	fmt.Println(c.CL + c.R + "Hello, world!")
}

func color(message string) {
	fmt.Println(c.CL + c.R + message)
}

func main() {
	if len(os.Args) > 1 {
		color("Hello, " + os.Args[2] + "!")
	} else {
		plain()
	}
}
