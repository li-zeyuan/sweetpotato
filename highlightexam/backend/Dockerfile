FROM golang:1.17.0-alpine AS builder
WORKDIR /apps
ADD . ./

RUN go build -o /apps/sweetpotato-highlightexam-backend

# 分段构建，image大小从560M到25M
FROM alpine
WORKDIR /apps
COPY --from=builder /apps/sweetpotato-highlightexam-backend /apps/sweetpotato-highlightexam-backend
RUN chmod +x /apps/sweetpotato-highlightexam-backend
ENTRYPOINT ["/apps/sweetpotato-highlightexam-backend", "run --config /apps/config.yaml"]
