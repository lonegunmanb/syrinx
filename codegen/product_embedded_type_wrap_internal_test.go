package codegen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const embeddedAssignTemplate = `
package test
%s
type Struct struct {
%s
}
type Struct2 struct {
}
type Interface interface {
Name() string
}
`

func TestEmbedded(t *testing.T) {
	testEmbeddedType(t, "", "Struct2", "Struct2", "*", "github.com/lonegunmanb/syrinx/test.Struct2", "*Struct2")
	testEmbeddedType(t, "", "*Struct2", "Struct2", "", "github.com/lonegunmanb/syrinx/test.Struct2", "*Struct2")
	testEmbeddedType(t, "", "Interface", "Interface", "", "github.com/lonegunmanb/syrinx/test.Interface", "Interface")
	importDecl := `import "github.com/lonegunmanb/syrinx/test_code/car"`
	testEmbeddedType(t, importDecl, "car.Car", "Car", "*", "github.com/lonegunmanb/syrinx/test_code/car.Car", "*car.Car")
	testEmbeddedType(t, importDecl, "*car.Car", "Car", "", "github.com/lonegunmanb/syrinx/test_code/car.Car", "*car.Car")
	importDecl = `import "github.com/lonegunmanb/syrinx/test_code/engine"`
	testEmbeddedType(t, importDecl, "engine.Engine", "Engine", "", "github.com/lonegunmanb/syrinx/test_code/engine.Engine", "engine.Engine")
}

type stubTypeCodegen struct{}

func (*stubTypeCodegen) GetName() string {
	panic("implement me")
}

func (*stubTypeCodegen) GetPkgName() string {
	panic("implement me")
}

func (*stubTypeCodegen) GetDepPkgPaths() []string {
	panic("implement me")
}

func (*stubTypeCodegen) GetFieldAssigns() []Assembler {
	panic("implement me")
}

func (*stubTypeCodegen) GetEmbeddedTypeAssigns() []Assembler {
	panic("implement me")
}

func (*stubTypeCodegen) GetPkgNameFromPkgPath(pkgPath string) string {
	return getPkgNameFromPkgPath(pkgPath)
}

func testEmbeddedType(t *testing.T, importString string, typeDecl string, assignedField string, star string, key string, convertType string) {
	code := fmt.Sprintf(embeddedAssignTemplate, importString, typeDecl)
	walker := parseCode(t, code)
	embeddedType := &productEmbeddedTypeWrap{walker.GetTypes()[0].GetEmbeddedTypes()[0], &stubTypeCodegen{}}
	expected := fmt.Sprintf(`product.%s = %scontainer.Resolve("%s").(%s)`, assignedField, star, key, convertType)
	assert.Equal(t, expected, embeddedType.AssembleCode())
}
