package lib

import (
	"fmt"
	"go/scanner"
	"go/token"
)

type Interface struct {
	Name   string
	Method Func
}

type Interfaces []Interface

type Struct struct {
	Name string
}

type Structs []Struct

func (ss *Structs) Scan(tok token.Token, lit string, s *scanner.Scanner) {
	st := Struct{}
	if st.scanStruct(tok, lit, s) {
		ss = append(ss, st)
	}
}

func (st *Struct) scanStruct(tok token.Token, lit string, s *scanner.Scanner) bool {
	if tok != token.TYPE {
		return false
	}

	_, tok, lit = s.Scan()
	if tok != token.IDENT {
		return false
	}
	st.Name = lit

	return true
}
