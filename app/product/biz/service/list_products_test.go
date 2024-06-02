package service

import (
	"context"
	product "gomall/app/product/kitex_gen/gomall/product"
	"testing"
)

func TestListProducts_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListProductsService(ctx)
	// init req and assert value

	req := &product.ListProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
