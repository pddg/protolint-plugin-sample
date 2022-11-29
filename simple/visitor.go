package main

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"github.com/yoheimuta/protolint/linter/visitor"
)

type sampleVisitor struct {
	// BaseAddVisitorを埋め込む。BaseAddVisitor自体の初期化は後から行う。
	*visitor.BaseAddVisitor
}

func (s *sampleVisitor) VisitPackage(pkg *parser.Package) bool {
	s.AddFailuref(pkg.Meta.Pos, "sample visitor: given=%v", pkg.Name)
	return true
}
