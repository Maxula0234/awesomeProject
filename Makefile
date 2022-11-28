APP?=testdatagonew
PORT?=8000
COMMIT?=$(shell git rev-parse --short HEAD)
GOOS?=linux
GOARCH?=amd64

clean:
	rm -f ${APP}

build: clean
		CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
        		-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
        		-X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
        		-o ${APP}

container: build
	docker build -t $(APP) .

run: container
	docker stop $(APP) || true && docker rm $(APP)|| true
	docker run --name ${APP} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" \
		$(APP)

test:
	go test -v -r	ace ./...