package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	car := Getcar()
	newCar := Getcar()
	assert.Equal(t, car, newCar)

}
