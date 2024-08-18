package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"sort"
)

type StatsList struct {
	Strength     int
	Str          int
	Dexterity    int
	Dex          int
	Constitution int
	Con          int
	Wisdom       int
	Wis          int
	Intelligence int
	Int          int
	Charisma     int
	Cha          int
}

type Character struct {
	StatsList
	Name           string `json:"name"`
	BaseHealth     int    `json:"BaseHealth"`
	InitativeBonus int    `json:"InitativeBonus"`
	Initative      int    `json:"Initative"`
	ItemList       []Item `json:"ItemList"`
}

type Item struct {
	Name  string
	Bonus int
}

type CombatList struct {
	CharacterList []Character
}

func roll() int {
	return rand.IntN(20)
}

func rollInitative(character Character) {
	roll := roll()
	character.Initative = roll + character.InitativeBonus
	fmt.Printf("%v rolled a %v. Their bonus is %v, so their initative is %v\n", character.Name, roll, character.InitativeBonus, character.Initative)
}

func sortOrderByInitative(characterSlice []Character) []Character {
	sort.Slice(characterSlice, func(i, j int) bool {
		return characterSlice[i].Initative > characterSlice[j].Initative
	})
	return characterSlice
}

func deriveBonus(stat int) int {
	bonus := (stat - 10) / 2
	return bonus
}

func createCharacter(name string, health int, initativeBonus int) Character {
	c := Character{Name: name, BaseHealth: health, InitativeBonus: initativeBonus}
	saveCharacter(c)
	return c
}

func promptCharacter() Character {
	var name string
	var health int
	var initativeBonus int
	fmt.Println("What is the name of this character?")
	fmt.Scanln(&name)
	fmt.Printf("What is the health of %v?\n", name)
	fmt.Scanln(&health)
	fmt.Printf("What is the initative bonus of %v?\n", name)
	fmt.Scanln(&initativeBonus)
	saveCharacter(createCharacter(name, health, initativeBonus))
	return createCharacter(name, health, initativeBonus)
}

func saveCharacter(c Character) {
	marshaled, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Marshal err:", err)
		return
	}

	file, err := os.Create(fmt.Sprintf("%v.json", c.Name))
	if err != nil {
		log.Fatalf("error with file creation: %v", err)
	}
	defer file.Close()

	_, err = file.Write(marshaled)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
}

func main() {
	characterSlice := []Character{}
	var amount int
	fmt.Println("Hello, go-dnd! How many monsters and characters will you be using?")
	fmt.Scanln(&amount)
	for i := 0; i < amount; i++ {
		characterSlice = append(characterSlice, promptCharacter())
	}
	fmt.Println("Rolling initative!")
	for i := 0; i < len(characterSlice); i++ {
		currentCharacter := characterSlice[i]
		rollInitative(currentCharacter)
	}
	sorted := sortOrderByInitative(characterSlice)
	for i := 0; i < len(sorted); i++ {
		fmt.Printf("Character name is %v, and their health is at %v\n", sorted[i].Name, sorted[i].BaseHealth)
	}
	for i := 0; i < len(sorted); i++ {
		character := sorted[i]
		var res string
		fmt.Printf("It is %v's turn\n", character.Name)
		fmt.Println("Press enter for next character, or type exit to exit")
		fmt.Scanln(&res)
		if res == "exit" {
			break
		}
		if i == len(sorted)-1 {
			i = -1
		}
	}
}
