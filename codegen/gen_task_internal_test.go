package codegen

import (
	"github.com/golang/mock/gomock"
	"github.com/lonegunmanb/syrinx/ast"
	"github.com/stretchr/testify/assert"
	"go-funk"
	"testing"
)

func TestGetPkgNameFromPkgPath(t *testing.T) {
	cases := []*funk.Tuple{
		{"testing", "testing"},
		{"go/ast", "ast"},
		{"github.com/lonegunman/syrinx/codegen", "codegen"},
	}
	for _, tuple := range cases {
		assert.Equal(t, tuple.Element2.(string), getPkgNameFromPkgPath(tuple.Element1.(string)))
	}
}

func TestGetDepPkgPathsWithPkgNameDuplicate(t *testing.T) {
	paths := []string{
		"a",
		"b/b",
		"c/b",
	}
	typeInfos := []ast.TypeInfo{}
	ctrl := gomock.NewController(t)
	for _, path := range paths {
		mockTypeInfo := NewMockTypeInfo(ctrl)
		mockTypeInfo.EXPECT().GetDepPkgPaths().Times(1).Return([]string{path})
		typeInfos = append(typeInfos, mockTypeInfo)
	}
	sut := &genTask{
		typeInfos: typeInfos,
	}
	pathsReceived := sut.GetDepPkgPaths()
	assert.Equal(t, paths, pathsReceived)
	assert.Equal(t, "a", sut.GetPkgNameFromPkgPath("a"))
	assert.Equal(t, "p0", sut.GetPkgNameFromPkgPath("b/b"))
	assert.Equal(t, "p1", sut.GetPkgNameFromPkgPath("c/b"))
	imports := sut.GenImportDecls()
	expected := []string{
		`"a"`,
		`p0 "b/b"`,
		`p1 "c/b"`,
	}
	assert.Equal(t, expected, imports)
}
