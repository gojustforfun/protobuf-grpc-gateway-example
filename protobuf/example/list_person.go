package example

import (
	"fmt"
	"io"
)

func ListPeople(w io.Writer, book *AddressBook) {
	for _, person := range book.People {
		writePerson(w, person)
	}
}

func writePerson(w io.Writer, p *Person) {
	fmt.Fprintln(w, "Person ID:", p.Id)
	fmt.Fprintln(w, " Name:", p.Name)
	if p.Email != "" {
		fmt.Fprintln(w, " E-mail address:", p.Email)
	}

	for _, pn := range p.Phones {
		switch pn.Type {
		case Person_MOBILE:
			fmt.Fprint(w, " Mobile phone #: ")
		case Person_HOME:
			fmt.Fprint(w, " Home phone #: ")
		case Person_WORK:
			fmt.Fprint(w, " Work phone #: ")
		}
		fmt.Fprintln(w, pn.Number)
	}
}
