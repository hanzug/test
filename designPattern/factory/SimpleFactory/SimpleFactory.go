package main

import "fmt"

// Shape 是一个接口，定义了一个方法 Draw
type Shape interface {
	Draw()
}

// Circle 实现了 Shape 接口
type Circle struct{}

func (c *Circle) Draw() {
	fmt.Println("Inside Circle::draw() method.")
}

// Rectangle 实现了 Shape 接口
type Rectangle struct{}

func (r *Rectangle) Draw() {
	fmt.Println("Inside Rectangle::draw() method.")
}

// Square 实现了 Shape 接口
type Square struct{}

func (s *Square) Draw() {
	fmt.Println("Inside Square::draw() method.")
}

// ShapeFactory 是一个简单工厂，用来生成基于给定信息的实体类对象
type ShapeFactory struct{}

func (sf *ShapeFactory) GetShape(shapeType string) Shape {
	switch shapeType {
	case "CIRCLE":
		return &Circle{}
	case "RECTANGLE":
		return &Rectangle{}
	case "SQUARE":
		return &Square{}
	default:
		return nil
	}
}

func main() {
	shapeFactory := ShapeFactory{}

	shape1 := shapeFactory.GetShape("CIRCLE")
	if shape1 != nil {
		shape1.Draw()
	}

	shape2 := shapeFactory.GetShape("RECTANGLE")
	if shape2 != nil {
		shape2.Draw()
	}

	shape3 := shapeFactory.GetShape("SQUARE")
	if shape3 != nil {
		shape3.Draw()
	}
}
