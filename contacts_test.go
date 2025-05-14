package main

import "annuaire/contact"
import "testing"

func TestCheckContact(t *testing.T) {

	type testCase struct {
		contacts contact.Annuaire
		contact  contact.Contact
		want     bool
	}
	var tests = map[string]testCase{
		"FilledAndContactExist": {
			contacts: map[string]contact.Contact{
				"john_doe": {
					FirstName: "John",
					LastName:  "Doe",
					Phone:     "555-0123",
				},
			},
			contact: contact.Contact{
				FirstName: "John",
				LastName:  "Doe",
				Phone:     "555-0123",
			},
			want: true,
		},
		"FilledAndContactNotExist": {
			contacts: map[string]contact.Contact{
				"john_doe": {
					FirstName: "John",
					LastName:  "Doe",
					Phone:     "555-0123",
				},
			},
			contact: contact.Contact{
				FirstName: "Jane",
				LastName:  "Smith",
				Phone:     "555-0456",
			},
			want: false,
		},
		"Empty": {
			contacts: make(map[string]contact.Contact),
			contact: contact.Contact{
				FirstName: "Jane",
				LastName:  "Smith",
				Phone:     "555-0456",
			},
			want: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := test.contacts.CheckContact(test.contact)
			if test.want != got {
				t.Errorf("expected = %t, got %t", test.want, got)
			}
		})
	}
}
