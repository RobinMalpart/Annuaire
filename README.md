# Groupe 7
Ranyl
Rosvalde
Robin

# Annuaire CLI en Go

Ce projet est une application en ligne de commande (CLI) écrite en Go.  
Elle permet de gérer un annuaire de contacts avec les fonctionnalités suivantes :

- Ajouter un contact
- Supprimer un contact
- Rechercher un contact par nom ou prénom
- Mettre à jour un contact
- Lister tous les contacts

## Pré-requis

- Go installé : https://go.dev/dl
- Un terminal

## Lancer le projet

Se placer dans le dossier racine du projet, puis exécuter :

```bash
go run main.go [options]
```

## Commandes disponibles

| Commande                                                                 | Description                                  |
|--------------------------------------------------------------------------|----------------------------------------------|
| `-add -first=Prénom -last=Nom -phone=Numéro`                             | Ajoute un contact                            |
| `-find -term=NomOuPrénom`                                               | Recherche un contact par nom ou prénom       |
| `-update -first=Prénom -last=Nom -phone=NouveauNuméro`                  | Met à jour un contact                        |
| `-delete -term=NomOuPrénom`                                             | Supprime un contact                          |
| `-list`                                                                 | Affiche tous les contacts enregistrés        |

## Exemples d'utilisation

```bash
go run main.go -add -first=Jean -last=Michel -phone=0600000000
go run main.go -find -term=Doe
go run main.go -list
go run main.go -update -first=Jean -last=Michel -phone=0666666666
go run main.go -delete -term=Jean
```

## Structure du projet

- `main.go` : point d’entrée, initialise des contacts de test
- `contact/` : logique métier (struct Contact, Annuaire, méthodes)
- `cli/commands.go` : gestion des commandes et des flags CLI

## Remarques

À chaque exécution, l’annuaire est réinitialisé avec 3 contacts de test.  
Aucune persistance n’est encore implémentée (les données sont perdues à chaque relance).

