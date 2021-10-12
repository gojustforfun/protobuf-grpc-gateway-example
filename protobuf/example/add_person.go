package example

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func PromptForAddress(r io.Reader) (*Person, error) {

	rd := bufio.NewReader(r)
	p := &Person{}

	fmt.Print("Enter person ID number: ")
	if _, err := fmt.Fscanf(rd, "%d\n", &p.Id); err != nil {
		return p, err
	}

	fmt.Print("Enter name: ")
	name, err := rd.ReadString('\n')
	if err != nil {
		return p, err
	}
	p.Name = strings.TrimSpace(name)

	fmt.Print("Enter email address (blank for none): ")
	email, err := rd.ReadString('\n')
	if err != nil {
		return p, err
	}
	p.Email = strings.TrimSpace(email)

	for {
		fmt.Print("Enter a phone number (or leave blank to finish): ")
		phone, err := rd.ReadString('\n')
		if err != nil {
			return p, err
		}
		phone = strings.TrimSpace(phone)
		if phone == "" {
			break
		}

		pn := &Person_PhoneNumber{Number: phone}

		fmt.Print("Is this a mobile, home, or work phone? ")
		pType, err := rd.ReadString('\n')
		if err != nil {
			return p, err
		}
		pType = strings.TrimSpace(pType)
		switch pType {
		case "mobile":
			pn.Type = Person_MOBILE
		case "home":
			pn.Type = Person_HOME
		case "work":
			pn.Type = Person_WORK
		default:
			fmt.Printf("Unknown phone type %q.  Using default.\n", pType)
		}
		p.Phones = append(p.Phones, pn)
	}
	return p, nil
}
