package organization

import (
	"errors"
	"strings"
)

type Identifiable interface {
	ID() string
}

type Person struct {
	firstName string
	lastName string
	twitterHandle string
}

func NewPerson (FirstName, LastName string) Person {
	return Person{
		firstName: FirstName,
		lastName: LastName,
	}
}

func (p *Person) SetTwitterHandle(handle string) error {
	if(!strings.HasPrefix(handle, "@")) {
		return errors.New("twitter handle must start with an @ character")
	}

	p.twitterHandle = handle
	return nil
}

func (p *Person) TwitterHandle() string {
	return p.twitterHandle
}

func (p *Person) Fullname() string {
	return p.firstName + p.lastName
}

func (p *Person) ID() string {
	return "1234"
}