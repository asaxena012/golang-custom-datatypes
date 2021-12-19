package organization

type Identifiable interface {
	ID() string
}

type Person struct {
	firstName string
	lastName string
}

func NewPerson (FirstName, LastName string) Person {
	return Person{
		firstName: FirstName,
		lastName: LastName,
	}
}

func (p Person) Fullname() string {
	return p.firstName + p.lastName
}

func (p Person) ID() string {
	return "1234"
}