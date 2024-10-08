package main

import (
	"fmt"

	"github.com/janhaans/godesignpatterns/furniturefactory"
)
func createFurniture(furnitureFactory furniturefactory.FurnitureAbstractFactory) {
	chair := furnitureFactory.CreateChair()
	table := furnitureFactory.CreateTable()
	sofa := furnitureFactory.CreateSofa()
	fmt.Println(chair)
	fmt.Println(table)
	fmt.Println(sofa)
}

func main() {
	createFurniture(furniturefactory.ArtDecoFurnitureFactory{})
	createFurniture(furniturefactory.VictorianFurnitureFactory{})
	createFurniture(furniturefactory.ModernFurnitureFactory{})
}