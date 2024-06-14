package kata

import (
	"sort"
	"strings"
)

// Meeting / https://www.codewars.com/kata/59df2f8f08c6cec835000012/train/go
func Meeting(s string) string {
	persons := ParsePersons(s)
	return persons.Sorted().String()
}

type Person struct {
	FirstName string
	LastName  string
}

func NewPerson(name string) Person {
	names := strings.Split(name, ":")

	return Person{
		FirstName: names[0],
		LastName:  names[1],
	}
}

func (p Person) String() string {
	person := "(" + p.LastName + ", " + p.FirstName + ")"
	return strings.ToUpper(person)
}

func (p Person) Compare(other Person) int {
	myLastName := strings.ToLower(p.LastName)
	otherLastName := strings.ToLower(other.LastName)

	lastNameCompare := strings.Compare(myLastName, otherLastName)

	if lastNameCompare == 0 {
		myFirstName := strings.ToLower(p.FirstName)
		otherFirstName := strings.ToLower(other.FirstName)

		return strings.Compare(myFirstName, otherFirstName)
	}

	return lastNameCompare
}

func (p Person) IsBefore(other Person) bool {
	return p.Compare(other) < 0
}

type Persons []Person

func ParsePersons(s string) Persons {
	names := strings.Split(s, ";")
	persons := make([]Person, len(names))

	for i, name := range names {
		persons[i] = NewPerson(name)
	}

	return persons
}

// Sorted / sorts persons without modifying the original slice
func (p Persons) Sorted() Persons {
	result := make(Persons, len(p))
	copy(result, p)

	sort.Slice(result, func(i, j int) bool {
		personA := result[i]
		personB := result[j]

		return personA.IsBefore(personB)
	})

	return result
}

func (p Persons) String() string {
	result := make([]string, len(p))

	for i, person := range p {
		result[i] = person.String()
	}

	return strings.Join(result, "")
}
