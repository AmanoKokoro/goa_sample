package goasimplesample

import (
	"context"
	"log"
	actions "sample/gen/actions"
)

// actions service example implementation.
// The example methods log the requests and return zero values.
type actionssrvc struct {
	logger *log.Logger
}

// NewActions returns the actions service implementation.
func NewActions(logger *log.Logger) actions.Service {
	return &actionssrvc{logger}
}

// サーバーとの導通確認
func (s *actionssrvc) Ping(ctx context.Context) (res string, err error) {
	s.logger.Print("actions.ping")
	return "OK", err
}

// 挨拶する
func (s *actionssrvc) Hello(ctx context.Context, p *actions.HelloPayload) (res string, err error) {
	s.logger.Print("actions.hello")
	name := &p.Name
	res = "Hello" + *name
	return
}

// 複数アクション（:id）
func (s *actionssrvc) ID(ctx context.Context, p *actions.IDPayload) (res int, err error) {
	s.logger.Print("actions.ID")
	if *p.ID == 0 {
		return 1, ctx.Err()
	}

	return *p.ID, nil
}
