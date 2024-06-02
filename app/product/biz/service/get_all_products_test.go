package service

import (
	"context"
	product "gomall/app/product/kitex_gen/gomall/product"
	"testing"
)

func TestGetAllProducts_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetAllProductsService(ctx)
	// init req and assert value

	req := &product.GetAllProductsReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
