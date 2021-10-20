package simple_factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAnimalEat(t *testing.T) {
	dog1 := NewAnimal("dog")
	cat1 := NewAnimal("cat")

	dog1.eat("b")
	cat1.eat("c")

	assert.Equal(t, dog1.eat("beef"), "dog wang eating beef")
	assert.Equal(t, cat1.eat("beef"), "cat miao eating beef")

}
