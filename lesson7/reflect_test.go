package main

import (
	"log"
	"testing"
)

func TestProcessAnimalCat(t *testing.T) {
	valuesCat := map[string]interface{}{
		"Name":   "Cat",
		"Legs":   4,
		"Weight": 3.5,
		"Voice":  "meow",
		"Roar":   false,
	}

	var cat Animal

	err := processAnimal(cat, valuesCat)
	if err != nil {
		log.Println(err)
	}
}

func TestProcessAnimalDog(t *testing.T) {
	valuesDog := map[string]interface{}{
		"Name":   "Dog",
		"Legs":   4,
		"Weight": 6.5,
		"Voice":  "bark",
		"Roar":   true,
	}

	var dog Animal

	err := processAnimal(dog, valuesDog)
	if err != nil {
		log.Println(err)
	}
}

func TestProcessAnimalBird(t *testing.T) {
	valuesBird := map[string]interface{}{
		"Name":   "Bird",
		"Legs":   2,
		"Weight": 0.5,
		"Voice":  "squack",
		"Roar":   false,
	}

	var bird Animal

	err := processAnimal(bird, valuesBird)
	if err != nil {
		log.Println(err)
	}
}
