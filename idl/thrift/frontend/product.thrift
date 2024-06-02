include "common.thrift"
namespace go frontend.product

struct ProductReq {
    1: i32 id(api.query="id")
}

struct SearchProductsReq {
    1: string q(api.query="q")
}

service ProductService {
    common.Empty GetProduct(1:ProductReq req)(api.get="/product")
    common.Empty SearchProducts(1:SearchProductsReq req)(api.get="/search")
}