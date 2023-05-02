FROM golang:1.17.0-alpine AS builder
WORKDIR /app
ADD . ./
RUN go env -w GOPROXY=https://goproxy.cn
RUN go env -w GO111MODULE=on
RUN go mod download
RUN go build -o ./sweetpotato-highlightexam-backend

# 分段构建，image大小从560M到25M
FROM alpine
WORKDIR /app
COPY --from=builder /app/sweetpotato-highlightexam-backend /app/sweetpotato-highlightexam-backend
ENTRYPOINT ["/apps/sweetpotato-highlightexam-backend", "run --config /apps/config.yaml"]
