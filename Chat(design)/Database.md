# Thiết kế DB đối với hệ thống
## 1. User table

| Trường    | Kiểu dữ liệu | Mô tả                            |
| --------- | ------------ | -------------------------------- |
| id        | INT(6)       | Số thứ tự của bản ghi trong bảng |
| username  | VARCHAR(20)  | Tên đăng nhập                    |
| passwords | VARCHAR(30)  | Mật khẩu                         |
| firstname | VARCHAR(30)  | Họ                               |
| lastname  | VARCHAR(30)  | Tên                              |

## 2. Chat room

| Trường  | Kiểu dữ liệu | Mô tả                                                                                                  |
| ------- | ------------ | ------------------------------------------------------------------------------------------------------ |
| ID      | INT(6)       | Số thứ tự của bản ghi trong bảng                                                                       |
| UsersID | VARCHAR(30)  | Mảng ID của các user trong phòng được chuyển thành string theo một mẫu có sẵn, sắp xếp thứ tự tăng dần |


## 3. Message

| Trường     | Kiểu dữ liệu | Mô tả                                     |
| ---------- | ------------ | ----------------------------------------- |
| ID         | INT(6)       | Số thứ tự của bản ghi tin nhắn trong bảng |
| MsgID      | INT(6)       | ID của người gửi                          |
| MsgContext | VARCHAR(50)  | Nội dung của tin nhắn                     |
| RoomID     | INT(6)       | Tin nhắn của phòng nào                    |

## Quan hệ giữa các bảng
![Quan hệ giữa các bảng](image/Screenshot%20from%202020-09-15%2016-09-30.png)
