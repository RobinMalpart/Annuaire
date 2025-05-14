package contact

import (
	"fmt"
	"strings"
)

type Contact struct {
	FirstName string
	LastName  string
	Phone     string
}

func BuildKey(c Contact) string {
	return strings.ToLower(c.FirstName + "_" + c.LastName)
}

type Annuaire map[string]Contact

func (a Annuaire) AddContact(c Contact) error {
	if a.CheckContact(c) {
		key := c.FirstName + "_" + c.LastName
		return fmt.Errorf("Contact %s already exists", key)
	}

	key := c.FirstName + "_" + c.LastName
	a[key] = c

	fmt.Printf("Contact %s added successfully\n", key)

	return nil
}

func (a Annuaire) DeleteContact(input string) bool {
	for key, contact := range a {
		if contact.FirstName == input || contact.LastName == input {
			delete(a, key)
			fmt.Printf("Contact %s deleted successfully\n", key)
			return true
		}
	}
	fmt.Printf("Contact %s not found\n", input)
	return false
}

func (a Annuaire) FindContact(term string) []Contact {
	var results []Contact

	for _, contact := range a {
		if contact.FirstName == term || contact.LastName == term {
			results = append(results, contact)
		}
	}
	return results
}

func (a *Annuaire) UpdateContact(c Contact) bool {
	key := BuildKey(c)
	if _, exists := (*a)[key]; exists {
		(*a)[key] = c
		return true
	}
	return false
}

func (a Annuaire) GetAllContact() {
	if len(a) == 0 {
		fmt.Println("L'annuaire est vide.")
		return
	}

	for _, c := range a {
		fmt.Printf("%s %s - %s\n", c.FirstName, c.LastName, c.Phone)
	}
}

func (a Annuaire) CheckContact(c Contact) bool {
	key := strings.ToLower(c.FirstName) + "_" + strings.ToLower(c.LastName)
	_, exists := a[key]
	return exists
}
