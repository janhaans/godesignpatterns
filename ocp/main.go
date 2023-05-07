package main

import "fmt"

type Color int
type Size int

const (
	green Color = iota
	red
	blue
)

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct{}

func (f *Filter) FilterByColor(products []Product, color Color) []Product {
	result := make([]Product, 0)
	for _, v := range products {
		if v.color == color {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", red, medium}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	f := Filter{}

	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf("Product %s is green", v.name)
	}

}
