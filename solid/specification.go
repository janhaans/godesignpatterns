package solid

type CarBrand string
type Color string
type Size string

const (
	Mercedes   CarBrand = "mercedes"
	Volkswagen CarBrand = "volkswagen"
	Volvo      CarBrand = "volvo"
	BMW        CarBrand = "bmw"
	Audi       CarBrand = "audi"
)

const (
	Green  Color = "green"
	Red    Color = "red"
	Yellow Color = "yellow"
	Blue   Color = "blue"
)

const (
	Small  Size = "small"
	Medium Size = "medium"
	Large  Size = "large"
)

type Car struct {
	LicensePlate string
	CarBrand
	Color
	Size
}

type Specification interface {
	IsSpecified(c *Car) bool
}

type ColorSpecification struct {
	Color
}

func (c ColorSpecification) IsSpecified(car *Car) bool {
	if car.Color == c.Color {
		return true
	} else {
		return false
	}
}

type SizeSpecification struct {
	Size
}

func (c SizeSpecification) IsSpecified(car *Car) bool {
	if car.Size == c.Size {
		return true
	} else {
		return false
	}
}

type ColorAndSizeSpecification struct {
	Color
	Size
}

func (c ColorAndSizeSpecification) IsSpecified(car *Car) bool {
	if c.Color == car.Color && c.Size == car.Size {
		return true
	} else {
		return false
	}
}

func FilterBySpecification(cars []Car, spec Specification) []*Car {
	result := make([]*Car, 0)
	for i := range cars {
		if spec.IsSpecified(&cars[i]) {
			result = append(result, &cars[i])
		}
	}
	return result
}
