package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Identifiable interface {
	ID() string
}

type Name struct {
	first string
	last string
}

type Person struct {
	Name
	twitterHandle TwitterHandle
}

func NewPerson (FirstName, LastName string) Person {
	return Person{
		Name: Name{
			first: FirstName,
			last: LastName,
		},
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

func (n *Name) Fullname() string {
	return n.first + n.last
}

func (p *Person) ID() string {
	return "1234"
}