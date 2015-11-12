package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	"os"
	"strings"
)

func plain() {

	fmt.Println(c.CL + c.Rc() + "Hello, world!")
}

func color(message string) {
	fmt.Println(c.CL + c.Rc() + message)
}

func multi(message string) {
	fmt.Println(c.CL + c.Multi(message))
}

func main() {

	message := []string{}
	var option = ""

	if len(os.Args) > 1 {
		if strings.HasPrefix(os.Args[1], "-") {
			option = os.Args[1]
			message = os.Args[2:]
		} else {
			message = os.Args[1:]
		}
	} else {
		plain()
		exit()
	}

	if option == "-c" {
		color(message)
	} else if option == "-m" {
		multi(message)
	}
}
