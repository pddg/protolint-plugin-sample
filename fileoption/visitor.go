package main

import (
	"sync"

	"github.com/yoheimuta/go-protoparser/v4/parser"
	"github.com/yoheimuta/protolint/linter/visitor"
)

type requiredOptionVisitor struct {
	*visitor.BaseAddVisitor

	requiredOptions []string

	mutex   sync.Mutex
	visited map[string]bool
	meta    *parser.ProtoMeta
}

func (r *requiredOptionVisitor) OnStart(proto *parser.Proto) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	// 開始時に訪問済みのfile optionのmapを初期化する
	r.visited = make(map[string]bool)
	for _, opt := range r.requiredOptions {
		r.visited[opt] = false
	}
	// 最後にAddFailureするときにファイルの情報が欲しいのでメタデータを持っておく
	r.meta = proto.Meta
	return nil
}

func (r *requiredOptionVisitor) VisitOption(opt *parser.Option) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	// 訪れたfile optionは全てtrueにする。
	// 初期化時にfalseにされ、ファイル内に記述されていないものはfalseのまま残る。
	r.visited[opt.OptionName] = true
	return true
}

func (r *requiredOptionVisitor) Finally() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for opt, exists := range r.visited {
		// 記述が無く必至なfile optionがあった場合、違反を報告する
		if !exists {
			r.AddFailurefWithProtoMeta(r.meta, "'%s' is required", opt)
		}
	}
	return nil
}
