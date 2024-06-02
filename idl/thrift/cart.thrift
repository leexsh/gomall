namespace go gomall.cart

struct CartItem {
    1: i32 product_id
    2: i32 quantity
}

struct AddItemReq {
    1: i32 user_id
    2: CartItem item
}

struct AddItemResp {

}

struct GetCartReq {
    1: i32 user_id
}

struct GetCartResp {
    1: list<CartItem> items
}

struct EmptyCartReq {
    1: i32 user_id
}

struct EmptyCartResp {
}

service CartService {
    AddItemResp AddItem(1: AddItemReq req)
    GetCartResp GetCart(1: GetCartReq req)
    EmptyCartResp EmptyCart(1:EmptyCartReq req)
}