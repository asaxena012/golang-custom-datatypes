package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Identifiable interface {
	ID() string
}

type Person struct {
	firstName string
	lastName string
	twitterHandle TwitterHandle
}

func NewPerson (FirstName, LastName string) Person {
	return Person{
		firstName: FirstName,
		lastName: LastName,
	}
}

type TwitterHandle string

func (th TwitterHandle) RedirectURL() string {
	cleanHandle := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandle)
}

func (p *Person) SetTwitterHandle(handle TwitterHandle) error {
	if(!strings.HasPrefix(string(handle), "@")) {
		return errors.New("twitter handle must start with an @ character")
	}

	p.twitterHandle = handle
	return nil
}

func (p *Person) TwitterHandle() TwitterHandle {
	return p.twitterHandle
}

func (p *Person) Fullname() string {
	return p.firstName + p.lastName
}

func (p *Person) ID() string {
	return "1234"
}