package simple_factory

import (
	"fmt"
)

type animal interface {
	eat(food string) string
}

type dog struct {
	name string
	age  int
}

func (d *dog) eat(food string) string {
	return fmt.Sprintf("dog %s eating %s", d.name, food)
}

type cat struct {
	name string
	age  int
}

func (c *cat) eat(food string) string {
	return fmt.Sprintf("cat %s eating %s", c.name, food)
}

func NewAnimal(a string) animal {
	switch a {
	case "cat":
		return &cat{"miao", 2}
	case "dog":
		return &dog{"wang", 1}
	default:
		return &cat{"m", 0}
	}
}

func main() {
	dog1 := NewAnimal("dog")
	cat1 := NewAnimal("cat")

	fmt.Printf("%s", dog1.eat("b"))
	fmt.Printf("%s", cat1.eat("c"))
}
