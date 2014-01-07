package lib

import (
	"fmt"
	"go/scanner"
	"go/token"
)

type Tokens struct {
	Funcations Funcs
}

func Scan(path string) (res Tokens) {

	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, scanner.ScanComments)

	for {
		_, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		if tok == token.FUNC {
			funcation := t.Funcations.Scan(tok, lit, s)
			fmt.Println(funcation.String())
		}
	}
	return
}
