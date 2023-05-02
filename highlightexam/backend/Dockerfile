FROM golang:1.17.0-alpine AS builder
WORKDIR /apps
ADD . ./

RUN go build -o /apps/highlight-exam

# 分段构建，image大小从560M到25M
FROM alpine
WORKDIR /apps
COPY --from=builder /apps/highlight-exam /apps/highlight-exam
RUN chmod +x /apps/highlight-exam
ENTRYPOINT ["/apps/highlight-exam", "run", "--config=/apps/config.yaml"]
