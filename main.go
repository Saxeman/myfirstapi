package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	var api = new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("Error Registering API", err)
	}

	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("listen error: ", err)
	}

	log.Printf("serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error Serving: ", err)
	}

	/*fmt.Println("Current Database: ", db)

	var deleted = DeleteCard(avacyn)
	fmt.Println("Deleted Card: ", deleted)
	fmt.Println("Updated Database: ", db)

	updated_bigblob := Card{"Big Blob", "Black", "Uncommon", "Phyrex", [2]uint8{5, 5}, [6]uint8{0, 0, 2, 0, 0, 2}, 10.00, 1}

	var changed = EditCard(bigblob, updated_bigblob)
	fmt.Println("Updated Card: ", changed)
	fmt.Println("Updated Database: ", db)*/
}

type API int

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

func (a *API) GetDB(name string, reply *[]Card) error {
	*reply = db
	return nil
}

func (a *API) GetByName(name string, reply *Card) error {
	var GetItem Card

	// TODO: POSSIBLY NEED TO ADD IN FUNCTIONALITY FOR DUPLICATES (FOILS, REPRINTS, ETC) BUT OK FOR NOW
	for _, val := range db {
		if val.Name == name {
			GetItem = val
		}
	}

	*reply = GetItem

	// TODO: ADD IN ERROR HANDLING
	return nil
}

func (a *API) AddCard(card Card, reply *Card) error {
	db = append(db, card)
	fmt.Println(db)
	fmt.Println(card)
	*reply = card
	// TODO: ADD IN ERROR HANDLING
	return nil
}

func (a *API) EditCard(edit Card, reply *Card) error {
	var changed Card

	for idx, val := range db {
		if val.Name == edit.Name && val.Set == edit.Set {
			db[idx] = edit
			changed = edit
		}
	}
	*reply = changed

	// TODO: ADD IN ERROR HANDLING
	return nil
}

func (a *API) DeleteCard(card Card, reply *Card) error {
	var Del Card

	for idx, val := range db {
		if val.Name == card.Name && val.Set == card.Set {
			db = append(db[:idx], db[idx+1:]...)
			Del = val
			break
		}
	}
	*reply = Del
	// TODO: ADD IN ERROR HANDLING
	return nil
}
