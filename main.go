package main

import (
	"fmt"

	"io/ioutil"
	"os"

	"interface/lib"
)

func main() {
	f, err := os.Open("lib/test.go")

	if err != nil {
		fmt.Println(err)
		return
	}

	src, _ := ioutil.ReadAll(f)

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
			funcation := ScanFunc(tok, lit, &s)
			fmt.Println(funcation.String())
		}
	}
}
