package main

import (
	"fmt"
	"github.com/godesignpatterns/abstractfactory"
)
func createFurniture(furnitureFactory FurnitureAbstractFactory) {
	chair := furnitureFactory.CreateChair()
	table := furnitureFactory.CreateTable()
	sofa := furnitureFactory.CreateSofa()
	fmt.Println(chair)
	fmt.Println(table)
	fmt.Println(sofa)
}

func main() {
	createFurniture(ArtDecoFurnitureFactory{})
	createFurniture(VictorianFurnitureFactory{})
	createFurniture(ModernFurnitureFactory{})
}