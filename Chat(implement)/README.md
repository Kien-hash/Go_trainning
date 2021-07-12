# Thực hiện cài đặt môi trường dependencies
## Golang 
- Cài đặt ngôn ngữ Golang thống qua Ubuntu Software 
![](image/Screenshot%20from%202020-09-07%2017-02-15.png)

- Cài đặt GOPATH, biến môi trường để chạy các project golang thông qua câu lệnh:
> go env -w GOPATH=<tên đường dẫn để workspace Golang>

- Trong đường dẫn của GOPATH, tạo folder "src" để chứa code của 2 project ta sẽ chạy: ws-server & resful-server
- Các thư viện ngoài được dùng trong chương trình
  - github.com/gorilla/mux
  - github.com/gorilla/websocket
  - github.com/go-sql-driver/mysql

- Muốn chương trình chạy thành công ta cần cài đặt các thư viện trên, thực hiện lệnh:
   > go get <tên thư viện>

## Mysql Server
- Cài đặt hệ quản trị cơ sở dữ liệu mySQL thông qua các lệnh 
   > sudo apt update \
   > sudo apt install mysql-server

- Truy cập vào mysql để chỉnh chế độ đăng nhập của user root từ auth_socket sang mysql_native_password
  > sudo mysql;\
  > SELECT user,authentication_string,plugin,host FROM mysql.user;\
  > ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';

- Thoát khỏi mysql
  > exit

## MySQL Workbench
- Là một GUI cho phép đễ dàng thao tác với mysql trong máy 
- Tương tự như Golang, tìm và cài đặt nó tại Ubuntu Software
- Trong chương trình của này ta cần 3 bảng trong mySQL, chúng được khỏi tạo như sau:
```mysql
CREATE TABLE IF NOT EXISTS Users (
id INT(6) UNSIGNED AUTO_INCREMENT ,
username VARCHAR(20) NOT NULL unique key,
passwords varchar(30) NOT NULL,
firstname VARCHAR(30) NOT NULL,
lastname VARCHAR(30) NOT NULL,
PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS Rooms (
id INT(6) UNSIGNED AUTO_INCREMENT ,
UsersID VARCHAR(30) NOT NULL unique key,
PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS Message (
id INT(6) UNSIGNED AUTO_INCREMENT ,
RoomID INT(6) UNSIGNED,
SenderID INT(6) UNSIGNED,
MsgContent VARCHAR(100),
PRIMARY KEY (id),
FOREIGN KEY (RoomID) REFERENCES Rooms(id),
foreign key (SenderID) REFERENCES Users(id)
);
```

# Build/compile từ source
- Đưa 2 folder ws-server & resful-server vào trong thư mục src của GOPATH
- Mở 2 cửa sổ terminal trong 2 thư mục ws-server & resful-server
- Trên mỗi terminal chạy lệnh: 
  > go run main.go

# GUI
- Truy cập vào [localhost:8080](localhost:8080) để thực hiện xem giao diện

# Server log 
- Server sẽ log ra các thông tin về các tin nhắn được gửi cũng như các tin nhắn mà server trả về cho client
- Log được thể hiện trong terminal