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
					Phone:     "555-0000",
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

func TestUpdateContact(t *testing.T) {
	type testCase struct {
		initial  contact.Annuaire
		input    contact.Contact
		expected bool
		finalNum string
	}

	tests := map[string]testCase{
		"UpdateExisting": {
			initial: contact.Annuaire{
				contact.BuildKey(contact.Contact{FirstName: "John", LastName: "Doe"}): {
					FirstName: "John",
					LastName:  "Doe",
					Phone:     "555-0000",
				},
			},
			input: contact.Contact{
				FirstName: "John",
				LastName:  "Doe",
				Phone:     "555-9999",
			},
			expected: true,
			finalNum: "555-9999",
		},
		"UpdateNonExisting": {
			initial: contact.Annuaire{
				contact.BuildKey(contact.Contact{FirstName: "John", LastName: "Doe"}): {
					FirstName: "John",
					LastName:  "Doe",
					Phone:     "555-0000",
				},
			},
			input: contact.Contact{
				FirstName: "Jane",
				LastName:  "Smith",
				Phone:     "555-8888",
			},
			expected: false,
			finalNum: "",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ann := test.initial
			result := (&ann).UpdateContact(test.input)

			if result != test.expected {
				t.Errorf("expected = %t, got = %t", test.expected, result)
			}

			if test.finalNum != "" {
				key := contact.BuildKey(test.input)
				updatedContact, exists := ann[key]
				if !exists {
					t.Errorf("expected contact to exist with key %s", key)
				} else if updatedContact.Phone != test.finalNum {
					t.Errorf("expected phone = %s, got = %s", test.finalNum, updatedContact.Phone)
				}
			}
		})
	}
}
