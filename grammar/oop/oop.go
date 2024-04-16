package oop

import "fmt"

// 封装：首字符大小写控制访问权限
//type Animal struct {
//	name string
//}
//
//func NewAnimal() *Animal {
//	return &Animal{}
//}
//
//func (p *Animal) SetName(name string) {
//	p.name = name
//}
//
//func (p *Animal) GetName() string {
//	return p.name
//}

// 继承：go中的继承是通过组合的方式，没有层次结构，相对更加简单。
//type Animal struct {
//	Name string
//}
//
//type Cat struct {
//	Animal
//	FeatureA string
//}
//
//type Dog struct {
//	Animal
//	FeatureB string
//}
//
//func NewAnimal() *Animal {
//	return &Animal{}
//}
//
//func (p *Animal) SetName(name string) {
//	p.Name = name
//}
//
//func (dog *Dog) HelloWorld() {
//	fmt.Println("脑子进煎鱼了")
//}
//
//func (cat *Cat) HelloWorld() {
//	fmt.Println("煎鱼进脑子了")
//}
//
//func main() {
//	p := NewAnimal()
//	p.SetName("我是搬运工，去给煎鱼点赞~")
//
//	dog := Dog{Animal: *p}
//	fmt.Println(dog.Name)
//}

// 多态：go的多态是通过实现接口中的方法来实现的。
type AnimalSounder interface {
	MakeDNA()
}

func MakeSomeDNA(animalSounder AnimalSounder) {
	animalSounder.MakeDNA()
}

type Cat struct{}
type Dog struct{}

func (c *Cat) MakeDNA() {
	fmt.Println("煎鱼是煎鱼")
}

func (c *Dog) MakeDNA() {
	fmt.Println("煎鱼其实不是煎鱼")
}

func main() {
	MakeSomeDNA(&Cat{})
	MakeSomeDNA(&Dog{})
}
