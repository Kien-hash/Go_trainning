# Hệ thống nhắn tin sử dụng websocket và restfulAPI
## Yêu cầu hệ thống
- Ứng dụng chat viết bằng Golang/nodejs/java/python gồm client + server
- Sử dụng websocket và restful
- Tính năng:
  - Đăng ký người dùng với user name + password
  - Chat giữa 2 người dùng biết user name của nhau
  - Chat nhóm giữa nhiều người biết user name của nhau

## Kiến trúc hệ thống đề xuất
Từ các yêu cầu trên, hệ thống được đề xuất:
![System Architecture](./image/architecture.png)

## Kết nối giữa các khối

| Kết nối | Mô tả                                                                                                            |
| ------- | ---------------------------------------------------------------------------------------------------------------- |
| 1       | Main server kết nối với database nhằm truy xuất các dữ liệu về người dùng như: tên đăng nhập, mật khẩu,...       |
| 2       | Kết nối này nhằm phục vụ truy xuất các tin nhắn trước đó về server trong cuộc hội thoại nhất định                |
| 3       | Client thực hiện các kết nối http theo kiểu restfulAPI nhằm thực hiện các nhiệm vụ đăng nhập, đăng kí, đăng xuất |
| 4       | Liên kết dự phòng                                                                                                |
| 5       | Thực hiện việc gửi tin nhắn và nhận tin nhắn thông qua liên kết websocket từ server lên người dùng               |

## Hệ thống gồm các khối

|          | Main server                                                                                    | Chat server                                                                                                 | Database                                                                   | Client                                                                                                  |
| -------- | ---------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| Nhiệm vụ | Giao tiếp với client nhằm phục vụ những chức năng: đăng kí, đăng nhập,...                      | Duy trì kết nối với các Client đã đăng nhập bằng liên kết websocket => thông báo cho users biết user online | Lưu trữ thông tin về mà 2 server cần: User, meessage, room                 | Cung cấp giao diện với người dùng: Đăng nhập, đăng kí, xem useronline, hiển thị tin nhắn chat đơn, nhóm |
| Thiết kế | Khối Main server được viết bằng golang, thực hiện kết nối với các Client thông qua các RestAPI | Sử dụng golang, sử dụng websocket trong thư viện gorilla websocket                                          | Sử dụng mySQL. Được thiết kế tại [Tài liệu thiết kế database](Database.md) | Sử dụng web app: html, javascript                                                                       |
