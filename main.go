package main

import (
	"fmt"
	"dog"
	// "github.com/place-your-name-on-github-here/myGoDogDoc"
)

type doggy struct {
	name  string
	age int
}

func main() {
	fido := doggy{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}
