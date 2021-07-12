Nội dung chính
- [Docker và các khái niệm cơ bản](#docker-và-các-khái-niệm-cơ-bản)
- [Dockerfile](#dockerfile)
- [Docker compose](#docker-compose)
- [Các lệnh cơ bản trong docker](#các-lệnh-cơ-bản-trong-docker)
- [Triển khai ứng dụng chat với docker](#triển-khai-ứng-dụng-chat-với-docker)
  - [1. Triển khai thực hiện với Docker Compose](#1-triển-khai-thực-hiện-với-docker-compose)
  - [2. Triển khai với Dockerfile và Docker Compose](#2-triển-khai-với-dockerfile-và-docker-compose)
  - [3. Thực hành với docker file và push lên docker hub](#3-thực-hành-với-docker-file-và-push-lên-docker-hub)
# Docker và các khái niệm cơ bản
![Tổng quan về docker](https://topdev.vn/blog/wp-content/uploads/2020/03/1_8-Gor5TMPqvkeNwSP3LmFw.png)

- **Docker Client**: Là terminal của chính người dùng, Docker Client thực hiện nhận yêu cầu của người dùng thông qua command trong terminal để gửi lệnh tới Docker Daemon.
- **Docker Daemon**: là server dùng cho yêu cầu từ docker API. Nó quản lý các images, containners, volume.
- **Docker Volume**: Là một thư mục, reposi,... được ánh xạ từ máy tính thật vào trong môi trường của containers cho phép containers linh hoạt trong việc lưu và tiếp nhận dữ liệu.
- **Docker Registry**: Là nơi lưu giữ riêng của Dockers Images. Images được push vào Registry và client sẽ pull images từ registry. Có nhiều loại registry như: local, Google Cloud, AWS,..., 
- **Docker Hub**: là Registry lớn nhất của Docker Images. Đặc biệt là free :D
- **Docker Repository**: Là images cùng tên những khác tags. VD: golang:1.11-alpine, golang:latest,...
- **Docker Networking**: cho phép kết nối nhiều container với nhau. Kết nối này có thể ở trong một máy local hoặc có thể ở trên nhiều máy tính với nhau. Có thể được khai báo trong docker-compose.yml
- **Docker Compose**:  công cụ cho phép run app với nhiều Docker containers 1 cách dễ dàng hơn. Docker Compose cho phép bạn config các command trong file docker-compose.yml để sử dụng lại. 
- **Docker Swarm**: để phối hợp triển khai container.
- **Docker Services**: là các containers trong production. 1 service chỉ run 1 images nhưng nó mã hóa cách thức để run image - sử dụng port nào, bao nhiêu bản sao container run để service có hiệu năng cần thiết và ngay lập tức.

# Dockerfile
![Tổng quan về Dockerfile](https://blog.cloud365.vn/images/img-docker/docker4/Dockerfile.png)

Dockerfile là file config để thực hiện build ra một images. Nó dùng một images cơ bản để xây dựng lớp images ban đầu. Một số image cơ bản: python, unbutu and alpine. Sau đó nếu có các lớp bổ sung thì nó được xếp chồng lên lớp cơ bản. Cuối cùng một lớp mỏng có thể được xếp chồng lên nhau trên các lớp khác trước đó.
Các config :
- `FROM` — chỉ định image gốc: python, unbutu, alpine…
- `LABEL` — cung cấp metadata cho image. Có thể sử dụng để add thông tin maintainer. Để xem các label của images, dùng lệnh docker inspect.
- `ENV` — thiết lập một biến môi trường.
- `RUN` — Có thể tạo một lệnh khi build image. Được sử dụng để cài đặt các package vào container.
- `COPY` — Sao chép các file và thư mục vào container.
- `ADD` — Sao chép các file và thư mục vào container.
- `CMD` — Cung cấp một lệnh và đối số cho container thực thi. Các tham số có thể được ghi đè và chỉ có một CMD.
- `WORKDIR` — Thiết lập thư mục đang làm việc cho các chỉ thị khác như: RUN, CMD, ENTRYPOINT, COPY, ADD,…
- `ARG` — Định nghĩa giá trị biến được dùng trong lúc build image.
- `ENTRYPOINT` — cung cấp lệnh và đối số cho một container thực thi.
- `EXPOSE` — khai báo port nào của máy thật thông với port nào của images.
- `VOLUME` — tạo một điểm gắn thư mục để truy cập và lưu trữ data.

Các lệnh trong Dockerfile:
- `docker build [OPTIONS] PATH | URL | -`: lệnh để build một images tại đường dẫn path 
- `docker build -t <image_name>:<version> .`: build images tại Dockerfile tại ví trí `pwd`. Nếu trường version trống, docker tự hiểu đây là `latest`
- `docker build -t <image_name>:<version> -f <path>`: build docker images tại bất kì Dockerfile nào nằm trong path mà cờ `-f` trỏ vào.  

# Docker compose 
![Docker compose](https://images.viblo.asia/997828e9-4dd9-4193-94bc-a7c7886b78f9.png)

Docker compose là công cụ dùng để định nghĩa và run multi-container cho Docker application. Với compose bạn sử dụng file YAML để config các services cho application của bạn. Sau đó dùng command để create và run từ những config đó.
Sử dụng cũng khá đơn giản chỉ với ba bước:
- Khai báo app’s environment trong Dockerfile.
- Khai báo các services cần thiết để chạy application trong file docker-compose.yml.
- Run docker-compose up để start và run app.
Nếu như Dockerfile dùng để tạo ra các images thì Dockercompose là thực hiện việc kết hợp các containers với nhau để có thế thực hiện được một service cụ thể nào đấy.
Docker compose thường được viết theo cú pháp của 2 phiên bản phổ biến là 2 và 3, chi tiết xem tại:
- [Docker compose ver 2](https://docs.docker.com/compose/compose-file/compose-file-v2/)
- [Docker compose ver 3](https://docs.docker.com/compose/compose-file/)
- Các lệnh hay gặp: 
  - `docker-compose up`: Chạy file `docker-compose.yml` nằm ở `pwd`
  - `docker-compose up && docker-compose rm -fsv`: Chạy file `docker-compose.yml` nằm ở `pwd` và xóa các containers liên quan khi dừng lại.
  - `docker-compose up --remove-orphans`: Chạy file `docker-compose.yml` nằm ở `pwd` và thực hiện xóa bỏ đi các containers không còn dùng trong file config
  - `docker-compose down`: dừng các containers đang chạy dựa trên file config nằm ở `pwd`
  - `docker-compose exec [options] [-e KEY=VAL...] SERVICE COMMAND [ARGS...]`: thực thi lệnh `SERVICE COMMAND` ở trong các service mà mình chạy
  - `docker-compose up --abort-on-container-exit`: dừng tất cả các container nếu một container bị dừng
  - `docker-compose up --build $(<services_name_1> <services_name_2>)`: Thực hiện build và chạy chạy các service nằm trong file .yml tại `pwd` mà ta khai báo 
  - `docker-compose up -d`: Chạy file `docker-compose.yml` nằm ở `pwd` trong chế độ background
  - `docker-compose up -f <my_custom_docker_compose>.yml`: chạy một file có tên được đặt theo tên người sử dụng muốn hoặc một file nằm trong đường dẫn mà người đặt muốn.

# Các lệnh cơ bản trong docker
- List image:
> $ docker images
- List container:
> $ docker ps (List tất cả các container đang chạy)\
> $ docker ps –a (List tất cả các container)
- Xóa container:
> $ docker rm <tên/ID container>\
> $ docker rm $(docker ps -a -q)  (Xóa tất cả các container)
- Xóa image:
> $ docker rmi <tên/ID image>
- Dừng container:
> $ docker stop <tên/ID image>
- Chạy containers:
> $ docker run –name <tên_container> -v <thư mục trên máy tính>:<thư mục trong container> -p<port_máy tính>:<port_container>  <image name> bash
- Tải images từ Docker Hub:
> $ docker pull <name_image:tag>
- Start một container(đã stop):
> $ docker start <container_id/name_container>
- Show log containers:
> $ docker logs <container_id/name_container>
- Build một images từ container: 
> $ docker build -t <container_id/name_container> .
- Tạo một container chạy ngầm:
> $ docker run -d <name_image:tag>
- Chạy lệnh trong container ở Docker Client:
> $ docker exec <tên/ID container> <lệnh muốn chạy>
- Truy cập vào 1 container đang chạy
> docker exec -it <container_id/name_container> bash\
> docker attach <container_id/name_container> (có thể không truy cập được trong một số trường hợp cụ thể)
- Export bản container
> docker export <container_id/name_container> | gzip > file_export.tar.gz
- Import container => image
> zcat file_export.tar.gz | docker <new_name_image>
- Help
> docker -h
- Xoá các layer thừa nằm trong docker images khi build:
> docker image prune

Một số cờ hay dùng trong Docker run:
- `--name`: tạo tên cho container
- `-it`: nhảy vào container sau khi chạy
- `--rm`: xóa container khi stop
- `-v`: tham chiếu thư mục trên máy tính vào thư mục trong containers
- `-e`: viết tắt cho env, khởi tạo các biến môi trường cho containers
- ...

# Triển khai ứng dụng chat với docker
Trong phần này em sẽ trình bày việc triển khai ứng dụng chat của mình với docker. Chi tiết như sau:
## 1. Triển khai thực hiện với Docker Compose 
Trong phần này em sẽ thực hiện việc triển khai với Docker Compose mà không dùng Dockerfile để tối ưu cho các images em sử dụng trong hệ thống. Tất cả các images em dùng đều là mặc định. 
Hệ thống của docker em dùng sẽ có dạng như sau: 

![Sơ đồ hệ thống triển khai bằng docker](../Chat(design)/image/Chat_system_in_Docker.png)

Hệ thống của em có 3 service là RESTful server, Websocket Server và MySQL. Trong đó 2 server được nối ra ngoài local bằng cồng 8000 và 8080 như trong hình, MySql không cần phải kết nối với bên ngoài nên chỉ cần tạo một mạng và kết nối nội bộ với cấu hình là Bridge. Dựa vào cấu trúc của hệ thống, ta có file config như sau:
```yaml
version: "3"

networks: 
  my-network:
    driver: bridge

services:
  database: 
    image: mysql:5.7
    hostname: mysql
    container_name: mysql
    environment:
      MYSQL_ROOT_PASSWORD: root
      DATABASE_USER_NAME: t
      DATABASE_PASSWORD: 1
      MYSQL_DATABASE: testdb
    volumes:
      - ./db_data:/var/lib/mysql
      - ./schema.sql:/docker-entrypoint-initdb.d/schema.sql
    networks: 
      - my-network
    healthcheck:
      test: mysqladmin ping -h 127.0.0.1 -u $$MYSQL_USER --password=$$MYSQL_PASSWORD
      timeout: 20s
      retries: 1

  restful-server:
    image: golang:1.15-alpine
    container_name: restful-server
    volumes:
      - ./resful-server:/src
    working_dir: /src
    ports:
      - 8080:8080
    networks: 
      - my-network
    depends_on: 
      "database":
        condition: service_healthy
    command: sh -c "go build -o restful-server; ./restful-server"

  websocket-server:
    image: golang:1.15-alpine
    container_name: websocket-server
    volumes: 
      - ./ws-server:/src
    working_dir: /src
    ports: 
      - 8000:8000
    networks: 
      - my-network
    depends_on: 
      "database":
        condition: service_healthy
    command: sh -c "go build -o websocket-server; ./websocket-server"
```

Thực hiện chạy chỉ cần:
> git clone http://10.1.110.10:8088/nguyentrungkien/kien-training/tree/17-tri-n-khai-ng-d-ng-chat-d-n-gi-n-b-ng-docker/DockerDeploy#1-tri%E1%BB%83n-khai-th%E1%BB%B1c-hi%E1%BB%87n-v%E1%BB%9Bi-docker-compose

> docker-compose up && docker-compose rm -fsv

## 2. Triển khai với Dockerfile và Docker Compose
Nếu ở phần 1 em chỉ gọi đến 2 file là `docker-compose.yml` như và file `schema.sql` đóng vai trò là một file khởi tạo các database cho mysql. Một thư mục để lưu trữ các bảng và cơ sở dữ liệu. Trong phần 2 này em sẽ đóng gói 2 server là của mình lại trở thành các images để có thể thực hiện việc khởi tạo một môi trường cũng như tốt ưu tài nguyên mà các container chiếm.
```yaml
version: "3"

networks: 
  my-network:
    driver: bridge

services:
  database: 
    image: mysql:5.7
    hostname: mysql
    container_name: mysql
    environment:
      MYSQL_ROOT_PASSWORD: root
      DATABASE_USER_NAME: t
      DATABASE_PASSWORD: 1
      MYSQL_DATABASE: testdb
    volumes:
      - ./db_data:/var/lib/mysql
      - ./schema.sql:/docker-entrypoint-initdb.d/schema.sql
    networks: 
      - my-network
    healthcheck:
      test: mysqladmin ping -h 127.0.0.1 -u $$MYSQL_USER --password=$$MYSQL_PASSWORD
      timeout: 20s
      retries: 3

  restful-server:
    container_name: restful-server
    build: 
      dockerfile: Dockerfile
      context: ./resful-server
    hostname: restful-server
    restart: always
    ports:
      - 8080:8080
    networks: 
      - my-network
    depends_on: 
      "database":
        condition: service_healthy

  websocket-server:
    container_name: websocket-server
    build: 
      dockerfile: Dockerfile
      context: ./ws-server
    hostname: websocket-server
    restart: always
    ports: 
      - 8000:8000
    networks: 
      - my-network
    depends_on: 
      "database":
        condition: service_healthy
```

```Dockerfile
# Docker file from websocket server
FROM golang:1.15 AS builder
WORKDIR /go/src/ws-server/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ws-server .

FROM alpine:latest
WORKDIR /root/ws-server
COPY --from=builder /go/src/ws-server .
CMD ["./ws-server"]
```

```Dockerfile
# Docker file from restful server
FROM golang:1.15 AS builder
WORKDIR /go/src/resful-server/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o resful-server .

FROM alpine:latest
WORKDIR /root/resful-server
COPY --from=builder /go/src/resful-server .
CMD ["./resful-server"]
```

Để chạy cũng tương tự như phần 1 ta chỉ cần clone từ git về và thực hiện `docker-compose up`

## 3. Thực hành với docker file và push lên docker hub
1. Đăng kí tài khoản trên docker-hub, sau đó về máy local chạy lệnh
`docker login` trên terminal
2. Viết dockerfile cho việc build image resful-server:
```Dockerfile
# Docker file from restful server
FROM golang:1.15 AS builder
WORKDIR /go/src/resful-server/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o resful-server .

FROM alpine:latest
WORKDIR /root/resful-server
COPY --from=builder /go/src/resful-server .
CMD ["./resful-server"]
```
3. Build docker image với tên quy ước `<tên đăng nhập>/<tên images>:<tag>`

Chạy lệnh 
> `docker build -t <image_name>:<version> .`

Nếu đã có image trong hệ thống và không muốn build lại, ta có thể thực hiện đặt thêm tag mới cho image bằng lệnh:
> docker tag `<tên cũ>:<tag cũ>` `<tên mới>:<tag mới>`

Trong trường hợp cụ thể của em:
> docker tag resful-server:v1 ntkien4docker/resful-server:v1

4. `docker push <image_name>:<version>` để thực đấy lên repo docker-hub
