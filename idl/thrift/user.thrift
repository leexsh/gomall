namespace go gomall.user

struct RegisterReq {
    1:string email
    2:string password
    3:string password_confrim
}

struct RegisterResp {
    1:string userId
}

struct LoginReq {
    1:string email
    2:string password
}
struct LoginResp {
    1:string userId
}

service UserService {
    RegisterResp Register(1:RegisterReq req)
    LoginResp Login(1:LoginReq req)

}