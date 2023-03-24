package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	var reply Card
	var db []Card

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Error Dialing:", err)
	}

	avacyn := Card{"Avacyn, Angel of Hope", "White", "Mythic", "Avacyn Restored", [2]uint8{8, 8}, [6]uint8{3, 0, 0, 0, 0, 5}, 54.54, 1}
	bigblob := Card{"Big Blob", "Black", "Uncommon", "Phyrex", [2]uint8{3, 3}, [6]uint8{0, 0, 2, 0, 0, 2}, 5.00, 1}
	crazedgorilla := Card{"Crazed Gorilla", "Green", "Common", "blahblah", [2]uint8{6, 5}, [6]uint8{5, 0, 0, 3, 0, 0}, 0.45, 1}
	monsterhunter := Card{"Monster Hunter", "Red", "Mythic", "foofoobarbar", [2]uint8{4, 2}, [6]uint8{1, 0, 0, 0, 2, 1}, 6.74, 1}

	err = client.Call("API.AddCard", avacyn, &reply)
	if err != nil {
		fmt.Printf("Error Adding Card", err)
	}
	err = client.Call("API.AddCard", bigblob, &reply)
	if err != nil {
		fmt.Printf("Error Adding Card", err)
	}
	err = client.Call("API.AddCard", crazedgorilla, &reply)
	if err != nil {
		fmt.Printf("Error Adding Card", err)
	}
	err = client.Call("API.AddCard", monsterhunter, &reply)
	if err != nil {
		fmt.Printf("Error Adding Card", err)
	}
	err = client.Call("API.GetDB", "", &db)
	if err != nil {
		fmt.Printf("Error Getting DB", err)
	}

	var temp Card
	client.Call("API.GetByName", "Avacyn, Angel of Hope", &temp)
	fmt.Println("This is is in our database: ", db)
	fmt.Println("This is the value of our getter ", temp)
}

type Card struct {
	Name   string
	Color  string
	Rarity string
	Set    string
	Stats  [2]uint8
	Cost   [6]uint8
	Price  float32
	Stock  uint8
}
