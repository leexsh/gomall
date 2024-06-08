package email

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	email "gomall/rpc_gen/kitex_gen/frontend/email"
)

func Send(ctx context.Context, req *email.EmailReq, callOptions ...callopt.Option) (resp *email.EmailResp, err error) {
	resp, err = defaultClient.Send(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Send call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
