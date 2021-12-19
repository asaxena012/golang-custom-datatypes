package main

import (
	"fmt"

	"custom-datatypes/organization"
)

func main() {
	fmt.Println("hello world")
	p := organization.NewPerson("aadi", "saxena")
	fmt.Println(p.ID())
	fmt.Println(p.Fullname())
}