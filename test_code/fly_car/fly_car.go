package fly_car

import (
	"github.com/lonegunmanb/syrinx/test_code/car"
	"github.com/lonegunmanb/syrinx/test_code/flyer"
)

type FlyCar struct {
	*car.Car
	flyer.Plane
	Decoration Decoration
}

type Decoration interface {
	LookAndFeel() string
}

type FancyDecoration struct {
}

func (f *FancyDecoration) LookAndFeel() string {
	return "Fancy"
}