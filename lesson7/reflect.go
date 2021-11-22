package main

import (
	"reflect"

	"github.com/davecgh/go-spew/spew"
)

type Animal struct {
	Name   string
	Legs   int
	Weight float64
	Voice  string
	Roar   bool
}

func processAnimal(animal Animal, values map[string]interface{}) error {
	valuesOfAnimal := reflect.ValueOf(&animal)
	elemOfAnimal := valuesOfAnimal.Elem()

	for i := 0; i < elemOfAnimal.Type().NumField(); i++ {
		vName := elemOfAnimal.Type().Field(i).Name
		vPar := elemOfAnimal.FieldByName(vName)

		for k, v := range values {
			if k == vName && vPar.Type().AssignableTo(reflect.TypeOf(v)) {
				vPar.Set(reflect.ValueOf(v))
			}
		}
	}
	spew.Dump(animal)
	return nil
}
