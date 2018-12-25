package codegen

//go:generate mockgen -package=codegen -destination=./mock_type_info.go github.com/lonegunmanb/syrinx/ast TypeInfo
import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenPackageDecl(t *testing.T) {
	testGen(t, func(typeInfo *MockTypeInfo) {
		typeInfo.EXPECT().GetPkgName().Times(1).Return("ast")
	}, func(gen *codegen) error {
		return gen.genPkgDecl()
	}, "package ast")
}

func TestGenImportsDecl(t *testing.T) {
	testGen(t, func(typeInfo *MockTypeInfo) {
		depImports := []string{
			"go/ast",
			"go/token",
			"go/types",
		}
		setupMockToGenImports(typeInfo, depImports)
	}, func(gen *codegen) error {
		return gen.genImportDecls()
	}, `
import (
"github.com/lonegunmanb/syrinx/ioc"
"go/ast"
"go/token"
"go/types"
)`)
}

func TestShouldNotGenExtraImportsIfDepPathsEmpty(t *testing.T) {
	testGen(t, func(typeInfo *MockTypeInfo) {
		var depImports []string
		setupMockToGenImports(typeInfo, depImports)
	}, func(gen *codegen) error {
		return gen.genImportDecls()
	}, `
import (
"github.com/lonegunmanb/syrinx/ioc"
)`)
}

func testGen(t *testing.T, setupMockFunc func(info *MockTypeInfo), testMethod func(gen *codegen) error, expected string) {
	writer := &bytes.Buffer{}
	ctrl, typeInfo := prepareMock(t)
	defer ctrl.Finish()
	//
	setupMockFunc(typeInfo)
	codegen := newCodegen(typeInfo, writer)
	err := testMethod(codegen)
	assert.Nil(t, err)
	code := writer.String()
	assert.Equal(t, expected, code)
}

func setupMockToGenImports(typeInfo *MockTypeInfo, depImports []string) {
	typeInfo.EXPECT().GetDepPkgPaths().Times(1).Return(depImports)
}

func prepareMock(t *testing.T) (*gomock.Controller, *MockTypeInfo) {
	ctrl := gomock.NewController(t)
	mock := NewMockTypeInfo(ctrl)
	return ctrl, mock
}
