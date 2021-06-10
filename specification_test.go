package godesignpatterns

import (
	"testing"
)

var data = []Car{
	{"12-shd-6", Mercedes, Green, Large},
	{"63-jhd-9", BMW, Red, Medium},
	{"72-skn-8", Volvo, Yellow, Large},
	{"80-owp-2", Volkswagen, Blue, Small},
	{"99-kfx-5", Audi, Green, Medium},
	{LicensePlate: "76-wkl-6", CarBrand: BMW, Size: Large},
}

func TestFilterByGreen(t *testing.T) {
	testData := []struct {
		cars []Car
		ColorSpecification
		want int
	}{
		{data, ColorSpecification{Green}, 2},
		{data, ColorSpecification{Red}, 1},
		{data, ColorSpecification{Yellow}, 1},
		{data, ColorSpecification{Blue}, 1},
	}

	for _, test := range testData {
		filteredCars := FilterBySpecification(test.cars, test.ColorSpecification)
		if len(filteredCars) != test.want {
			t.Errorf("FilterBySpecification got %d %v cars but expected %d %[2]v cars\n ", len(filteredCars), test.ColorSpecification.Color, test.want)
			continue
		}
	}

}
