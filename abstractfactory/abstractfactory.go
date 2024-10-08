package abstractfactory

// FurnitureAbstractFactory is an interface for creating furniture
type FurnitureAbstractFactory interface {
	CreateChair() Chair
	CreateTable() Table
	CreateSofa() Sofa
}

// Chair is a struct for a chair
type Chair struct {
	name string
	style string
}

// Table is a struct for a table
type Table struct {
	name string
	style string
}

// Sofa is a struct for a sofa
type Sofa struct {
	name string
	style string
}

// ArtDecoFurnitureFactory is a struct for creating Art Deco furniture
type ArtDecoFurnitureFactory struct {}

// CreateChair creates a chair
func (a ArtDecoFurnitureFactory) CreateChair() Chair {
	return Chair{"Art Deco Chair", "Art Deco"}
}

// CreateTable creates a table
func (a ArtDecoFurnitureFactory) CreateTable() Table {
	return Table{"Art Deco Table", "Art Deco"}
}

// CreateSofa creates a sofa
func (a ArtDecoFurnitureFactory) CreateSofa() Sofa {
	return Sofa{"Art Deco Sofa", "Art Deco"}
}
// VictorianFurnitureFactory is a struct for creating Victorian furniture
type VictorianFurnitureFactory struct {}

// CreateChair creates a Victorian chair
func (v VictorianFurnitureFactory) CreateChair() Chair {
	return Chair{"Victorian Chair", "Victorian"}
}

// CreateTable creates a Victorian table
func (v VictorianFurnitureFactory) CreateTable() Table {
	return Table{"Victorian Table", "Victorian"}
}

// CreateSofa creates a Victorian sofa
func (v VictorianFurnitureFactory) CreateSofa() Sofa {
	return Sofa{"Victorian Sofa", "Victorian"}
}

// ModernFurnitureFactory is a struct for creating Modern furniture
type ModernFurnitureFactory struct {}

// CreateChair creates a Modern chair
func (m ModernFurnitureFactory) CreateChair() Chair {
	return Chair{"Modern Chair", "Modern"}
}

// CreateTable creates a Modern table
func (m ModernFurnitureFactory) CreateTable() Table {
	return Table{"Modern Table", "Modern"}
}

// CreateSofa creates a Modern sofa
func (m ModernFurnitureFactory) CreateSofa() Sofa {
	return Sofa{"Modern Sofa", "Modern"}
}

