include "common.thrift"
namespace go frontend.auth

struct LoginRequest {
    1: string email(api.form="email")
    2: string password(api.form="password")
}

service AuthService {
    common.Empty login(1: LoginRequest req)(api.post="/auth/login")
}