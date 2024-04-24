package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	Test int
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.Name, p.Age)
}

func inspectStruct(s interface{}) {
	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	fmt.Println("Inspecting:", typ.Name())

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fmt.Printf("%d. Field: %s, Type: %s, Value: %v\n", i+1, typ.Field(i).Name, field.Type(), field.Interface())
	}

	if method := val.MethodByName("Greet"); method.IsValid() {
		result := method.Call(nil)
		if len(result) > 0 {
			fmt.Println("Greet Method Call Result:", result[0].Interface())
		}
	}
}

func main() {
	p := Person{Name: "John", Age: 30}
	inspectStruct(p)
}
