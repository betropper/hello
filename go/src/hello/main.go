package main

import (
	"fmt"
	h "hello/lib"
	"os"
	"strings"
)

func main() {
	message := "Hello, world!"
	var option = ""

	if len(os.Args) > 1 {
		if strings.HasPrefix(os.Args[1], "-") {
			option = os.Args[1]
			if len(os.Args) > 2 {
				message = strings.Join(os.Args[2:], " ")
			}
		} else {
			message = strings.Join(os.Args[1:], " ")
		}
	}
	if strings.Contains(message, "ayy lmao") {
		h.ayylmao()
	} else {
		plain(message)
		os.Exit(0)
	}

	if option == "-c" {
		h.color(message)
	} else if option == "-m" {
		h.multi(message)
	} else if option == "-n" {
		h.nyan(message)
	} else {
		h.plain(message)
	}
}
