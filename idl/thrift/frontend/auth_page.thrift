include "common.thrift"
namespace go frontend.auth

struct LoginRequest {
    1: string email(api.form="email")
    2: string password(api.form="password")
    3: string next(api.query="next")
}
struct RegisterRequest {
    1: string email(api.form="email")
    2: string password(api.form="password")
    3: string passwordConfirm(api.form="password_confirm")
}

service AuthService {
    common.Empty login(1: LoginRequest req)(api.post="/auth/login")
    common.Empty register(1: RegisterRequest req)(api.post="/auth/register")
    common.Empty logout(1: common.Empty req)(api.post="/auth/logout")
}