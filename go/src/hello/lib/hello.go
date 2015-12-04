package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	"strings"
	"time"
)

func plain(message string) {

	fmt.Println(c.Clear + message)
}

func color(message string) {
	fmt.Println(c.Clear + c.Rc() + message)
}

func multi(message string) {
	for {
		fmt.Println(c.Clear + c.Multi(message))
	}
}

func nyan(message string) {
	fmt.Print(c.Clear)
	for {
		fmt.Print(c.Rc() + message + " ")
	}
}

func ayylmao() {
	fmt.Println(c.Clear + c.Red + "Ayy..")
	time.Sleep(1000 * time.Millisecond)
	for {
		fmt.Print(c.Rc() + "LMAO" + " ")
	}
}
