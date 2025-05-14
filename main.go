package Annuaire

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

	fmt.Println("Recherche de 'Jane':")
	results := contacts.FindContact("Jane")
	for _, c := range results {
		fmt.Printf("- %s %s (%s)\n", c.FirstName, c.LastName, c.Phone)
	}

	fmt.Println("Mise à jour du numéro de Alice Dupont...")
	ok := contacts.UpdateContact(contact.Contact{
		FirstName: "Alice",
		LastName:  "Dupont",
		Phone:     "555-0999",
	})
	if ok {
		fmt.Println("Mise à jour réussie.")
	} else {
		fmt.Println("Contact introuvable.")
	}

	fmt.Println("Annuaire final:")
	for _, c := range contacts {
		fmt.Printf("- %s %s (%s)\n", c.FirstName, c.LastName, c.Phone)
	}
}
