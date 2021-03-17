package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Book struct {
	SBN    int    `json:"SBN"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}
type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Library struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
	Open    bool    `json:"open,omitempty"`
	Books   []Book  `json:"books,omitempty"` //with ,omitempty  I wont have the books key if it is null/empty when parsing to json
}
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func fromBookToJson(book Book) (jsonString string) {
	fmt.Printf("I've this book obj: %+v\n", book)

	byteArray, err := json.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}
	jsonString = string(byteArray)
	fmt.Println("I get this json string: " + jsonString)
	return

}
func fromJsonToAddress(jsonString string) (address Address) {

	fmt.Println("I've this address json string: " + jsonString)

	err := json.Unmarshal([]byte(jsonString), &address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("I get this address obj: %+v\n", address)
	return
}

func fromLibraryToJson(library Library) (jsonString string) {
	fmt.Printf("I've this library obj: %+v\n", library)

	byteArray, err := json.Marshal(library)
	if err != nil {
		log.Fatal(err)
	}
	jsonString = string(byteArray)
	fmt.Println("I get this json string: " + jsonString)
	return
}
func playingWithJson() {
	fmt.Printf("\n\n#####\tjson tutorial\t#####\n\n")
	authorInput := Author{Name: "Mr. Buzi", Age: 8}
	bookInput := Book{SBN: 42, Title: "cats of the world", Author: authorInput}
	book2 := Book{SBN: 9000, Title: "street life", Author: authorInput}

	jsonStringOutput1 := fromBookToJson(bookInput)
	fmt.Println(" - Output1 result:\n" + jsonStringOutput1)
	jsonStringInput1 := `{"street": "Corso Duca degli Abruzzi", "city" : "Torino", "country": "it"}`

	addressOutput := fromJsonToAddress(jsonStringInput1)
	var books []Book = []Book{bookInput, book2}
	library1 := Library{Name: "PoliTO Library", Address: addressOutput, Open: true, Books: books}
	jsonStringOutput2 := fromLibraryToJson(library1)
	fmt.Println(" - Output2 result:\n" + jsonStringOutput2)

	library2 := Library{Name: "UniTO Library", Address: addressOutput}
	fmt.Println(" - Output3 result:\n" + fromLibraryToJson(library2))
}
