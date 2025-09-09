FROM golang:1.23-alpine AS builder

WORKDIR /app

# 复制 go mod 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -o debank-scheduler ./cmd/scheduler

FROM alpine:latest

WORKDIR /app

# 从 builder 阶段复制二进制文件
COPY --from=builder /app/debank-scheduler .
COPY --from=builder /app/configs ./configs
COPY --from=builder /app/data ./data
COPY --from=builder /app/go.mod ./go.mod
COPY --from=builder /app/go.sum ./go.sum

# 运行应用
CMD ["./debank-scheduler"]