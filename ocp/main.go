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

type Specification interface {
	IsSpecified(p Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSpecified(p Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSpecified(p Product) bool {
	return p.size == s.size
}

type AndSpecification struct {
	first  Specification
	second Specification
}

func (a AndSpecification) IsSpecified(p Product) bool {
	return a.first.IsSpecified(p) && a.second.IsSpecified(p)
}

func FilterBySpecification(products []Product, spec Specification) []Product {
	result := make([]Product, 0)
	for _, v := range products {
		if spec.IsSpecified(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, medium}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	greenSpecification := ColorSpecification{green}
	largeSpecification := SizeSpecification{large}
	mediumSpecification := SizeSpecification{medium}
	greenMediumSpecification := AndSpecification{greenSpecification, mediumSpecification}

	for _, v := range FilterBySpecification(products, greenSpecification) {
		fmt.Printf("Product %s is green\n", v.name)
	}

	for _, v := range FilterBySpecification(products, largeSpecification) {
		fmt.Printf("Product %s is large\n", v.name)
	}

	for _, v := range FilterBySpecification(products, greenMediumSpecification) {
		fmt.Printf("Product %s is green and medium\n", v.name)
	}

}
