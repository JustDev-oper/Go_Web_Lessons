package main

// skip lesson

import "fmt"

// Определение интерфейса
type Animal interface {
	Speak() string
}

// Структура Dog
type Dog struct {
	Name string
}

// Метод для Dog
func (d Dog) Speak() string {
	return "Woof! My name is " + d.Name
}

// Структура Cat
type Cat struct {
	Name string
}

// Метод для Cat
func (c Cat) Speak() string {
	return "Meow! I'm " + c.Name
}

// Функция, использующая интерфейс Animal
func AnimalSound(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	dog := Dog{Name: "Rex"}
	cat := Cat{Name: "Whiskers"}

	AnimalSound(dog)
	AnimalSound(cat)

	// Использование композиции
	type PetShop struct {
		Animals []Animal
	}

	shop := PetShop{
		Animals: []Animal{dog, cat},
	}

	for _, animal := range shop.Animals {
		AnimalSound(animal)
	}
}
