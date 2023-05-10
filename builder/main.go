package main

import (
	"fmt"
	"strings"
)

type Car struct {
	model  string
	engine string
	seats  int
}

func (c Car) String() string {
	sb := strings.Builder{}
	sb.WriteString("The car has ")
	if c.model != "" {
		sb.WriteString(fmt.Sprintf("%s model, ", c.model))
	}
	if c.engine != "" {
		sb.WriteString(fmt.Sprintf("%s engine, ", c.engine))
	} else {
		sb.WriteString(fmt.Sprint("no engine, "))
	}
	if c.seats != 0 {
		sb.WriteString(fmt.Sprintf("%d seats ", c.seats))
	} else {
		sb.WriteString(fmt.Sprint("no seats "))
	}

	return sb.String()
}

type CarBuilder struct {
	car Car
}

func (c *CarBuilder) SetModel(model string) *CarBuilder {
	c.car.model = model
	return c
}

func (c *CarBuilder) SetEngine(engine string) *CarBuilder {
	c.car.engine = engine
	return c
}

func (c *CarBuilder) SetSeats(seats int) *CarBuilder {
	c.car.seats = seats
	return c
}

func (c *CarBuilder) GetCar() *Car {
	return &c.car
}

func main() {
	car1 := Car{"sport", "large", 2}
	fmt.Println(car1)

	cb2 := &CarBuilder{Car{}}
	car2 := cb2.SetModel("sport").SetEngine("large").SetSeats(2).GetCar()
	fmt.Println(car2)

	cb3 := &CarBuilder{Car{}}
	car3 := cb3.SetEngine("large").SetSeats(4).GetCar()
	fmt.Println(car3)

}
