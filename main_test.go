package main

import (
	"reflect"
	"testing"
)

func TestSortInitativeDescending(t *testing.T) {
	characters := []Character{
		{
			Name:           "Bob",
			BaseHealth:         20,
			InitativeBonus: 2,
		},
		{
			Name:           "Jones",
			BaseHealth:         30,
			InitativeBonus: 1,
		},		
		{
			Name:           "Billy",
			BaseHealth:         10,
			InitativeBonus: 3,
		},
	}
	expected := []Character{
		{
			Name:           "Billy",
			BaseHealth:         10,
			InitativeBonus: 3,
		},
		{
			Name:           "Bob",
			BaseHealth:         20,
			InitativeBonus: 2,
		},
		{
			Name:           "Jones",
			BaseHealth:         30,
			InitativeBonus: 1,
		},		
	}
	result := sortInitativeDescending(characters)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result was incorrect. Got: %v, Expected: %v", result, expected)
	}
}

// func TestSortOrderByInitative(t *testing.T) {
// 	cSlice := []Character{}
// 	result := sortOrderByInitative(cSlice)
// 	if result != []Character{
// 		{
// 			Name:           "Billy",
// 			BaseHealth:     10,
// 			InitativeBonus: 3,
// 		},
// 		{
// 			Name:           "Bob",
// 			BaseHealth:     20,
// 			InitativeBonus: 2,
// 		},
// 		{
// 			Name:           "Jones",
// 			BaseHealth:     30,
// 			InitativeBonus: 1,
// 		},
// 	} {
// 		t.Errorf("Result was incorrect %v", 8)
// 	}
// }

func TestDeriveBonus(t *testing.T) {
	result := deriveBonus(20)
	expected := 5

	if result != expected {
		t.Errorf("deriveBonus(20) = %d; want %d", result, expected)
	}
}
