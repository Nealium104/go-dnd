package main

import (
	"testing"
)

// type Character struct {
// 	Name           string
// 	Health         int
// 	InitativeBonus int
// 	Initative      int
// 	ItemList       []Item
// }

func TestSortOrderByInitative(t *testing.T) {
	cSlice := []Character{}
	result := sortOrderByInitative(cSlice)
	if result != []Character{
		{
			Name:           "Billy",
			Health:         10,
			InitativeBonus: 3,
		},
		{
			Name:           "Bob",
			Health:         20,
			InitativeBonus: 2,
		},
		{
			Name:           "Jones",
			Health:         30,
			InitativeBonus: 1,
		},
	} {
		t.Errorf("Result was incorrect %v", 8)
	}
}
