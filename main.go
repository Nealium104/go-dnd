package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

var characterDirectory string

type Character struct {
	Name   string `json:"name"`
	Id     string `json:"id"`
	Health int    `json:"health"`
}

func main() {
	characterDirectory = "./assets/characters/"
	c := gatherCLICharacter()

	c.writeCharacter(characterDirectory)
	displayCharacter(c)

	allCharacters := allCharacters()
	for _, c := range allCharacters {
		fmt.Println(c.Name)
	}

	character := loadCharacter(c.Id)
	fmt.Printf("loadCharacter Function Output: %v\n", character.Name)
}

func gatherCLICharacter() Character {
	var c Character
	var n string
	var h int

	fmt.Println("What is this character's name?")
	fmt.Scan(&n)

	fmt.Println("What is the base health of this character?")
	fmt.Scan(&h)

	c.Name = n
	c.Id = uuid.New().String()
	c.Health = h
	return c
}

func (c *Character) writeCharacter(cDirectory string) {
	jsonData, err := json.Marshal(c)
	if err != nil {
		log.Fatal("Error marshaling json: ", err)
	}
	os.WriteFile(fmt.Sprintf("%s%s.json", cDirectory, c.Id), jsonData, 0764)
}

func displayCharacter(c Character) {
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Health: %d\n", c.Health)
}

func loadCharacter(id string) Character {
	cFile, err := os.Open(fmt.Sprintf("%s%s.json", characterDirectory, id))
	if err != nil {
		log.Fatal("Error opening loadCharacter file: ", err)
	}
	defer cFile.Close()

	byteValue, err := io.ReadAll(cFile)
	if err != nil {
		log.Fatal("Err reading loadCharacter file: ", err)
	}

	var c Character
	marshalErr := json.Unmarshal(byteValue, &c)
	if marshalErr != nil {
		log.Fatal("Error unmarshaling in loadCharacter: ", err)
	}
	return c
}

func allCharacters() []Character {
	characterSlice := make([]Character, 0)
	entries, err := os.ReadDir(characterDirectory)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		character := Character{}
		if entry.IsDir() {
			continue
		}

		filePath := filepath.Join(characterDirectory, entry.Name())

		data, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(data, &character)
		if err != nil {
			log.Fatal(err)
		}

		c := Character{Name: character.Name, Id: character.Id, Health: character.Health}
		characterSlice = append(characterSlice, c)
	}
	return characterSlice
}
