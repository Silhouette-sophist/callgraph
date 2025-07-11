package main

import (
	"fmt"

	"golang.org/x/tools/go/packages"
)

func main() {
	fmt.Println("callgraph")
	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedTypesInfo,
		Dir:  ".",
	}
	pkgs, err := packages.Load(cfg, "github.com/Silhouette-sophist/callgraph/service/...")
	if err != nil {
		fmt.Printf("packages.Load err %v\n", err)
		return
	}
	fmt.Printf("packages.Load len %d\n", len(pkgs))
	for _, pkg := range pkgs {
		fmt.Printf("load pkg info %v\n", pkg)
	}
}
