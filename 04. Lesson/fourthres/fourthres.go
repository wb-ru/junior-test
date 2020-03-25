package fourthres

import "fmt"

type zoo interface {
	setInfo(name string, age int)
	GetInfo() (string, int)
}

type animal struct {
	name string
	age  int
}

func (a animal) setInfo(name string, age int) {
	a.name = name
	a.age = age
}

func (a animal) GetInfo() (string, int) {
	return a.name, a.age
}

type hippopotamus struct {
	animal
}

type zebra struct {
	animal
}

func newHippopotamus(name string, age int) zoo {
	return hippopotamus{
		animal: animal{
			name: name,
			age:  age,
		},
	}
}

func newZebra(name string, age int) zoo {
	return zebra{
		animal: animal{
			name: name,
			age:  age,
		},
	}
}

func adoptPet(kind string, name string, age int) (zoo, error) {
	switch kind {
	case "hippopotamus":
		return newHippopotamus(name, age), nil
	case "zebra":
		return newZebra(name, age), nil
	default:
		return nil, fmt.Errorf("We can`t adopt this kind yet")
	}
}
