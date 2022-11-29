package main

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"github.com/yoheimuta/protolint/linter/report"
	"github.com/yoheimuta/protolint/linter/visitor"
)

type RequiredOptionRule struct {
	requiredOptions []string
}

func (r *RequiredOptionRule) ID() string {
	return "REQUIRED_OPTIONS"
}

func (r *RequiredOptionRule) Purpose() string {
	return "Ensure that the required file options exist"
}

func (r *RequiredOptionRule) IsOfficial() bool {
	return false
}

func (r *RequiredOptionRule) Apply(proto *parser.Proto) ([]report.Failure, error) {
	return visitor.RunVisitor(&requiredOptionVisitor{
		requiredOptions: r.requiredOptions,
		BaseAddVisitor:  visitor.NewBaseAddVisitor(r.ID()),
	}, proto, r.ID())
}
