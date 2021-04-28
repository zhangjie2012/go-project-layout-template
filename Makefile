TARGET=myapp
BRANCH=`git rev-parse --abbrev-ref HEAD`
TAG=`git rev-parse --short HEAD`

build:
	go build -o bin/${TARGET} ./cmd/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/${TARGET} ./cmd/main.go

build-docker:
	GOOS=linux GOARCH=amd64 go build -o bin/${TARGET} ./cmd/main.go
	docker build -t ${TARGET}:${BRANCH}-${TAG} .

push-docker:
	docker push ${TARGET}:${BRANCH}-${TAG}
