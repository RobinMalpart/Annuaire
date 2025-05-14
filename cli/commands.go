package cli

import (
	"annuaire/contact"
	"flag"
	"fmt"
	"os"
)

func ParseAndExecute(a contact.Annuaire) {
	add := flag.Bool("add", false, "Ajouter un contact")
	delete := flag.Bool("delete", false, "Supprimer un contact")
	find := flag.Bool("find", false, "Rechercher un contact")
	update := flag.Bool("update", false, "Mettre à jour un contact")
	list := flag.Bool("list", false, "Lister tous les contacts")

	first := flag.String("first", "", "Prénom du contact")
	last := flag.String("last", "", "Nom du contact")
	phone := flag.String("phone", "", "Téléphone du contact")
	term := flag.String("term", "", "Nom ou prénom pour recherche/suppression")

	flag.Parse()

	switch {
	case *add:
		if *first == "" || *last == "" || *phone == "" {
			fmt.Println("Erreur : -first, -last et -phone sont requis avec -add")
			os.Exit(1)
		}
		err := a.AddContact(contact.Contact{FirstName: *first, LastName: *last, Phone: *phone})
		if err != nil {
			fmt.Println(err)
		}
	case *delete:
		if *term == "" {
			fmt.Println("Erreur : -term requis avec -delete")
			os.Exit(1)
		}
		a.DeleteContact(*term)
	case *find:
		if *term == "" {
			fmt.Println("Erreur : -term requis avec -find")
			os.Exit(1)
		}
		results := a.FindContact(*term)
		for _, c := range results {
			fmt.Printf("%s %s - %s\n", c.FirstName, c.LastName, c.Phone)
		}
	case *update:
		if *first == "" || *last == "" || *phone == "" {
			fmt.Println("Erreur : -first, -last et -phone sont requis avec -update")
			os.Exit(1)
		}
		ok := a.UpdateContact(contact.Contact{FirstName: *first, LastName: *last, Phone: *phone})
		if ok {
			fmt.Println("Contact mis à jour.")
		} else {
			fmt.Println("Contact introuvable.")
		}
	case *list:
		a.GetAllContact()
	default:
		fmt.Println("Utilisation :")
		fmt.Println("  -add -first=John -last=Doe -phone=0600000000")
		fmt.Println("  -delete -term=Doe")
		fmt.Println("  -find -term=John")
		fmt.Println("  -update -first=John -last=Doe -phone=0600000000")
		fmt.Println("  -list")
	}
}
