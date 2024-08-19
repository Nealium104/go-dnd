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
