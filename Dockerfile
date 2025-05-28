ARG GOLANG_VERSION="1.19.1"

FROM golang:$GOLANG_VERSION-alpine as builder

# Cài đặt các phụ thuộc cần thiết
RUN apk --no-cache add tzdata git

# Tạo thư mục làm việc
WORKDIR /app

# Copy toàn bộ mã nguồn local (bao gồm go.mod, go.sum và code đã sửa)
COPY . .

# Build ứng dụng với các cờ tối ưu và tắt debug
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-s -w -X main.disableLogs=true" \  # Truyền cờ tắt log qua biến build-time
    -a -installsuffix cgo \
    -o socks5 .

# Giai đoạn runtime
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/socks5 /  # Sửa đường dẫn đúng vị trí build

ENTRYPOINT ["/socks5"]
