package main

import (
	"fmt"
	"math/rand"
)

type Item interface {
	Name() string
	Info() string
}
type Fish struct {
	name string
	rarity int
	price int
}

type LegendaryFish struct {
	name string
	rarity string
	price int 
}

func (f Fish) Name() string { return f.name }
func (f Fish) Info() string { return fmt.Sprintf("Редкость: %d, Цена: %d золота", f.rarity, f.price)}
func (lf LegendaryFish) Name() string { return lf.name }
func (lf LegendaryFish) Info() string { return fmt.Sprintf("Редкость: %s, Цена: %d золота", lf.rarity, lf.price)}
type Bait struct {
	name string
	power int
}
type Player struct {
	fishPower int
}
var (
	commonFish = []Fish{
		{name: "окунь", rarity: 10, price: 1},
		{name: "рыба ласточка", rarity: 15, price: 1},
	}
	rareFish = []LegendaryFish{
		{name:"люфтяга", rarity: "бесценно", price: 999},
	}
)

func (p Player) Cast(b Bait) Item{
	totalPower := p.fishPower + b.power
	fmt.Printf("Игрок закидывает: %s, Мастерство: %d\n", b.name, totalPower)
	roll := rand.Intn(100) + totalPower

	if roll > 120 {
		return rareFish[0]
	} else {
		return commonFish[rand.Intn(len(commonFish))]
	}
}

func main() {
	p := Player{fishPower: 50}
	b := Bait{name: "Нубская наживка", power: 40}

	fmt.Println("Игрок начинает рыбалку")

	for i := 1; i <= 5; i++{
		catch := p.Cast(b)
		fmt.Printf("Улов %d: %s (%s)\n\n", i, catch.Name(), catch.Info())
	}
}