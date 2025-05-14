package contact

type Contact struct {
	FirstName string
	LastName  string
	Phone     string
}

type Annuaire struct {
	Contacts []Contact
}

func (a Annuaire) AddContact(c Contact) {}

func (a Annuaire) DeleteContact(c Contact) {}

func (a Annuaire) FindContact(c Contact) {}

func (a Annuaire) UpdateContact(c Contact) {}

func (a Annuaire) GetAllContact(c Contact) {}

func (a Annuaire) CheckContact(c Contact) {}
