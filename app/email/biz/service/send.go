package service

import (
	"context"
	email "gomall/app/email/kitex_gen/frontend/email"
)

type SendService struct {
	ctx context.Context
} // NewSendService new SendService
func NewSendService(ctx context.Context) *SendService {
	return &SendService{ctx: ctx}
}

// Run create note info
func (s *SendService) Run(req *email.EmailReq) (resp *email.EmailResp, err error) {
	// Finish your business logic.

	return
}
