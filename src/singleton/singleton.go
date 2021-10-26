package singleton

import (
	"fmt"
	"sync"
)

type Car struct {
	name string
}

var car *Car
var once sync.Once

func Getcar() *Car {
	once.Do(func() {
		car = &Car{"suv"}
	})
	return car
}

func main() {
	car := Getcar()
	fmt.Printf("car's name is %s\n", car.name)
	newCar := Getcar()
	fmt.Printf("car == newCar = %t\n", car == newCar)
}
