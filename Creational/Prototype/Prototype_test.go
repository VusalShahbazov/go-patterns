package Prototype

import (
	"testing"
)

func TestCharacter_Clone(t *testing.T) {
	character1 := Character{
		Name:  "Enemy",
		Color: "Black",
	}

	character2 := character1.Clone()

	if !isEqual(character2 , character1){
		t.Error("Does not equal")
	}
}

func isEqual(a, b Character) bool {
	if a.Name != b.Name {
		return false
	}
	if a.Color != b.Color {
		return false
	}
	return true
}