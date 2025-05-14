package contact

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

func (a Annuaire) GetAllContact(c Contact) {

}

func (a Annuaire) CheckContact(c Contact) Contact {
	key := c.FirstName + "_" + c.LastName
	if existingContact, exists := a[key]; exists {
		return existingContact
	}
	a[key] = c
	return c
}
