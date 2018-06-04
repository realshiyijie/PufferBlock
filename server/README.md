# 系统操作说明

## 启动`fabric`网络

```bash
./network_setup.sh up
```

如果网络启动失败，重新启动

```bash
./network_setup.sh restart
```

## 启动服务端

```bash
cd $GOPATH/src/PufferBlock/server/
go run server.go
```

## 启动用户界面

打开FireFox浏览器，前往[localhost:8080](http://localhost:8080)。