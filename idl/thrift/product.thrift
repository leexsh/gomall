namespace go gomall.product

struct ListProductReq {
    1: i32 page
    2: i32 page_size
    3: string category_name
}

struct Product {
    1: i32 id
    2: string name
    3: string description
    4: string picture
    5: double price
    6: list<string> categories
}

struct ListProductResp {
    1: list<Product> products
}

struct GetProductReq {
    1: i32 id
}

struct GetProductResp {
    1: Product product
}
struct SearchProductReq {
    1: string query
}

struct SearchProductResp {
    1:list<Product> results
}

struct GetAllProductsReq {
}

struct GetAllProductsResp {
    1:list<Product> results
}
service ProductService {
    ListProductResp ListProducts(1: ListProductReq req)
    GetProductResp GetProduct(1: GetProductReq req)
    SearchProductResp SearchProducts(1:SearchProductReq req)
    GetAllProductsResp getAllProducts(1:GetAllProductsReq req)
}