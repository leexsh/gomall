.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --type RPC --service demo_proto --idl ../../idl/proto/echo.proto  

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server --type RPC --service demo_thrift  --idl ../../idl/thrift/echo.thrift

.PHONY: gen-frontend
gen-frontend:
#	 @cd app/frontend && cwgo server -I ../../idl/proto --type HTTP --service frontend --module gomall/app/frontend --idl ../../idl/proto/frontend/home.proto
#	@cd app/frontend && cwgo server -I ../../idl/thrift --type HTTP --service frontend --module gomall/app/frontend --idl ../../idl/thrift/frontend/home.thrift
#	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module gomall/app/frontend --idl ../../idl/thrift/frontend/auth_page.thrift
#	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module gomall/app/frontend --idl ../../idl/thrift/frontend/product.thrift
#	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module gomall/app/frontend --idl ../../idl/thrift/frontend/category.thrift
#	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module gomall/app/frontend --idl ../../idl/thrift/frontend/cart_page.thrift
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module gomall/app/frontend --idl ../../idl/thrift/frontend/checkout_page.thrift

.PHONY: gen-user
gen-user:
	# rpc client
	#@cd rpc_gen && cwgo client --type RPC --service user --module gomall/rpc_gen --I ../idl --idl ../idl/thrift/user.thrift
	# rpc server
	@cd app/user && cwgo server --type RPC --service user --module gomall/app/user --I ../../idl --idl ../../idl/thrift/user.thrift

.PHONY: gen-product
gen-product:
	@cd rpc_gen && cwgo client --type RPC --service product --module gomall/rpc_gen -I ../idl --idl ../idl/thrift/product.thrift
	@cd app/product && cwgo server --type RPC --service product  --module gomall/app/product --I ../../idl --idl ../../idl/thrift/product.thrift

.PHONY: gen-cart
gen-cart:
	@cd rpc_gen && cwgo client --type RPC --service cart --module gomall/rpc_gen -I ../idl --idl ../idl/thrift/cart.thrift
	@cd app/cart && cwgo server --type RPC --service cart  --module gomall/app/cart --I ../../idl --idl ../../idl/thrift/cart.thrift

.PHONY: gen-payment
gen-payment:
	@cd rpc_gen && cwgo client --type RPC --service payment --module gomall/rpc_gen -I ../idl --idl ../idl/thrift/payment.thrift
	@cd app/payment && cwgo server --type RPC --service payment  --module gomall/app/payment --I ../../idl --idl ../../idl/thrift/payment.thrift

.PHONY: gen-checkout
gen-checkout:
	@cd rpc_gen && cwgo client --type RPC --service checkout --module gomall/rpc_gen -I ../idl --idl ../idl/thrift/checkout.thrift
	@cd app/checkout && cwgo server --type RPC --service checkout  --module gomall/app/checkout --I ../../idl --idl ../../idl/thrift/checkout.thrift

.PHONY: gen-email
gen-email:
	@cd rpc_gen && cwgo client --type RPC --service email --module gomall/rpc_gen -I ../idl --idl ../idl/thrift/email.thrift
	@cd app/email && cwgo server --type RPC --service email  --module gomall/app/email --I ../../idl --idl ../../idl/thrift/email.thrift

gen-rpc-client:
#	@cd app/frontend/rpc/userClient && cwgo client --type RPC --service user --module gomall/app/frontend/rpc/userClient --I ../../../../idl --idl ../../../..//idl/thrift/user.thrift
	@cd rpc_gen && cwgo client --type RPC --service user --module gomall/rpc_gen --I ../idl --idl ../idl/thrift/user.thrift


