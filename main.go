package main

import (
	"fmt"
)

type Character struct {
	Name string
}

func main() {
	var characterName string
	fmt.Println("Hello, go-dnd! What's your name?")
	fmt.Scanln(&characterName)
	player := Character{Name: characterName}
	fmt.Printf("Character Name is: %v \n", player.Name)
}