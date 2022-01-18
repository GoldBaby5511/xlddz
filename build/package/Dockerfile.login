FROM golang:1.15.2-alpine

# 环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

RUN mkdir /mango
WORKDIR /mango

COPY . .
RUN go build ./cmd/login/

EXPOSE 10070

FROM scratch as login
COPY --from=0 /mango /
CMD ["./login","-Type=5","-Id=70","-DockerRun=1"]