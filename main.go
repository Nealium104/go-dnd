package main

import (
	"fmt"
	"math/rand/v2"
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
	Name           string
	BaseHealth     int
	InitativeBonus int
	Initative      int
	ItemList       []Item
}

type Item struct {
	Name  string
	Bonus int
}

func main() {
	characterSlice := []Character{}
	var amount int
	fmt.Println("Hello, go-dnd! How many monsters and characters will you be using?")
	fmt.Scanln(&amount)
	for i := 0; i < amount; i++ {
		characterSlice = append(characterSlice, createCharacter())
	}
	fmt.Println("Rolling initative!")
	for i := 0; i < len(characterSlice); i++ {
		currentCharacter := characterSlice[i]
		rollInitative(currentCharacter)
	}
	sorted := sortInitativeDescending(characterSlice)
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

func roll() int {
	return rand.IntN(20)
}

func rollInitative(character Character) {
	roll := roll()
	character.Initative = roll + character.InitativeBonus
	fmt.Printf("%v rolled a %v. Their bonus is %v, so their initative is %v\n", character.Name, roll, character.InitativeBonus, roll+character.Initative)
}

func sortInitativeDescending(characterSlice []Character) []Character {
	if len(characterSlice) < 2 {
		return characterSlice
	} else {
		pivot := characterSlice[0]
		less := []Character{}
		greater := []Character{}

		for _, i := range characterSlice[1:] {
			if i.Initative >= pivot.Initative {
				less = append(less, i)
			} else {
				greater = append(greater, i)
			}
		}

		sortedLess := sortInitativeDescending(less)
		sortedGreater := sortInitativeDescending(greater)

		return append(append(sortedLess, pivot), sortedGreater...)
	}

}

func createCharacter() Character {
	var name string
	var health int
	var initativeBonus int
	fmt.Println("What is the name of this character?")
	fmt.Scanln(&name)
	fmt.Printf("What is the health of %v?\n", name)
	fmt.Scanln(&health)
	fmt.Printf("What is the initative bonus of %v?\n", name)
	fmt.Scanln(&initativeBonus)
	return Character{Name: name, BaseHealth: health, InitativeBonus: initativeBonus}
}
