package main

import (
	"design_mode/house_builder"
	"fmt"
)

// 生成器设计模式
func generate_test() {
	normalBuilder := house_builder.GetBuilder("normal")
	iglooBuilder := house_builder.GetBuilder("igloo")

	// 客户端与主管交互生成产品, 创建步骤由主管负责执行
	director := house_builder.NewDirector(normalBuilder)
	normalHouse2 := director.BuildHouse()
	printHouseInfo(normalHouse2)

	director.SetBuilder(iglooBuilder)
	iglooHouse2 := director.BuildHouse()
	printHouseInfo(iglooHouse2)
}

func printHouseInfo(house house_builder.House) {
	fmt.Printf("房子的门是: %s, 窗户是 %s, 楼层是: %d\n", house.DoorType, house.WindowType, house.Floor)
}

// output:
// 房子的门是: Wooden Door, 窗户是 Wooden Window, 楼层是: 2
// 房子的门是: Snow Door, 窗户是 Snow Window, 楼层是: 1
