package main

import (
	"fmt"

	"github.com/janhaans/godesignpatterns/abstractfactory"
)
func createFurniture(furnitureFactory abstractfactory.FurnitureAbstractFactory) {
	chair := furnitureFactory.CreateChair()
	table := furnitureFactory.CreateTable()
	sofa := furnitureFactory.CreateSofa()
	fmt.Println(chair)
	fmt.Println(table)
	fmt.Println(sofa)
}

func main() {
	createFurniture(abstractfactory.ArtDecoFurnitureFactory{})
	createFurniture(abstractfactory.VictorianFurnitureFactory{})
	createFurniture(abstractfactory.ModernFurnitureFactory{})
}