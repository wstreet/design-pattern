package simple_factory

import (
	"simple_factory"
	"testing"
)

func TestNewAnimal(t *testing.T) {
	dog1 := NewAnimal("dog")
	cat1 := NewAnimal("cat")

	dog1.eat()
	cat1.eat()

}
