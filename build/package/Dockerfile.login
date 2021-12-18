FROM golang:1.15.2-alpine

# 环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

RUN mkdir /xlddz
WORKDIR /xlddz

COPY . .
RUN go build ./cmd/login/

EXPOSE 10010

FROM scratch as login
COPY --from=0 /xlddz /
CMD ["./login","/DockerRun","1"]