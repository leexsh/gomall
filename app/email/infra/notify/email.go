package notify

import (
	"gomall/app/email/kitex_gen/frontend/email"

	"github.com/kr/pretty"
)

type NoopEmail struct {

}

func (n *NoopEmail)Send(req *email.EmailReq) error {
	pretty.Printf("%v\n", req)
	return nil
}

func NewNoopEmail()*NoopEmail  {
	return &NoopEmail{}
}