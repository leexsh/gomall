include "common.thrift"
namespace go frontend.home

service HomeService {
    common.Empty Home(1:common.Empty req)(api.get="/")
}
