package rpc_client
import "github.com/cloudwego/hertz/pkg/common/hlog"
func MustHandleErr(err error) {
	if err == nil {
		return
	}
	hlog.Fatal(err)
}
