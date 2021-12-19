package main

import (
	"fmt"

	"custom-datatypes/organization"
)

func main() {
	fmt.Println("hello world")
	p := organization.NewPerson("aadi", "saxena")
	err := p.SetTwitterHandle("@aadi")

	fmt.Println(p.ID())
	fmt.Println(p.Fullname())

	if(err != nil){
		fmt.Println(err)
	}

	fmt.Println(p.TwitterHandle())
}