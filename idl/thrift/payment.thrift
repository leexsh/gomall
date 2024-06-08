namespace go gomall.payment

struct CreditCardInfo {
    1: string credit_card_number
    2: i32 credit_card_cvv
    3: i32 credit_card_expiration_year
    4: i32 credit_card_expiration_month
}

struct ChargeReq {
    1: double amount
    2: CreditCardInfo credit_info
    3: string order_id
    4: i32 user_id
}

struct ChargeResp {
    1: string transaction_id
}

service PaymentService {
    ChargeResp Charge(1:ChargeReq req)    
}