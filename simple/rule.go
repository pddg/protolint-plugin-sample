package main

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"github.com/yoheimuta/protolint/linter/report"
	"github.com/yoheimuta/protolint/linter/visitor"
)

type SimpleDenyRule struct{}

// ID - アッパースネークケースのIDを返す
func (r *SimpleDenyRule) ID() string {
	return "SIMPLE_DENY_RULE"
}

// Purpose - このルールの目的。
func (r *SimpleDenyRule) Purpose() string {
	return "sample implementation of protolint plugin"
}

// IsOfficial - このルールがGoogleの公式ガイドラインのものかどうか。基本的にfalse。
func (r *SimpleDenyRule) IsOfficial() bool {
	return false
}

// Apply - このルールを対象のprotoファイルに適用する。
func (r *SimpleDenyRule) Apply(proto *parser.Proto) ([]report.Failure, error) {
	return visitor.RunVisitor(&sampleVisitor{
		BaseAddVisitor: visitor.NewBaseAddVisitor(r.ID()),
	}, proto, r.ID())
}
