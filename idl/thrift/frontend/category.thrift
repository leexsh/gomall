include "common.thrift"
namespace go frontend.category
struct CategoryReq {
    1: string category(api.path="category")
}

service CategoryService {
    common.Empty Category(1:CategoryReq req)(api.get="/category/:category")
}