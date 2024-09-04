package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

var characterDirectory string

type Character struct {
	Name   string `json:"name"`
	Id     string `json:"id"`
	Health int    `json:"health"`
	Str    int    `json:"str"`
	Dex    int    `json:"dex"`
	Con    int    `json:"con"`
	Int    int    `json:"int"`
	Wis    int    `json:"wis"`
	Cha    int    `json:"cha"`
	Stats
}

type Stats struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
}

func main() {
	characterDirectory = "./assets/characters/"
	for {
		cliOptions()
	}
}

// func calculateStatBonus(stat int) int {
// 	statBonus := (stat - 10) / 2
// 	return statBonus
// }

func cliOptions() {
	blue := "\033[34m"
	reset := "\033[0m"
	var command string

	fmt.Println(string(blue), "What would you like to do?", string(reset))
	fmt.Println("c to create a character, l for list all current characters, or e to exit")
	fmt.Scanln(&command)

	if strings.ToLower(command) == "c" {
		c := gatherCLICharacter()
		c.writeCharacter(characterDirectory)
		displayCharacter(c)
	} else if strings.ToLower(command) == "l" {
		allCharacters := allCharacters()
		for _, c := range allCharacters {
			fmt.Println(c.Name)
		}
	} else if strings.ToLower(command) == "e" {
		os.Exit(1)
	}
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

	var command string
	fmt.Println("Would you like to fill out stats? Y or n?")
	fmt.Scan(&command)
	var s Stats
	if strings.ToLower(command) == "y" {
		fmt.Print("Strength:")
		fmt.Scan(&s.Strength)

		fmt.Print("Dexterity:")
		fmt.Scan(&s.Dexterity)

		fmt.Print("Constitution:")
		fmt.Scan(&s.Constitution)

		fmt.Print("Intelligence:")
		fmt.Scan(&s.Intelligence)

		fmt.Print("Wisdom:")
		fmt.Scan(&s.Wisdom)

		fmt.Print("Charisma:")
		fmt.Scan(&s.Charisma)
	}

	return c
}

func (c *Character) writeCharacter(cDirectory string) error {
	jsonData, err := json.Marshal(c)
	if err != nil {
		return err
	}
	os.WriteFile(fmt.Sprintf("%s%s.json", cDirectory, c.Id), jsonData, 0764)
	return nil
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
