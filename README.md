# go-project-layout-template

docs: https://www.zhangjiee.com/topic/go-project/project-layout.html :cn:

```
git clone https://github.com/zhangjie2012/go-project-layout-template.git
```

run:

``` sh
➜  go-project-layout-template git:(master) make build && ./bin/myapp -conf ./configs/app.yaml
go build -o bin/myapp ./cmd/main.go
INFO[2021-04-28 17:19:49.355]server.go:85 Run server run on: 0.0.0.0:8080

➜  ~ curl http://localhost:8080/api/v1/users
"ok"%
```

docker:

```
make build-docker
make push-docker
```
