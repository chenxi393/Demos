package main

import "fmt"

// 工厂方法（Factory Method）模式
type Product interface {
	Show()
}

type DefaultProduct struct {
}

func (p *DefaultProduct) Show() {
}

type FoodProduct struct {
	Name string
}

func (p *FoodProduct) Show() {
	fmt.Println("FoodProduct产品name:" + p.Name)
}

type WaterProduct struct {
	Name string
}

func (p *WaterProduct) Show() {
	fmt.Println("WaterProduct产品name:" + p.Name)
}

// 简单工厂
func ProductFactor(name string) Product {
	if name == "food" {
		return &FoodProduct{}
	} else if name == "water" {
		return &WaterProduct{}
	} else {
		return &DefaultProduct{}
	}
}

// 方法工厂
var FactOne FactorOne
var FactTwo FactorTwo

type Factor interface {
	Produce()
}

type FactorOne struct {
}

type FactorTwo struct {
}

func (f *FactorOne) Produce() Product {
	return &FoodProduct{}
}

func (f *FactorTwo) Produce() Product {
	return &WaterProduct{}
}
