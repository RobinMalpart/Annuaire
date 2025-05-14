package Annuaire

import "annuaire/contact"

func main() {

	var contacts contact.Annuaire = map[string]contact.Contact{
		"john_doe": contact.Contact{
			FirstName: "John",
			LastName:  "Doe",
			Phone:     "555-0123",
		},
		"jane_smith": contact.Contact{
			FirstName: "Jane",
			LastName:  "Smith",
			Phone:     "555-0456",
		},
	}

}
