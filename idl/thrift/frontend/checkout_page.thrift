include "common.thrift"
namespace go frontend.checkout

struct CheckoutReq {
    1: string email(api.form="email")
    2: string firstName(api.form="firstname")
    3: string lastName(api.form="lastname")
    4: string street(api.form="street")
    5: string zipCode(api.form="zipcode")
    6: string province(api.form="province")
    7: string country(api.form="country")
    8: string city(api.form="city")
    9: string card_num(api.form="cardNum")
    10: i32 expiration_month(api.form="expirationMonth")
    11: i32 expiration_year(api.form="expirationYear")
    12: i32 cvv(api.form="cvv")
    13: string payment(api.form="payment")
}

service CheckoutService {
    common.Empty Checkout(1:common.Empty req)(api.get="/checkout")
    common.Empty CheckoutWaiting(1:CheckoutReq req)(api.post="/checkout/waiting")
    common.Empty CheckoutResult(1:common.Empty req)(api.get="/checkout/result")
}