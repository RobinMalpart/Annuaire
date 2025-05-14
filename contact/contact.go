package contact

import "fmt"

type Contact struct {
	FirstName string
	LastName  string
	Phone     string
}
type Annuaire map[string]Contact

func (a Annuaire) AddContact(c Contact) {}

func (a Annuaire) DeleteContact(c Contact) {}

func (a Annuaire) FindContact(c Contact) {}

func (a Annuaire) UpdateContact(c Contact) {}

func (a Annuaire) GetAllContact() {

}

func (a Annuaire) CheckContact(c Contact) bool {
	key := c.FirstName + "_" + c.LastName
	if existingContact, exists := a[key]; exists {
		fmt.Printf("Contact %q, %q already exist\n", existingContact.FirstName, existingContact.LastName)
		return true
	}
	return false
}
