package visitor

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"testing"
)

func TestBaseVisitor(t *testing.T) {
	filePath := "/Users/silhouette/codeworks/callgraph/service/simple_call.go"
	fset := token.NewFileSet()
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("ReadFile err %v", err)
		return
	}
	v := &BaseVisitor{
		FSet:      fset,
		RFilePath: filePath,
		FileBytes: fileContent,
	}
	astFile, err := parser.ParseFile(fset, filePath, fileContent, 0)
	if err != nil {
		t.Errorf("ParseFile err %v", err)
		return
	}
	ast.Walk(v, astFile)
}
