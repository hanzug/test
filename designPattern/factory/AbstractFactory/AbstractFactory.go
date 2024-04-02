package main

import (
	"fmt"
)

// 抽象工厂模式下，每种水果定义为一个interface
type Strawberry interface {
	// 草莓的甜蜜攻击
	SweetAttack()
}

type Lemon interface {
	// 柠檬的酸劲攻击
	AcidAttack()
}

// FruitFactory 抽象工厂类
type FruitFactory interface {
	CreateStrawberry() Strawberry
	CreateLemon() Lemon
}

type GoodfarmerStrawberry struct {
	brand string
	Strawberry
}

func (g *GoodfarmerStrawberry) SweetAttack() {
	fmt.Printf("sweet attack from %s, ", g.brand)
}

type GoodfarmerLemon struct {
	brand string
	Lemon
}

func (g *GoodfarmerLemon) AcidAttack() {
	fmt.Printf("acid attack from %s, ", g.brand)
}

type DoleStrawberry struct {
	brand string
	Strawberry
}

func (d *DoleStrawberry) SweetAttack() {
	fmt.Printf("sweet attack from %s, ", d.brand)
}

type DoleLemon struct {
	brand string
	Lemon
}

func (d *DoleLemon) AcidAttack() {
	fmt.Printf("acid attack from %s,", d.brand)
}

type GoodfarmerFactory struct{}

func (g *GoodfarmerFactory) myAspect() {
	fmt.Println("good farmer aspect...")
}

func (g *GoodfarmerFactory) CreateStrawberry() Strawberry {
	// 同一个产品族可以插入一个切面
	g.myAspect()
	defer g.myAspect()
	return &GoodfarmerStrawberry{
		brand: "goodfarmer",
	}
}
func (g *GoodfarmerFactory) CreateLemon() Lemon {
	// 同一个产品族可以插入一个切面
	g.myAspect()
	defer g.myAspect()
	return &GoodfarmerLemon{
		brand: "goodfarmer",
	}
}

type DoleFactory struct{}

func (d *DoleFactory) myAspect() {
	fmt.Println("dole aspect...")
}

func (d *DoleFactory) CreateStrawberry() Strawberry {
	// 同一个产品族可以插入一个切面
	d.myAspect()
	//defer d.Myspect()
	return &DoleStrawberry{
		brand: "dole",
	}
}
func (d *DoleFactory) CreateLemon() Lemon {
	// 同一个产品族可以插入一个切面
	d.myAspect()
	//defer d.Myspect()
	return &DoleLemon{
		brand: "dole",
	}
}

//倘若我们需要额外扩展一个新的水果品牌，比如佳沛 Zespri，此时需要额外新增如下代码：

type ZespriStrawberry struct {
	brand string
	Strawberry
}

func (z *ZespriStrawberry) SweetAttack() {
	fmt.Printf("sweet attack from %s, ", z.brand)
}

type ZespriLemon struct {
	brand string
	Lemon
}

func (z *ZespriLemon) AcidAttack() {
	fmt.Printf("acid attack from %s, ", z.brand)
}

type ZespriFactory struct{}

func (z *ZespriFactory) myAspect() {
	fmt.Println("dole aspect...")
}

func (z *ZespriFactory) CreateStrawberry() Strawberry {
	// 同一个产品族可以插入一个切面
	z.myAspect()
	//defer z.Myspect()
	return &ZespriStrawberry{
		brand: "zespri",
	}
}
func (z *ZespriFactory) CreateLemon() Lemon {
	// 同一个产品族可以插入一个切面
	z.myAspect()
	//defer z.Myspect()
	return &ZespriLemon{
		brand: "zespri",
	}
}

func main() {
	// 尝尝佳农品牌的水果
	goodfarmerFactory := GoodfarmerFactory{}
	goodfarmerStrawberry := goodfarmerFactory.CreateStrawberry()
	goodfarmerStrawberry.SweetAttack()

	goodfarmerLemon := goodfarmerFactory.CreateLemon()
	goodfarmerLemon.AcidAttack()

	// 尝尝都乐品牌的水果
	doleFactory := DoleFactory{}
	doleStrawberry := doleFactory.CreateStrawberry()
	doleStrawberry.SweetAttack()

	doleLemon := doleFactory.CreateLemon()
	doleLemon.AcidAttack()
}
