package main

import (
	"fmt"
)

// This defines the card struct, a representation of a Magic the Gathering card which will be stored in our database.
// The information stored in the struct are as followed: Color, Cost, Rarity, and stock (how much I personally own)
// the plan for this API is to be able to build our backend as a microservice which can simply be maintained and operated
// by Go, so that any issues do not affect the rest of our platform.
type Card struct {
	// TODO: UPDATE TO SQL
	// TODO: NEED TO ACCOUNT FOR FOIL VERSIONS
	Name string
	// TODO:possibly change color and rarity to an enum (int or string?) since options are limited.
	Color  string
	Rarity string
	Set    string
	Stats  [2]uint8
	// cost can only be a value 0 - 5 for each color: white, blue, black, red, green, and colorless
	Cost  [6]uint8
	Price float32
	Stock uint8
}

var db []Card

/* TODO: ADD MORE FUNCTIONALITY FOR MORE DATA ACCESS
 SUCH AS:
		SPECIFIC GETTER FUNCTIONS
		CHANGE INDIVIDUAL STATS INSTEAD OF NEEDING TO REMAKE THE CARD EACH TIME
		ETC.
*/

func GetByName(name string) Card {
	var GetItem Card

	for _, val := range db {
		if val.Name == name {
			GetItem = val
		}
	}
	return GetItem
}

func AddCard(card Card) Card {
	db = append(db, card)
	return card
}

func EditCard(card Card, edit Card) Card {
	var changed Card

	for idx, val := range db {
		if val.Name == card.Name && val.Set == card.Set {
			db[idx] = edit
			changed = edit
		}
	}
	return changed
}

func DeleteCard(card Card) Card {
	var Del Card

	for idx, val := range db {
		if val.Name == card.Name && val.Set == card.Set {
			db = append(db[:idx], db[idx+1:]...)
			Del = val
			break
		}
	}
	return Del
}

func main() {
	avacyn := Card{"Avacyn, Angel of Hope", "White", "Mythic", "Avacyn Restored", [2]uint8{8, 8}, [6]uint8{3, 0, 0, 0, 0, 5}, 54.54, 1}
	bigblob := Card{"Big Blob", "Black", "Uncommon", "Phyrex", [2]uint8{3, 3}, [6]uint8{0, 0, 2, 0, 0, 2}, 5.00, 1}
	crazedgorilla := Card{"Crazed Gorilla", "Green", "Common", "blahblah", [2]uint8{6, 5}, [6]uint8{5, 0, 0, 3, 0, 0}, 0.45, 1}
	monsterhunter := Card{"Monster Hunter", "Red", "Mythic", "foofoobarbar", [2]uint8{4, 2}, [6]uint8{1, 0, 0, 0, 2, 1}, 6.74, 1}

	AddCard(avacyn)
	AddCard(bigblob)
	AddCard(crazedgorilla)
	AddCard(monsterhunter)

	fmt.Println("Current Database: ", db)

	var deleted = DeleteCard(avacyn)
	fmt.Println("Deleted Card: ", deleted)
	fmt.Println("Updated Database: ", db)

	updated_bigblob := Card{"Big Blob", "Black", "Uncommon", "Phyrex", [2]uint8{5, 5}, [6]uint8{0, 0, 2, 0, 0, 2}, 10.00, 1}

	var changed = EditCard(bigblob, updated_bigblob)
	fmt.Println("Updated Card: ", changed)
	fmt.Println("Updated Database: ", db)
}
