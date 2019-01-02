package test_code_test

import (
	"github.com/lonegunmanb/syrinx/ioc"
	"github.com/lonegunmanb/syrinx/test_code/engine"
	"github.com/lonegunmanb/syrinx/test_code/fly_car"
	"github.com/lonegunmanb/syrinx/test_code/flyer"
	"github.com/lonegunmanb/syrinx/test_code/ioc_gen"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateFlyCarCarByManualSimulate(t *testing.T) {
	container := ioc_gen.CreateIoc()
	e := engine.NewEngine(100)
	container.RegisterFactory((*engine.Engine)(nil), func(ioc ioc.Container) interface{} {
		return e
	})
	decoration := &fly_car.FancyDecoration{}
	container.RegisterFactory((*fly_car.Decoration)(nil), func(ioc ioc.Container) interface{} {
		return decoration
	})
	wing := &flyer.AluminumWing{}
	container.RegisterFactory((*flyer.Wing)(nil), func(ioc ioc.Container) interface{} {
		return wing
	})
	c := fly_car.Create_FlyCar(container)
	assert.Equal(t, e, c.Engine)
	assert.Equal(t, wing, c.Wing)
	assert.Equal(t, decoration, c.Decoration)
	assert.NotNil(t, c.R1)
	assert.NotNil(t, c.R2)
}
