namespace go gomall.checkout
include "payment.thrift"

struct Address {
    1: string street_address
    2: string city
    3: string state
    4: string country
    5: string zip_code
}

struct CheckoutReq {
    1: i32 user_id
    2: string first_name
    3: string last_name
    4: string email
    5: Address address
    6: payment.CreditCardInfo credit_card
}

struct CheckoutResp {
    1: string order_id
    2: string transaction_id
}

service CheckoutService {
    CheckoutResp Checkout(1:CheckoutReq req)
}