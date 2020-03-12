package Prototype


type IClone interface {
	Clone () interface{}
}
type Character struct {
	Name string
	Color string
}

func (c *Character) Clone() Character {
	newCharacter := Character{}
	newCharacter.Name = c.Name
	newCharacter.Color = c.Color
	return newCharacter
}
