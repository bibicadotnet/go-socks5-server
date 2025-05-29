# SOCKS5 Proxy Server với logs tắt mặc định

Máy chủ SOCKS5 đơn giản dựa trên go-socks5 với:
- Xác thực người dùng
- Danh sách IP được phép
- Lọc FQDN đích
- **TẮT TOÀN BỘ LOG HỆ THỐNG MẶC ĐỊNH**

Có rất nhiều phiên bản socks5 trên Github, đặc biệt là các phiên bản của người Trung Quốc viết ra, hiệu năng bố đời, chạy cực nhẹ, chịu tải cao

Ở góc độ người dùng cuối tại Việt Nam, thường mục đích chính mở chặn các trang bị [nhà mạng khóa](https://bibica.net/giai-quyet-telegram-bi-nha-mang-viet-nam-chan-bang-mtproto-socks5-proton-vpn/), dùng các phiên bản socks5 đơn giản, **hỗ trợ xác thực người dùng, tùy chỉnh các port được là đủ**, đỡ phải mất thời gian tìm hiểu nhiều

Lượn lờ thì mình thấy bản `go-socks5-proxy` từ `serjs` có lượt kéo về hàng đầu trên [docker hub](https://hub.docker.com/r/serjs/go-socks5-proxy) (hơn 10 triệu lượt)

Việc cài đặt, xác thực người dùng, tùy chỉnh các port, sử dụng đơn giản như mong đợi, thứ duy nhất khó chịu, là tác giả vẫn duy trì, giữ lại 1 số thông báo `[INFO]` từ hệ thống (https://github.com/serjs/socks5-server/issues/68), `2025/05/28 10:57:54 [INFO] socks: Connection from allowed IP address: 212.179.155.163`

Ngoài chuyện nó lưu lại ngày giờ và IP kết nối tới socks, tần xuất logs `[INFO]` này xuất hiện ở cường độ rất cao, tầm 1s/1 lần, logs rác như thế không hiểu sao tác giả không xóa đi cho nhẹ VPS?

Phiên bản bạn thấy ở đây, dùng từ bản `bobpaul/go-socks5-server`, sau đó xóa sạch tất cả các logs hệ thống, còn lại cũng chẳng biết gì mà sửa, cấu hình, sử dụng, tương tư phiên bản gốc

- Cấu hình sử dụng nhanh qua docker `compose.yml`

```compose.yml
services:
  socks5:
    image: bibica/go-socks5-server-silent:latest  # Image SOCKS5
    container_name: socks5                        # Tên container
    restart: always                              # Tự động khởi động lại
    environment:
      - PROXY_USER=myusername                    # Username SOCKS5
      - PROXY_PASSWORD=mypassword                # Password SOCKS5
      - PROXY_PORT=7128                          # Port chạy trong container
    ports:
      - "12821:7128"                             # Port host → container

```
Đổi lại thông tin tài khoản `myusername` `mypassword` hoặc port `12821` thành 1 số khác tùy thích

- Bật chạy

```
docker compose up -d
```

- Kiểm tra xác thực tài khoản:

```
curl --socks5 myusername:mypassword@localhost:12821 http://ifconfig.me

Hoặc

curl --socks5 myusername:mypassword@localhost:12821 http://httpbin.org/ip

```

Thấy hiện ra được IP của VPS là chính xác

- Kết nối sử dụng

<p align="center">
  <img src="https://raw.githubusercontent.com/bibicadotnet/go-socks5-server-silent/refs/heads/merge_requests/img/2025-05-29_13-40-48.png" alt="Kết nối sử dụng" />
</p>

- Tùy chỉnh port

Ví dụ trên mình dùng 1 port khá ngẫu nhiên là `12821`, tránh trùng tới các port có sẵn của hệ thống

Về `PROXY_PORT=7128` bên trong container, bạn có thể đổi sang bất cứ port nào, không ảnh hưởng gì, miễn là ports mapping đúng (ví dụ: "12821:7128")

Tác giả serjs dùng port SOCKS tiêu chuẩn `1080`, chủ yếu cũng do thói quen, một số người khác thích dùng port HTTPS tiêu chuẩn `443`, vì không hệ thống firewall nào mặc định đi chặn 443 cả, nó cũng tránh được việc soi ra đang dùng SOCKS hơn, có điều `443` thường mọi người hay chạy webserver, dùng 443 trên socks dễ bị trùng, gây lỗi

Theo cấu hình ví dụ ở trên, __mở port `12821` trên VPS là được__

### __Kiểm tra logs hệ thống__
  
```
curl --socks5 myusername:mypassword@localhost:12821 http://ifconfig.me
docker logs socks5
```
- Không thấy logs nào hiện ra là đúng

- Ở phiên bản gốc bạn sẽ thấy khá nhiều logs rác tại đây, bản silent này tắt sẵn, đỡ phải ghi quá nhiều thông tin xuống ổ cứng, gây nặng khi chạy ở thời gian dài, cũng giúp người dùng đỡ lấn cấn vấn đề VPS ghi lại IP của người sử dụng
