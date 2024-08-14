package main

import(
	"testing"
)
// type Character struct {
// 	Name           string
// 	Health         int
// 	InitativeBonus int
// 	Initative      int
// 	ItemList       []Item
// }
cSlice := []Character{
	Character{
		Name: "Billy",
		Health: 10,
		InitativeBonus: 3,
	},
	Character{
		Name: "Bob",
		Health: 20,
		InitativeBonus: 2,			
	},
	Character{
		Name: "Jones",
		Health: 30,
		InitativeBonus: 1,			
	},
}
func TestSortOrderByInitative(t *testing.T){
	sortOrderByInitative(cSlice)
}