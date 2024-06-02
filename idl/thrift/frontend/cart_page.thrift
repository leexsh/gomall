include "common.thrift"
namespace go frontend.cart

struct AddCartItemReq {
    1:i32 product_id(api.form="productId")
    2:i32 product_num(api.form="productNum")
}

service ProductService {
    common.Empty GetCart(1:common.Empty req)(api.get="/cart")
    common.Empty AddCartItem(1:AddCartItemReq req)(api.post="/cart")
}