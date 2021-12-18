package main

import (
	"fmt"

	"custom-datatypes/organization"
)

func main() {
	fmt.Println("hello world")
	var p = organization.Person{FirstName: "aadi", LastName: "saxena"}
	fmt.Println(p.ID())
	fmt.Println(p.FirstName)
	fmt.Println(p.LastName)
}