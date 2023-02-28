# 个人网站服务端


## 打包部署
## 打包镜像

```
sh setup.sh
```

## 推送镜像

```
sh push.sh
```

# 交叉编译

## for Linux

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
```

## for windows

```
CGO_ENABLE=0 GOOS=windows GOARCH=amd64 go build
```