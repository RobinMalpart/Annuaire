package main

import (
	"annuaire/cli"
	"annuaire/contact"
)

func main() {
	ann := contact.Annuaire{
		contact.BuildKey(contact.Contact{FirstName: "John", LastName: "Doe"}): {
			FirstName: "John",
			LastName:  "Doe",
			Phone:     "555-0123",
		},
		contact.BuildKey(contact.Contact{FirstName: "Jane", LastName: "Smith"}): {
			FirstName: "Jane",
			LastName:  "Smith",
			Phone:     "555-0456",
		},
		contact.BuildKey(contact.Contact{FirstName: "Alice", LastName: "Dupont"}): {
			FirstName: "Alice",
			LastName:  "Dupont",
			Phone:     "",
		},
	}

	cli.ParseAndExecute(ann)
}
