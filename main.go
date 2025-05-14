package main

import (
	"annuaire/contact"
	"fmt"
)

func main() {
	var contacts contact.Annuaire = map[string]contact.Contact{
		"john_doe": {
			FirstName: "John",
			LastName:  "Doe",
			Phone:     "555-0123",
		},
		"jane_smith": {
			FirstName: "Jane",
			LastName:  "Smith",
			Phone:     "555-0456",
		},
		"alice_dupont": {
			FirstName: "Alice",
			LastName:  "Dupont",
			Phone:     "",
		},
	}

	results := contacts.FindContact("Jane")
	for _, c := range results {
		fmt.Printf("%s %s - %s\n", c.FirstName, c.LastName, c.Phone)
	}

	contacts.UpdateContact(contact.Contact{
		FirstName: "Alice",
		LastName:  "Dupont",
		Phone:     "555-0999",
	})

	for _, c := range contacts {
		fmt.Printf("%s %s - %s\n", c.FirstName, c.LastName, c.Phone)
	}
	// Add a new contact
	newContact := contact.Contact{
		FirstName: "Alice",
		LastName:  "Johnson",
		Phone:     "555-0789",
	}

	err := contacts.AddContact(newContact)
	if err != nil {
		println(err.Error())
	} else {
		println("Contact added successfully")
	}

}
