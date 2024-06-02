package service

import (
	"context"
	product "gomall/app/product/kitex_gen/gomall/product"
	"testing"
)

func TestSearchProducts_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSearchProductsService(ctx)
	// init req and assert value

	req := &product.SearchProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
