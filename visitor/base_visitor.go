package visitor

import (
	"fmt"
	"go/ast"
	"go/token"
)

type BaseVisitor struct {
	FSet         *token.FileSet
	RFilePath    string
	FileBytes    []byte
	Package      string
	PkgVarMap    map[string]*PkgVar
	PkgStructMap map[string]*PkgStruct
}

type BaseInfo struct {
	Name      string
	Content   string
	StartLine int
	EndLine   int
}

type PkgVar struct {
	BaseInfo
	Type string
}

type PkgStruct struct {
	BaseInfo
	Type       string
	DepStructs []string
}

func (v *BaseVisitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}
	// 检查节点是否为函数声明
	if funcDecl, ok := node.(*ast.FuncDecl); ok {
		fmt.Printf("发现函数: %s\n", funcDecl.Name.Name)
	}
	return v
}
