package factory_method

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactoryMethod(t *testing.T) {
	af := NewAcpuFactory()
	a1 := af.create()
	a2 := af.create()

	bf := NewBcpuFactory()
	b1 := bf.create()
	b2 := bf.create()

	assert.Equal(t, a1.calculate(), "Acpu has 2 core calculate result")
	assert.Equal(t, a1.calculate(), a2.calculate())
	assert.Equal(t, b1.calculate(), "Bcpu has 4 core calculate result")
	assert.Equal(t, b1.calculate(), b2.calculate())

}
