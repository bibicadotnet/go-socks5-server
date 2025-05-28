# SOCKS5 Proxy Server với log tắt mặc định

Máy chủ SOCKS5 đơn giản dựa trên go-socks5 với:
- Xác thực người dùng
- Danh sách IP được phép
- Lọc FQDN đích
- **TẮT TOÀN BỘ LOG HỆ THỐNG MẶC ĐỊNH**

## Cách sử dụng

### Chạy Docker container

1. **Với xác thực đơn giản**:
   ```bash
   docker run -d --name socks5 -p 1080:1080 \
     -e PROXY_USER=<TÊN_NGƯỜI_DÙNG> \
     -e PROXY_PASSWORD=<MẬT_KHẨU> \
     bibica/go-socks5-server-silent
   ```

2. **Không yêu cầu xác thực**:
   ```bash
   docker run -d --name socks5 -p 1090:9090 \
     -e PROXY_PORT=9090 \
     bibica/go-socks5-server-silent
   ```

3. **Xác thực nhiều người dùng**:
   ```bash
   docker run -d --name socks5 -p 1080:1080 \
     -e PROXY_CREDENTIALS='[{"username":"USER1","password":"pass1"},{"username":"USER2","password":"pass2"}]' \
     bibica/go-socks5-server-silent
   ```

## Danh sách cấu hình hỗ trợ

| Biến môi trường       | Kiểu dữ liệu | Mặc định | Mô tả |
|-----------------------|-------------|----------|-------|
| DIAL_TIMEOUT          | String      | 3s       | Thời gian chờ kết nối |
| PROXY_CREDENTIALS     | JSON        | EMPTY    | Danh sách user/password dạng JSON |
| PROXY_USER            | String      | EMPTY    | Tên người dùng (yêu cầu PROXY_PASSWORD) |
| PROXY_PASSWORD        | String      | EMPTY    | Mật khẩu xác thực |
| PROXY_PORT            | String      | 1080     | Cổng lắng nghe trong container |
| ALLOWED_DEST_FQDN     | String      | EMPTY    | Regex cho phép FQDN đích |
| ALLOWED_IPS           | String      | EMPTY    | Danh sách IP được phép kết nối, phân cách bằng dấu phẩy |

**Lưu ý đặc biệt**: Phiên bản silent này đã tắt toàn bộ log hệ thống mặc định để đảm bảo hoạt động tối ưu và bảo mật.

## Xây dựng image tùy chỉnh

```bash
docker-compose up --build -d
```
Cấu hình các tham số trong file .env khi cần thiết

## Kiểm tra hoạt động

1. **Không xác thực**:
   ```bash
   curl --socks5 <IP_DOCKER_HOST>:1080 https://ifcfg.co
   ```
   hoặc
   ```bash
   docker run --rm curlimages/curl:7.65.3 -s --socks5 <IP_DOCKER_HOST>:1080 https://ifcfg.co
   ```

2. **Có xác thực**:
   ```bash
   curl --socks5 <IP_DOCKER_HOST>:1080 -U <USER>:<MẬT_KHẨU> http://ifcfg.co
   ```
   hoặc
   ```bash
   docker run --rm curlimages/curl:7.65.3 -s --socks5 <USER>:<MẬT_KHẨU>@<IP_DOCKER_HOST>:1080 http://ifcfg.co
   ```

## Tác giả

- **Sergey Bogayrets** (Phiên bản gốc)
  
## Người đóng góp
  
- **[bobpaul](https://github.com/bobpaul/go-socks5-server)**

Xem thêm danh sách [người đóng góp](https://github.com/bibica/go-socks5-server-silent/graphs/contributors) cho dự án này.
