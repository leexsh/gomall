# 基于hertz和kitex的GOMALL(商城)

## frontend
```
cd app/frontend
go run .
```


## 目录说明
```
rpc_gen: rpc client generate code
app/frontend: http support by hertz
app/user: rpc support by kitex
```

## 问题
```azure
1. 在远端用consul作为服务注册和服务发现，但是失败，是因为远端的server无法ping通wsl，只能用本地的consul。
```