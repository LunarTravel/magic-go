# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# 依赖缓存
COPY go.mod go.sum ./
RUN go mod download

# 编译
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /bin/magic-service .

# Runtime stage
FROM alpine:3.20

RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Shanghai

WORKDIR /app
COPY --from=builder /bin/magic-service .
COPY configs/ configs/

EXPOSE 9501

ENTRYPOINT ["./magic-service"]
