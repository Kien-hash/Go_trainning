# 1. Hệ điều hành 
## 1.1 Tổng quan
Hệ điều hành là một phần mềm chạy trên máy tính và các thiết bị di động, dùng để điều hành, quản lý các thiết bị phần cứng và các tài nguyên phần mềm trên máy tính và các thiết bị di động.
![OS structure](https://i.ytimg.com/vi/vyECnqM9LJM/maxresdefault.jpg)
Nhiệm vụ của hệ điều hành:
* Điều khiển và quản lý trực tiếp các phần cứng như bo mạch chủ, bo mạch đồ họa và bo mạch âm thanh,...
* Thực hiện một số thao tác cơ bản trong máy tính như các thao tác đọc, viết tập tin, quản lý hệ thống tập tin (file system) và các kho dữ liệu.
* Cung ứng một hệ thống giao diện sơ khai cho các ứng dụng thường là thông qua một hệ thống thư viện các hàm chuẩn để điều hành các phần cứng mà từ đó các ứng dụng có thể gọi tới.
* Cung ứng một hệ thống lệnh cơ bản để điều hành máy. Các lệnh này gọi là lệnh hệ thống (system command).
* Ngoài ra hệ điều hành, trong vài trường hợp, cũng cung cấp các dịch vụ cơ bản cho các phần mềm ứng dụng thông thường như chương trình duyệt Web, chương trình soạn thảo văn bản....

## 1.2 Một số hoạt động chính của hệ điều hành
Hoạt động của hệ điều hành rất nhiều nhưng ta có thể khái quát chúng thành những hoạt động chính như sau:
1. Quản lý tiến trình: Tiến trình là một thực thể chương trình đang được thực thi, quản lý tiến trình là quá trình phân bổ tài nguyên hệ thống cho các chương trình lớp ứng dụng sao cho chúng có thể thực hiện được các chức năng mong muốn
2. Lập lịch: Thực hiện chia khe thời gian và phân bổ các tiến trình vào các khe thời gian tùy thuộc vào các chiến lược cụ thể đối với từng hàng đợi.
3. Xứ lý luồng và xử lý đồng bộ: Luồng là một thành phần con được chia nhỏ ra từ tiến trình, giúp giảm thời gian xử lý của một tiến trình, giảm tải nguyên, tận dụng được đa core hiện đại...
4. Quản lý bộ nhớ: Một chương trình muốn chạy được thì phài được cấp phát vùng nhớ, chức năng quản lý bộ nhớ theo dõi, cấp phát, giải phóng, truy xuất,... bộ nhớ để phục vụ cho các tiến trình. 
5. Quản lý I/O: Một trong những công việc quan trọng của OS là quản lý các thiết bị đầu vào / đầu ra (Input / Output) khác nhau, bao gồm chuột, bàn phím, touchpad, ổ đĩa cứng, màn hình, các thiết bị USB, thiết bị kết nối mạng, thiết bị âm thanh, máy in, v.v.
6. Ảo hóa phẩn cứng: (Virtualization) là công nghệ đặc biệt cho phép chúng ta tạo ra nhiều môi trường mô phỏng hoặc tài nguyên chuyên dụng từ một hệ thống phần cứng vật lý duy nhất (Tức là có thể coi như dùng được nhiều máy tính trên một phần cứng duy nhất).
7. Các hoạt động khác: Hệ điều hành còn có một số các hoạt động khác như quản lý tập tin, hệ thống dịch lệnh, giao tiếp liên tiến trình, phân tán tệp tin,...

# 2. Basic Network Concept
## 2.1 IP
IP là viết tắt của Internet Protocol - giao thức Internet là một giao thức hướng dữ liệu được sử dụng bởi các máy trạm đầu cúối để truyền dữ liệu trong một mạng chuyển mạch gói.
Dữ liệu trong một liên mạng IP được tổ chức dưới dạng gói (packet hoặc datagram), được thực hiện đóng gói ở tầng mạng và chuyển xuống tầng liên kết dữ liệu, dữ liệu được chuyển đi theo dịch vụ best effort không đảm bảo tính tin cậy.
![Protocol in layers](https://micrium.com/wp-content/uploads/2014/03/OSI-Seven-Layer-Model.png)
Các máy tính sử dụng giao thức IP phải có một địa chỉ IP để giao tiếp với nhau, địa chỉ IP có thể là đơn nhất trong mạng nội bộ hoặc đơn nhất trong mạng toàn cầu, tùy thuộc vào thiết kế và yêu cầu của hệ thống. Hiện đang có 2 phiên bản của địa chỉ IP là IPv4 sử dụng 32 bit để đánh địa chỉ, và IPv6 sử dụn 128 bit.
### 2.1.1 IPv4
Giao thức Internet phiên bản 4 hiện đang là phiên bản phổ biến nhất trong các giao thức Internet (IP). IPv4 sử dụng 32 bits để đánh địa chỉ, theo đó, số địa chỉ tối đa có thể sử dụng là 4.294.967.296 (2^32). Tuy nhiên, do một số được sử dụng cho các mục đích khác như: cấp cho mạng cá nhân (xấp xỉ 18 triệu địa chỉ), hoặc sử dụng làm địa chỉ quảng bá (xấp xỉ 16 triệu), nên số lượng địa chỉ thực tế có thể sử dụng cho mạng Internet công cộng bị giảm xuống. 

| Khối địa chỉ CIDR | Miêu tả                                                                | Tài liệu tham khảo |
| ----------------- | ---------------------------------------------------------------------- | ------------------ |
| 0.0.0.0/8         | Mạng hiện tại (Chỉ có giá trị với địa chỉ nguồn)                       | RFC 1700           |
| 10.0.0.0/8        | mạng riêng                                                             | RFC 1918           |
| 14.0.0.0/8        | "Mạng dữ liệu công cộng (sẵn sàng cho sử dụng từ 10 tháng 2 năm 2008)" | RFC 1700           |
| 127.0.0.0/8       | Localhost                                                              | RFC 3330           |
| 128.0.0.0/16      | Dự trữ (IANA)                                                          | RFC 3330           |
| 169.254.0.0/16    | Link-Local                                                             | RFC 3927           |
| 172.16.0.0/12     | mạng riêng                                                             | RFC 1918           |
| 191.255.0.0/16    | Dự trữ (IANA)                                                          | RFC 3330           |
| 192.0.0.0/24      | Dự trữ (IANA)                                                          | RFC 3330           |
| 192.0.2.0/24      | Tài liệu và mã ví dụ                                                   | RFC 3330           |
| 192.88.99.0/24    | tương tác với IPv6                                                     | RFC 3068           |
| 192.168.0.0/16    | mạng riêng                                                             | RFC 1918           |
| 198.18.0.0/15     | Dành cho thí nghiệm                                                    | RFC 2544           |
| 223.255.255.0/24  | Dự trữ (IANA)                                                          | RFC 3330           |
| 224.0.0.0/4       | Multicasts (Lớp D)                                                     | RFC 3171           |
| 240.0.0.0/4       | Dự trữ (lớp E)                                                         | RFC 1700           |
| 255.255.255.255   | Broadcast                                                              |

Với sự phát triển không ngừng của mạng Internet, nguy cơ thiếu hụt địa chỉ đã được dự báo, tuy nhiên, nhờ công nghệ NAT (Network Address Translation - Chuyển dịch địa chỉ mạng) tạo nên hai vùng mạng riêng biệt: Mạng riêng và Mạng công cộng, địa chỉ mạng sử dụng ở mạng riêng có thể dùng lại ở mạng công công mà không hề bị xung đột, qua đó trì hoãn được vấn đề thiếu hụt địa chỉ.\
#### Mạng riêng
Trong khoảng hơn bốn tỷ địa chỉ có thể sử dụng của IPv4, ba dải địa chỉ được dành riêng cho các mạng riêng (private network). Các dải này không xuất hiện trong bảng định tuyến ở bên ngoài mạng riêng. Các thiết bị trong mạng riêng cũng không thể trực tiếp liên lạc với các mạng công cộng. Để có thể liên lạc với internet công cộng, họ phải sử dụng công nghệ NAT.

| Tên         | Dải địa chỉ                 | Số lượng địa chỉ | Mô tả mạng đầy đủ            | Khối CIDR lớn nhất |
| ----------- | --------------------------- | ---------------- | ---------------------------- | ------------------ |
| Khối 24-bit | 10.0.0.0–10.255.255.255     | 16.777.216       | Một dải trọn vẹn thuộc lớp A | 10.0.0.0/8         |
| Khối 20-bit | 172.16.0.0–172.31.255.255   | 1.048.576        | Tổ hợp từ mạng lớp B         | 172.16.0.0/16      |
| Khối 16-bit | 192.168.0.0–192.168.255.255 | 65,536           | Tổ hợp từ mạng lớp C         | 192.168.0.0/24     |

#### NAT
Công nghệ "dịch địa chỉ mạng" (NAT: Network Address Translation)có thể nói đơn giản  như một bảng liên kết gồm hai phần, một phần là danh sách các địa chỉ của mạng riêng có nhu cầu giao tiếp với mạng công cộng, phần còn lại là địa chỉ công cộng được chỉ định làm đại diện cho các địa chỉ mạng riêng (tạm gọi là địa chỉ NAT).
![NAT](https://upload.wikimedia.org/wikipedia/commons/thumb/c/c7/NAT_Concept-en.svg/1280px-NAT_Concept-en.svg.png)
Khi một host trong mạng riêng muốn giao tiếp với một host (ví dụ như server) trong mạng công cộng, nó sẽ đóng gói bản tin với địa chỉ nguồn là địa chỉ mạng riêng của chính nó (host Pri, có địa chỉ IP là: 10.0.0.1), địa chỉ đích là địa chỉ của server trong mạng công cộng (Pub Server, có địa chỉ IP là: 200.100.10.1). Gói tin này sẽ được gửi tới NAT, tại đây, toàn bộ gói tin này được đóng gói thêm một lần nữa, địa chỉ đích của gói mới vẫn là 200.100.10.1, nhưng địa chỉ nguồn thì được đổi thành địa chỉ NAT(150.150.0.1). Để phản hồi về host, server đóng gói gói tin với địa chỉ nguồn là 200.100.10.1, địa chỉ đích là 150.150.0.1, sau đó tiếp tục sử dụng "lớp vỏ" NAT đóng gói tiếp một lần nữa để tạo thành gói NAT, gói NAT này có địa chỉ nguồn là 200.100.10.1, địa chỉ đích là địa chỉ 10.0.0.1.
#### Cấu trúc gói tin IPv4
![IPv4 formart](https://upload.wikimedia.org/wikipedia/commons/thumb/7/71/IPv4_Packet_-en.svg/1024px-IPv4_Packet_-en.svg.png)
* ***Phiên bản (Version)***: Trường đầu tiên trong header của gói tin IP chính là trường Phiên bản (Version) dài 4 bit. Với IPv4, nó có giá trị bằng 4.
* ***Độ lớn của header (Internet Header Length) (IHL)***: Trường thứ hai (4 bit) là độ lớn của header (Internet Header Length - IHL) cho biết số lượng các từ 32-bit trong header. 
* ***Differentiated Services (DS)***: Ban đầu được định nghĩa là trường TOS, hiện tại trường này được định nghĩa trong RFC 2474 là Differentiated services (DiffServ) và trong RFC 3168 là Explicit Congestion Notification (ECN), để phù hợp với IPv6. Chỉ định dịch vụ mong muốn khi truyền các gói tin qua router. 
* ***Total Length***: Chỉ định tổng chiều dài gói tin IPv4 (cả phần mào đầu và phần dữ liệu). Kích thước 16 bít, chỉ định rằng gói tin IPv4 nhỏ nhất là 20 byte (chỉ có header không có dữ liệu) và có thể lớn tới 65.535 byte.
* ***Identification***: Định danh gói tin. Kích thước 16 bít. Định danh cho gói tin được lựa chọn bởi nguồn gửi gói tin. Nếu gói tin IPv4 bị phân mảnh, mọi phân mảnh sẽ giữ lại giá trị trường định danh này, mục đích để nút đích có thể nhóm lại các mảnh, phục vụ cho việc phục hồi lại gói tin.

### 2.1.2 IPv6
Do sự phát triển như vũ bão của mạng và dịch vụ Internet, nguồn IPv4 dần cạn kiệt, đồng thời bộc lộ các hạn chế đối với việc phát triển các loại hình dịch vụ hiện đại trên Internet. Phiên bản địa chỉ Internet mới IPv6 được thiết kế để thay thế cho phiên bản IPv4, với hai mục đích cơ bản:
* Thay thế cho nguồn IPv4 cạn kiệt để tiếp nối hoạt động Internet.   
* Khắc phục các nhược điểm trong thiết kế của địa chỉ IPv4.

Địa chỉ IPv6 có chiều dài 128 bít, biểu diễn dưới dạng các cụm số hexa phân cách bởi dấu::, ví dụ 2001:0DC8:1005:2F43:0BCD:FFFF. Với 128 bít chiều dài, không gian địa chỉ IPv6 gồm 2^128 địa chỉ, cung cấp một lượng địa chỉ khổng lồ cho hoạt động Internet.

***Các loại địa chỉ IPv6***
Không gian địa chỉ IPv6 phân thành nhiều loại địa chỉ khác nhau. Mỗi loại địa chỉ có chức năng nhất định trong phục vụ giao tiếp. Khác với phiên bản IPv4, nơi mà một máy tính với một card mạng chỉ được gắn một địa chỉ IPv4 và xác định trên mạng Internet bằng địa chỉ này, một máy tính IPv6 với một card mạng có thể có nhiều địa chỉ, cùng loại hoặc khác loại. Địa chỉ IPv6 không còn duy trì khái niệm broadcast. Theo cách thức gói tin được gửi đến đích, IPv6 bao gồm ba loại địa chỉ sau:
* **Unicast**: Địa chỉ unicast xác định một giao diện duy nhất.
* **Multicast**: Địa chỉ multicast định danh một nhóm nhiều giao diện. Gói tin có địa chỉ đích là địa chỉ multicast sẽ được gửi tới tất cả các giao diện trong nhóm được gắn địa chỉ đó. Mọi chức năng của địa chỉ broadcast trong IPv4 được thay thế bởi địa chỉ IPv6 multicast.
* **Anycast**: Anycast là khái niệm mới của địa chỉ IPv6. Địa chỉ anycast cũng xác định tập hợp nhiều giao diện. Tuy nhiên, trong mô hình định tuyến, gói tin có địa chỉ đích anycast chỉ được gửi tới một giao diện duy nhất trong tập hợp. Giao diện đó là giao diện “gần nhất” theo khái niệm của thủ tục định tuyến..

## 2.2 Domain
Về cơ bản mà nói, muốn truy cập vào tài nguyên của một máy chủ công khai (VD vào một website), ta cần phải biết được địa chỉ IP công khai của nó. Nhưng với địa chỉ chỉ toàn là sổ thì thật khó để nhớ cũng như mường tượng ra ta vào máy chủ nào, cần tài nguyên gì. Vì thế tên miền được giúp ta tưởng tượng được ra tài nguyên cũng như máy chủ mà ta muốn truy cập.

Theo định nghĩa (RFC 1034, được cập nhật bằng RFC 1123), tên miền được tạo thành từ các nhãn không rỗng phân cách nhau bằng dấu chấm (.); những nhãn này giới hạn ở các chữ cái ASCII từ a đến z (không phân biệt hoa thường), chữ số từ 0 đến 9, và dấu gạch ngang (-), kèm theo những giới hạn về chiều dài tên và vị trí dấu gạch ngang.Đó là dấu gạch ngang không được xuất hiện ở đầu hoặc cuối của nhãn, và chiều dài của nhãn nên trong khoảng từ 1 đến 63 và tổng chiều dài của một tên miền không được vượt quá 255 (đây là hạn chế của DNS, xem RFC 2181, tiết đoạn 11).

Vì định nghĩa này không cho phép sử dụng nhiều ký tự thường thấy trong các ngôn ngữ không phải tiếng Anh, và không có các ký tự nhiều byte trong đa số ngôn ngữ châu Á, hệ thống Tên miền quốc tế hóa (IDN) đã được phát triển và hiện đang ở giai đoạn thử nghiệm với một tập tên miền cấp cao nhất được tạo ra vì mục đích này.

Công ty quản lý tên và số hiệu cấp phát Internet (Internet Corporation for Assigned Names and Numbers - ICANN) chịu trách nhiệm chung trong việc quản lý DNS. Nó có nhiệm vụ quản trị tên miền gốc, giao quyền điều hành mỗi tên miền cấp cao nhất cho một cơ quan đăng ký tên miền. Đối với tên miền quốc gia cấp cao nhất, cơ quan đăng ký tên miền thường do chính quyền của quốc gia đó thành lập. ICANN giữ vai trò cố vấn trong các cơ quan đó nhưng không được can thiệp vào các điều khoản và điều kiện về việc ủy quyền tên miền của mỗi cơ quan đăng ký tên miền cấp quốc gia. Tuy nhiên, tên miền cấp cao nhất dùng chung lại do ICANN quản lý trực tiếp, điều đó có nghĩa là tất cả các điều khoản và điều kiện sử dụng sẽ do ICANN quy định cùng với các cơ quan đăng ký tên miền đó.
## 2.3 DNS
Hệ thống phân giải tên miền (DNS - Domain Name System) về căn bản là một hệ thống giúp cho việc chuyển đổi các tên miền mà con người dễ ghi nhớ (dạng ký tự, ví dụ www.example.com) sang địa chỉ IP vật lý (dạng số, ví dụ 123.11.5.19) tương ứng của tên miền đó. DNS giúp liên kết với các trang thiết bị mạng cho các mục đích định vị và địa chỉ hóa các thiết bị trên Internet.
#### Nguyên tắc làm việc của DNS
* Mỗi nhà cung cấp dịch vụ vận hành và duy trì DNS server riêng của mình, gồm các máy bên trong phần riêng của mỗi nhà cung cấp dịch vụ đó trong Internet. Tức là, nếu một trình duyệt tìm kiếm địa chỉ của một website thì DNS server phân giải tên website này phải là DNS server của chính tổ chức quản lý website đó chứ không phải là của một tổ chức (nhà cung cấp dịch vụ) nào khác.
* INTERNIC (Internet Network Information Center) chịu trách nhiệm theo dõi các tên miền và các DNS server tương ứng. INTERNIC là một tổ chức được thành lập bởi NSF (National Science Foundation), AT&T và Network Solution, chịu trách nhiệm đăng ký các tên miền của Internet. INTERNIC chỉ có nhiệm vụ quản lý tất cả các DNS server trên Internet chứ không có nhiệm vụ phân giải tên cho từng địa chỉ.
* DNS có khả năng truy vấn các DNS server khác để có được một cái tên đã được phân giải. DNS server của mỗi tên miền thường có hai việc khác biệt. Thứ nhất, chịu trách nhiệm phân giải tên từ các máy bên trong miền về các địa chỉ Internet, cả bên trong lẫn bên ngoài miền nó quản lý. Thứ hai, chúng trả lời các DNS server bên ngoài đang cố gắng phân giải những cái tên bên trong miền nó quản lý.
* DNS server có khả năng ghi nhớ lại những tên vừa phân giải. Để dùng cho những yêu cầu phân giải lần sau. Số lượng những tên phân giải được lưu lại tùy thuộc vào quy mô của từng DNS.

#### Hoạt động
> ![Hoạt động của DNS](https://upload.wikimedia.org/wikipedia/commons/thumb/a/a5/Example_of_an_iterative_DNS_resolver.svg/400px-Example_of_an_iterative_DNS_resolver.svg.png)\
Một DNS recursor tham khảo 3 máy chủ khác để phân giải tên miền www.wikipedia.org sang dạng IP.

Trong ví dụ ở hình trên ta có thể mô tả sơ qua hoạt động của DNS như sau:\
Giả sử ta muốn tìm địa chỉ IP của www.wikipedia.org, ta biết được rằng tên miền này có các nhãn là www, wikipedia, org.
Trong đó: 
* **www** là tên host của World Wide Web
* **org** là tên miền cấp cao nhất, 
* **wikipedia** là tên miền cấp 2 . 
  
Khi máy trạm muốn truy cập đến một tên miền, nó sẽ gửi yêu cầu đến DNS interator, nếu trong CSDL hiện tại của DNS interator có dữ liệu về thông tin của www.wikipedia.org, nó sẽ trả về IP cho máy trạm. Nếu trong CSDL của DNS interator không có thông tin, nó sẽ truy vấn đến các namesever khác theo thứ tự như sau: 
1. Gửi thông tin đến nameserver cấp root để  truy vấn thông tin về tên miền cấp cao nhất (ở đây là nhãn org), sau đó ROOT nameserver trả về thông tin của các máy chủ quản lý thông tin về tên miền .org
2. DNS iterator lại gửi yêu cầu về máy chủ quản lý tên miền .org để tìm thông tin về tên miền wikipedia.org, ở đây vì tên miền chỉ có 2 mức nên nameserver quản lý tên miền org sẽ trả về IP của tên miền wikipedia.org. Sau đó DNS iterator sẽ lưu tên miền mới này vào trong cache của nó để tiện cho những lần truy vấn sau từ máy trạm.
3. Trong VD này chỉ có 2 mức tên miền nên chỉ cần 2 bước là hoàn thành, trong nhiều trường hợp với tên miền có nhiều mức hơn thì mọi thứ vẫn diễn ra tương tự 
   
## 2.4 TCP/UDP/HTTP/HTTPS/...
### 2.4.1 TCP
TCP (Transmission Control Protocol - "Giao thức điều khiển truyền vận") là một giao thức hướng kết nối nhằm truyền các gói tin đi một cách tin cậy và đúng thứ tự nằm ở tầng giao vận trong mô hình OSI. \
Tầng giao vận là nơi mà các dịch vụ của các ứng dụng trao đổi thông tin với nhau chứ không đơn thuần là các máy tính nữa, nên ở đây địa chỉ của các gói tin được gọi là các **socket**. **Socket** của tầng giao vận bao gồm 2 trường là (IP:port), trong đó IP là địa chỉ IP của máy trạm và port là mã số dịch vụ của các ứng dụng tầng giao vận cần truyền dữ liệu đi.\
Tính tin cậy của giao thức TCP được đảm bảo bằng các tính chất sau:
* Sử dụng gói tin ACK để báo nhận thành công, nếu quá RTT mà không nhận được ACK thì phải phát lại.
* Sử dụng sequence number để dánh số thứ tự cũng như đảm bảo không có gói tin nào bị thất lạc hay trùng lặp
* Sử dụng check sum để kiểm tra sai số của các byte
#### Hoạt động của giao thức
>![Handshankes](https://upload.wikimedia.org/wikipedia/commons/0/08/TCP_state_diagram.jpg)\
Sơ đồ trạng thái của TCP - phiên bản đơn giản hóa

Mô tả các trạng thái của một socket trong hình
1. LISTEN: đang đợi yêu cầu kết nối từ một TCP và cổng bất kỳ ở xa (trạng thái này thường do các TCP server đặt)
2. SYN-SENT: đang đợi TCP ở xa gửi một gói tin TCP với các cờ SYN và ACK được bật (trạng thái này thường do các TCP client đặt)
3. SYN-RECEIVED: đang đợi TCP ở xa gửi lại một tin báo nhận sau khi đã gửi cho TCP ở xa đó một tin báo nhận kết nối (connection acknowledgment) (thường do TCP server đặt)
4. ESTABLISHED: cổng đã sẵn sàng nhận/gửi dữ liệu với TCP ở xa (đặt bởi TCP client và server)
5. FIN-WAIT
6. CLOSE-WAIT
7. CLOSING
8. LAST-ACK
9. TIME-WAIT: đang đợi qua đủ thời gian để chắc chắn là TCP ở xa đã nhận được tin báo nhận về yêu cầu kết thúc kết nối của nó. Theo RFC 793, một kết nối có thể ở tại trạng thái TIME-WAIT trong vòng tối đa 4 phút.
10. CLOSER

**Thiết lập kết nối**\
Để thiết lập một kết nối, TCP sử dụng một quy trình bắt tay 3 bước (3-way handshake) Trước khi client thử kết nối với một server, server phải đăng ký một cổng và mở cổng đó cho các kết nối: đây được gọi là mở bị động. Một khi mở bị động đã được thiết lập thì một client có thể bắt đầu mở chủ động. Để thiết lập một kết nối, quy trình bắt tay 3 bước xảy ra như sau:
1. Client yêu cầu mở cổng dịch vụ bằng cách gửi gói tin SYN (gói tin TCP) tới server, trong gói tin này, tham số sequence number được gán cho một giá trị ngẫu nhiên X.
2. Server hồi đáp bằng cách gửi lại phía client bản tin SYN-ACK, trong gói tin này, tham số acknowledgment number được gán giá trị bằng X + 1, tham số sequence number được gán ngẫu nhiên một giá trị Y
3. Để hoàn tất quá trình bắt tay ba bước, client tiếp tục gửi tới server bản tin ACK, trong bản tin này, tham số sequence number được gán cho giá trị bằng X + 1 còn tham số acknowledgment number được gán giá trị bằng Y + 1

 Tại thời điểm này, cả client và server đều được xác nhận rằng, một kết nối đã được thiết lập.

 **Truyền dữ liệu**\
 Ở hai bước đầu tiên trong ba bước bắt tay, hai máy tính trao đổi một số thứ tự gói ban đầu (Initial Sequence Number -ISN). Số này có thể chọn một cách ngẫu nhiên. Số thứ tự này được dùng để đánh dấu các khối dữ liệu gửi từ mỗi máy tính. Sau mỗi byte được truyền đi, số này lại được tăng lên. Nhờ vậy ta có thể sắp xếp lại chúng khi tới máy tính kia bất kể các gói tới nơi theo thứ tự thế nào.\
 Trên lý thuyết, mỗi byte gửi đi đều có một số thứ tự và khi nhận được thì máy tính nhận gửi lại tin báo nhận (ACK). Trong thực tế thì chỉ có byte dữ liệu đầu tiên được gán số thứ tự trong trường số thứ tự của gói tin và bên nhận sẽ gửi tin báo nhận bằng cách gửi số thứ tự của byte đang chờ.\
Ví dụ: Máy tính A gửi 4 byte với ISN là 100 (theo lý thuyết thì 4 byte sẽ có thứ tự là 100, 101, 102, 103) thì bên nhận sẽ gửi tin báo nhận có nội dung là 104 vì đó là thứ tự của byte tiếp theo nó cần. Bằng cách gửi tin báo nhận là 104, bên nhận đã ngầm thông báo rằng nó đã nhận được các byte 100, 101, 102 và 103. Trong trường hợp 2 byte cuối bị lỗi thì bên nhận sẽ gửi tin báo nhận với nội dung là 102 vì 2 byte 100 và 101 đã được nhận thành công.

**Kết thúc truyền**\
Để kết thúc kết nối hai bên sử dụng quá trình bắt tay 4 bước và chiều của kết nối kết thúc độc lập với nhau. Khi một bên muốn kết thúc, nó gửi đi một gói tin FIN và bên kia gửi lại tin báo nhận ACK. Vì vậy, một quá trình kết thúc tiêu biểu sẽ có 2 cặp gói tin trao đổi.\
Một kết nối có thể tồn tại ở dạng "nửa mở": một bên đã kết thúc gửi dữ liệu nên chỉ nhận thông tin, bên kia vẫn tiếp tục gửi.

#### Gói tin TCP
![Gói tin TCP](https://wiki.mikrotik.com/images/a/ac/Image1006.png)\
**Source port**\
Số hiệu của cổng tại máy tính gửi.\
**Destination port**\
Số hiệu của cổng tại máy tính nhận.\
**Sequence number**\
Trường này có 2 nhiệm vụ. Nếu cờ SYN bật thì nó là số thứ tự gói ban đầu và byte đầu tiên được gửi có số thứ tự này cộng thêm 1. Nếu không có cờ SYN thì đây là số thứ tự của byte đầu tiên.\
**Acknowledgement number**\
Nếu cờ ACK bật thì giá trị của trường chính là số thứ tự gói tin tiếp theo mà bên nhận cần.\
**Data offset**\
Trường có độ dài 4 bít quy định độ dài của phần header (tính theo đơn vị từ 32 bít). Phần header có độ dài tối thiểu là 5 từ (160 bit) và tối đa là 15 từ (480 bít).\
**Reserved**\
Dành cho tương lai và có giá trị là 0.\
**Flags (hay Control bits)**\
Bao gồm 6 cờ:
>**URG**\
Cờ cho trường Urgent pointer\
**ACK**\
Cờ cho trường Acknowledgement\
**PSH**\
Hàm Push\
**RST**\
Thiết lập lại đường truyền\
**SYN**\
Đồng bộ lại số thứ tự\
**FIN**\
Không gửi thêm số liệu\

**Window**\
Số byte có thể nhận bắt đầu từ giá trị của trường báo nhận (ACK)\
**Checksum**\
16 bít kiểm tra cho cả phần header và dữ liệu. Phương pháp sử dụng được mô tả trong RFC 793\
**Urgent pointer**\
Nếu cờ URG bật thì giá trị trường này chính là số từ 16 bít mà số thứ tự gói tin (sequence number) cần dịch trái.\
**Options**\
Đây là trường tùy chọn. Nếu có thì độ dài là bội số của 32 bít.\

### 2.4.2 UDP
Là giao thức hướng tin nhắn của tầng giao vận (message-oriented), khác với TCP, UDP không cung cấp sự tin cậy và thứ tự truyền nhận mà TCP làm; các gói dữ liệu có thể đến không đúng thứ tự hoặc bị mất mà không có thông báo. Tuy nhiên UDP nhanh và hiệu quả hơn đối với các mục tiêu như kích thước nhỏ và yêu cầu khắt khe về thời gian. Do bản chất không trạng thái của nó nên nó hữu dụng đối với việc trả lời các truy vấn nhỏ với số lượng lớn người yêu cầu.
#### Cổng
UDP dùng cổng để cho phép các giao tiếp giữa các ứng dụng diễn ra.\
Cổng dùng 16 bit để đánh địa chỉ, vì vậy số của cổng nằm trong khoản 0 đến 65.535. Cổng 0 được để dành và không nên sử dụng.\
Cổng từ 1 đến 1023 được gọi là cổng "well-known" và trên các hệ điều hành tựa Unix, việc gắn kết tới một trong những cổng này đòi hỏi quyền root.\
Cổng 1024 đến 49.151 là cổng đã đăng ký.\
Cổng từ 49.152 đến 65.535 là các cổng tạm, được dùng chủ yếu bởi client khi liên lạc với server.
#### Cấu trúc gói tin
UDP là giao thức hướng thông điệp nhỏ nhất của tầng giao vận hiện được mô tả trong RFC 768 của IETF.\
Trong bộ giao thức TCP/IP, UDP cung cấp một giao diện rất đơn giản giữa tầng mạng bên dưới (thí dụ, IPv4) và tầng phiên làm việc hoặc tầng ứng dụng phía trên.\
UDP không đảm bảo cho các tầng phía trên thông điệp đã được gửi đi và người gửi cũng không có trạng thái thông điệp UDP một khi đã được gửi (Vì lý do này đôi khi UDP còn được gọi là Unreliable Datagram Protocol).\
UDP chỉ thêm các thông tin multiplexing và giao dịch. Các loại thông tin tin cậy cho việc truyền dữ liệu nếu cần phải được xây dựng ở các tầng cao hơn.

| +   | Bits 0 - 15 |     16 - 31      |
| --- | :---------: | :--------------: |
| 0   | Source Port | Destination Port |
| 32  |   Length    |     Checksum     |
| 64  |    Data     |                  |

Phần header của UDP chỉ chứa 4 trường dữ liệu, trong đó có 2 trường là tùy chọn (ô nền đỏ trong bảng).\
**Source port**\
Trường này xác định cổng của người gửi thông tin và có ý nghĩa nếu muốn nhận thông tin phản hồi từ người nhận. Nếu không dùng đến thì đặt nó bằng 0.\
**Destination port**\
Trường xác định cổng nhận thông tin, và trường này là cần thiết.\
**Length**\
Trường có độ dài 16 bit xác định chiều dài của toàn bộ datagram: phần header và dữ liệu. Chiều dài tối thiểu là 8 byte khi gói tin không có dữ liệu, chỉ có header.\
**Checksum**\
Trường checksum 16 bit dùng cho việc kiểm tra lỗi của phần header và dữ liệu. Phương pháp tính checksum được định nghĩa trong RFC 768.\
Do thiếu tính tin cậy, các ứng dụng UDP nói chung phải chấp nhận mất mát, lỗi hoặc trùng dữ liệu. Một số ứng dụng như TFTP có nhu cầu phải thêm những kỹ thuật làm tin cậy cơ bản vào tầng ứng dụng. Hầu hết các ứng dụng UDP không cần những kỹ thuật làm tin cậy này và đôi khi nó bị bỏ đi. Streaming media, game trực tuyến và voice over IP (VoIP) là những thí dụ cho các ứng dụng thường dùng UDP. Nếu một ứng dụng đòi hỏi mức độ cao hơn về tính tin cậy, những giao thức như TCP hoặc mã erasure có thể được sử dụng để thay thế.\
Thiếu những cơ chế kiểm soát tắc nghẽn và kiểm soát luồng, các kỹ thuật dựa trên mạng là cần thiết để giảm nguy hiệu ứng cơ tắc nghẽn dây chuyền do không kiểm soát, tỷ lệ tải UDP cao. Nói cách khác, vì người gởi gói UDP không thể phát hiện tắc nghẽn, các thành phần dựa trên mạng như router dùng hàng đợi gói (packet queueing) hoặc kỹ thuật bỏ gói như là những công cụ để giảm tải của UDP. Giao thức Datagram Congestion Control Protocol (DCCP) được thiết kế như một giải pháp cho vấn đề bằng cách thêm hành vi kiểm soát tắc nghẽn cho thiết bị đầu cuối cho các dòng dữ liệu UDP như streaming media.
### 2.4.3 HTTP
HTTP (Tiếng Anh: HyperText Transfer Protocol - Giao thức truyền tải siêu văn bản) là một trong năm giao thức chuẩn của mạng Internet, được dùng để liên hệ thông tin giữa Máy cung cấp dịch vụ (Web server) và Máy sử dụng dịch vụ (Web client) trong mô hình Client/Server dùng cho World Wide Web-WWW, HTTP là một giao thức thuộc tầng ứng dụng, nằm trên cặp giao thức tầng giao vận & tầng mạng là TCP/IP.
#### HTTP Request methods
HTTP Request Method chỉ phương thức để được thực hiện trên nguồn được nhận diện bởi Request-URI đã cung cấp\
>**GET**\
GET được sử dụng để lấy lại thông tin từ Server đã cung cấp bởi sử dụng một URI đã cung cấp. Các yêu cầu sử dụng GET chỉ nhận dữ liệu và không có ảnh hưởng gì tới dữ liệu.\
**HEAD**\
Tương tự như GET, nhưng nó truyền tải dòng trạng thái và khu vực Header.\
**POST**\
Một yêu cầu POST được sử dụng để gửi dữ liệu tới Server, ví dụ, thông tin khách hàng, file tải lên, …, bởi sử dụng các mẫu HTML.\
**PUT**\
Thay đổi tất cả các đại diện hiện tại của nguồn mục tiêu với nội dung được tải lên.\
**DELETE**\
Gỡ bỏ tất cả các đại diện hiện tại của nguồn mục tiêu bởi URI.\
**CONNECT**\
Thiết lập một tunnel tới Server được xác định bởi URI đã cung cấp.\
**OPTIONS**\
Miêu tả các chức năng giao tiếp cho nguồn mục tiêu.\
**TRACE**\
Trình bày một vòng lặp kiểm tra thông báo song song với path tới nguồn mục tiêu.\
#### HTTP Response
Khi nhận và phiên dịch một HTTP Request, Server sẽ gửi tín hiệu phản hồi là một HTTP Response bao gồm các thành phần sau:
* Một dòng trạng thái (Status-Line)
* Không hoặc nhiều hơn các trường Header (General|Response|Entity) được theo sau CRLF
* Một dòng trống chỉ dòng kết thúc của các trường Header
* Một phần thân thông báo tùy ý
### 2.4.4 HTTPS
HTTPS, viết tắt của Hypertext Transfer Protocol Secure, là một giao thức kết hợp giữa giao thức HTTP và giao thức bảo mật SSL hay TLS cho phép trao đổi thông tin một cách bảo mật trên Internet. Giao thức HTTPS thường được dùng trong các giao dịch nhạy cảm cần tính bảo mật cao.\
Giao thức HTTPS sử dụng port 443, giúp đảm bảo các tính chất sau của thông tin:
* Confidentiality: sử dụng phương thức mã hóa (encryption) để đảm bảo rằng các thông điệp được trao đổi giữa client và server không bị kẻ thứ ba đọc được.
* Integrity: sử dụng phương thức hashing để cả người dùng (client) và máy chủ (server) đều có thể tin tưởng rằng thông điệp mà chúng nhận được có không bị mất mát hay chỉnh sửa.
* Authenticity: sử dụng chứng chỉ số (digital certificate) để giúp client có thể tin tưởng rằng server/website mà họ đang truy cập thực sự là server/website mà họ mong muốn vào, chứ không phải bị giả mạo.


Việc nhờ đến bên thứ 3 (thường là CA) để xác thực danh tính của website cộng thêm sự chú ý của người dùng rằng website đó có sử dụng HTTPS và SSL certificate của nó còn hiệu lực sẽ giúp loại bỏ hoàn toàn nguy cơ bị lừa đảo.
#### Quá trình giao tiếp giữa client và server thông qua HTTPS
1. Client gửi yêu cầu (request) cho một trang bảo mật (secure page) (có URL bắt đầu với https://)
2. Server gửi lại cho client certificate của nó.
3. Client (trình duyệt web) tiến hành xác thực certificate này bằng cách kiểm tra (verify) tính hợp lệ của chữ ký số của CA được kèm theo certificate.\
Giả sử certificate đã được xác thực và còn hạn sử dụng hoặc client vẫn cố tình truy cập mặc dù Web browser đã cảnh báo rằng không thể tin cậy được certificate này (do là dạng self-signed SSL certificate hoặc certificate hết hiệu lực, thông tin trong certificate không đúng) thì mới xảy ra bước 4 sau.
4. Client tự tạo ra ngẫu nhiên một symmetric encryption key (hay session key), rồi sử dụng public key (lấy trong certificate) để mã hóa session key này và gửi về cho server.
5. Server sử dụng private key (tương ứng với public key trong certificate ở trên) để giải mã ra session key ở trên.
6. Sau đó, cả server và client đều sử dụng session key đó để mã hóa/giải mã các thông điệp trong suốt phiên truyền thông.

## 3 Một số vấn đề cơ bản về linux
## 3.1 Cấu trúc cây thư mục trong Linux
Linux quan niệm bất cứ thứ j tồn tại trong máy tính của bạn đều là FILE (tệp tin) cả. Ngay cả DIRECTORY (thư mục) cũng là 1 loại tệp tin. Do đó mà lệnh "file" luôn luôn tương thích với mọi thứ trong hệ thống. Nó sẽ cho ra kết quả là filetype (định dạng tệp tin) của file mà ta kiểm tra. Sơ lược thì hệ thống thư mục của ubuntu sẽ như sau:

![Cấu trúc cây thư mục](https://images.viblo.asia/00898935-d72c-4019-b584-4cfc6af8593d.png)
### 3.1.1. / – Root
Đúng với tên gọi của mình: nút gốc (root) đây là nơi bắt đầu của tất cả các file và thư mục. Chỉ có root user mới có quyền ghi trong thư mục này. Chú ý rằng /root là thư mục home của root user chứ không phải là /.

### 3.1.2. /bin – Chương trình của người dùng
Thư mục này chứa các chương trình thực thi. Các chương trình chung của Linux được sử dụng bởi tất cả người dùng được lưu ở đây. Ví dụ như: ps, ls, ping…

### 3.1.3. /sbin – Chương trình hệ thống
Cũng giống như /bin, /sbinn cũng chứa các chương trình thực thi, nhưng chúng là những chương trình của admin, dành cho việc bảo trì hệ thống. Ví dụ như: reboot, fdisk, iptables…

### 3.1.4. /etc – Các file cấu hình
Thư mục này chứa các file cấu hình của các chương trình, đồng thời nó còn chứa các shell script dùng để khởi động hoặc tắt các chương trình khác. Ví dụ: /etc/resolv.conf, /etc/logrolate.conf

### 3.1.5. /dev – Các file thiết bị
Các phân vùng ổ cứng, thiết bị ngoại vi như USB, ổ đĩa cắm ngoài, hay bất cứ thiết bị nào gắn kèm vào hệ thống đều được lưu ở đây. Ví dụ: /dev/sdb1 là tên của USB bạn vừa cắm vào máy, để mở được USB này bạn cần sử dụng lệnh mount với quyền root: # mount /dev/sdb1 /tmp

### 3.1.6. /tmp – Các file tạm
Thư mục này chứa các file tạm thời được tạo bởi hệ thống và các người dùng. Các file lưu trong thư mục này sẽ bị xóa khi hệ thống khởi động lại.

### 3.1.7. /proc – Thông tin về các tiến trình
Thông tin về các tiến trình đang chạy sẽ được lưu trong /proc dưới dạng một hệ thống file thư mục mô phỏng. Ví dụ thư mục con /proc/{pid} chứa các thông tin về tiến trình có ID là pid (pid ~ process ID). Ngoài ra đây cũng là nơi lưu thông tin về về các tài nguyên đang sử dụng của hệ thống như: /proc/version, /proc/uptime…

### 3.1.8. /var – File về biến của chương trình
Thông tin về các biến của hệ thống được lưu trong thư mục này. Như thông tin về log file: /var/log, các gói và cơ sở dữ liệu /var/lib…

### 3.1.9. /usr – Chương trình của người dùng
Chứa các thư viện, file thực thi, tài liệu hướng dẫn và mã nguồn cho chương trình chạy ở level 2 của hệ thống. Trong đó:

| Thư mục thứ cấp  | Description       |
| ---------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| /usr/bin         | Lưu nhiều file thi hành của hệ thống.                                                                                                                                                      |
| /usr/etc         | Lưu nhiều file cấu hình hệ thống                                                                                                                                                           |
| /usr/include     | Tại đây và trong nhiều thư mục cấp dưới của /usr/include là nơi lưu tất cả các file kèm theo bộ biên dịch C. Những file header này định nghĩa các hằng và hàm dùng trong lập trình bằng C. |
| /usr/g++-include | Lưu các file kèm theo bộ biên dịch C.                                                                                                                                                      |
| /usr/lib         | Chứa các thư viện để chương trình sử dụng trong khi kết nối                                                                                                                                |
| /usr/share/man   | Chứa các trang thủ công cho chương trình. Bên dưới /usr/share/man là nhiều thư mục tương ứng với các đoạn trong trang man.                                                                 |
| /dev/pty*        | Driver hỗ trợ terminal giả, dùng cho việc đăng nhập từ xa, chẳng hạn như những phiên đăng nhập qua Telnet.                                                                                 |
| /usr/src         | Chứa các thư mục mã nguồn của nhiều chương trình trên hệ thống. Nếu nhận được gói phần mềm chờ cài đặt, bạn nên lưu vào /usr/src/tên-gói trước khi cài đặt.                                |
| /usr/local       | Dành riêng cho việc thiết kế hoặc tùy chỉnh các ứng dụng cho phù hợp với hệ thống máy bạn. Nhìn chung, hầu hết phần mềm dùng tại chỗ được lưu trong các thư mục cấp dưới của thư mục này.  |

### 3.1.10. /home – Thư mục người của dùng
Thư mục này chứa tất cả các file cá nhân của từng người dùng. Ví dụ: /home/john, /home/marie

### 3.1.11. /boot – Các file khởi động
Tất cả các file yêu cầu khi khởi động như initrd, vmlinux. grub được lưu tại đây. Ví dụ vmlixuz-2.6.32-24-generic

### 3.1.12. /lib – Thư viện hệ thống
Chứa cá thư viện hỗ trợ cho các file thực thi trong /bin và /sbin. Các thư viện này thường có tên bắt đầu bằng ld* hoặc lib*.so.*. Ví dụ như ld-2.11.1.so hay libncurses.so.5.7

### 3.1.13. /opt – Các ứng dụng phụ tùy chọn
Tên thư mục này nghĩa là optional (tùy chọn), nó chứa các ứng dụng thêm vào từ các nhà cung cấp độc lập khác. Các ứng dụng này có thể được cài ở /opt hoặc một thư mục con của /opt

### 3.1.14. /mnt – Thư mục để mount
Đây là thư mục tạm để mount các file hệ thống. Ví dụ như # mount /dev/sda2 /mnt

### 3.1.15. /media – Các thiết bị gắn có thể gỡ bỏ
Thư mục tạm này chứa các thiết bị như CdRom /media/cdrom. floppy /media/floopy hay các phân vùng đĩa cứng /media/Data (hiểu như là ổ D:/Data trong Windows)

### 3.1.16. /srv – Dữ liệu của các dịch vụ khác
Chứa dữ liệu liên quan đến các dịch vụ máy chủ như /srv/svs, chứa các dữ liệu liên quan đến CVS.

## 3.2 Process, Thread,...
### 3.2.1 Process
#### Khái quát về tiến trình
Khác với chương trình, tiến trình (process) là một phiên bản đang chạy của một chương trình. Ví dụ một chương trình helloworld.exe đang chạy trong máy tính được gọi là một tiến trình. Tiến trình được gọi là thực thể chủ động (active entity) để phân biệt với chương trình, vì nó sẽ được copy từ ổ cứng sang bộ nhớ RAM để chạy trên hệ điều hành máy tính. Máy tính xây dựng một tiến trình từ một chương trình, bằng cách copy chương trình đang nằm trên ổ cứng vào RAM, cấp phát các tài nguyên cần thiết để xây dựng một tiến trình. Bạn có thể tạo được nhiều tiến trình từ một chương trình, ví dụ như khi bạn double click để chạy chương trình helloworld.exe nhiều lần.

Trên góc nhìn từ kernel, một tiến trình bao gồm (1) không gian bộ nhớ người dùng (user-space) chứa mã nguồn chương trình và (2) các cấu trúc dữ liệu của kernel chứa thông tin trạng thái của tiến trình đó. Kernel duy trì một cấu trúc task_struct cho mỗi tiến trình để lưu các thông tin cho mỗi tiến trình đó như các process ID liên quan, virtual memory table, bảng mô tả file (open file descriptor), đường dẫn hiện tại (current working directory)...

#### Định danh tiến trình (process ID)
Mỗi tiến trình có một số định danh (process ID, hay pid), là một số dương để xác định tiến trình đó là duy nhất trong hệ thống. Lập trình viên có thể tác động lên tiến trình (ví dụ kill tiến trình) bằng một số system call với đối số truyền vào là process ID đó. Có thể lấy process ID của tiến trình đang chạy bằng system call getpid() với prototype như sau:

Có thể dùng câu lệnh “ps aux” để liệt kê các tiến trình đang chạy trên hệ thống, ví dụ như sau:
![PID](https://lh5.googleusercontent.com/T3kSJh_f6UGZa84lUgzMdtu3lyZt6zOps2GPPBR7AJOqixfKx_8EjKM5irJ0EjJJRyLkBM9aW332NSEWBHf38-QsWl550LbnRfV-S9_v7BBfEMaRXWs4KwwPdaVZD6d9o6kX3rqK)

Cũng trong hình vẽ trên ta có thể thấy, trong Linux tiến trình init (/sbin/init) là tiến trình được khởi tạo đầu tiên bởi hệ thống, vì vậy luôn có pid là 1. Tiến trình init sau khi được khởi tạo bởi kernel, sẽ làm nốt các phần việc còn lại là khởi tạo các tiến trình cần thiết khác đề hoàn tất quá trình booting của hệ thống.

Mỗi tiến trình sẽ có một tiến trình cha (parent process), là tiến trình tạo ra nó. Khi một tiến trình trở thành orphan (tạm dịch là tiến trình mồ côi) do tiến trình cha của nó bị kết thúc, nó sẽ trở thành tiến trình con của tiến trình đầu tiên “init”.

#### Cấu trúc bộ nhớ của tiến trình
Về cơ bản thì một tiến trình sẽ có một không gian bộ nhớ ảo (virtual memory) 4 GB (trong kiến trúc x86-32). Hình vẽ dưới đây phác thảo một cách đơn giản về các memory segment trong không gian bộ nhớ của một tiến trình:

![Cấu trúc bộ nhớ của một tiến trình](https://lh4.googleusercontent.com/QhcA0kKuZmvnwFqgIEvAJazaBX4sm8W9IcaCnzGSYzXEzSA7TgCJEoBcc-cXimCiXmslGekglljh10QwOvn2-o7lPY1vh9yndiOfr0UA-dtF4ncXbj01teG7M8tT6icu8Cr4BDRZ)

Bộ nhớ được cấp phát cho một tiến trình bao gồm nhiều thành phần được gọi là các segment. Cụ thể các segment như sau:
- **Text segment:** Chứa các chỉ lệnh ngôn ngữ máy (machine-language instruction) của chương trình mà tiến trình đó chạy. Text segment chính là các chỉ lệnh được biên dịch từ source code của lập trình viên cho chương trình đó. Nội dung của phân vùng này không thay đổi trong suốt quá trình process tồn tại nên được kernel thiết lập chế độ read-only để bảo vệ khỏi sự truy cập vô tình hay cố ý của người dùng. Như đã nói ở trên, vì nhiều tiến trình có thể chạy một chương trình nên text segment cũng được thiết lập sharable để các tiến trình có thể sử dụng chung để tiết kiệm tài nguyên.
- **Initialized data segment:** Vùng này chứa các biến toàn cục (global) và biến static mà đã được khởi tạo từ code của chương trình. Giá trị của các biến này được đọc từ các file thực thi khi chương trình được tải vào RAM. Ví dụ khi lập trình viên khai báo biến tĩnh “static var = 10;”, biến var này sẽ được lưu vào vùng nhớ của initialized data segent.
- **Uninitialized data segment:** Còn được gọi là vùng bss segment. Segment này chứa các biến toàn cục và static mà chưa được khởi tạo từ source code. Ví dụ khi lập trình viên khai báo biến tĩnh “static var;”, biến var sẽ được chứa ở vùng này và được khởi tạo giá trị 0. Nếu bạn thắc mắc tại sao biến var không được lưu vào vùng initialized data segment cho đơn giản. Câu trả lời là khi một chương trình được lưu trữ vào ổ cứng, không cần thiết phải cấp phát tài nguyên cho uninitialized data segment; thay vào đó chương trình chỉ cần nhớ vị trí và kích thước biến được yêu cầu cho vùng này, các biến này sẽ được cấp phát run time khi chương trình được tải vào RAM.
- **Stack segment:** Chứa stack frame của tiến trình. Chúng ta sẽ tìm hiểu sâu hơn về stack và stack frame ở phần dưới đây. Tạm thời hiểu là khi 1 hàm được gọi, một stack frame sẽ được cấp phát cho hàm đó (các biến được khai báo trong hàm, các đối số truyền vào hay giá trị return) và sẽ bị thu hồi khi hàm đó kết thúc. Vì vậy, stack segment có thể giãn ra hoặc co lại khi tiến trình cấp phát/thu hồi các stack frame.
- **Heap segment:** Là vùng bộ nhớ lưu các biến được cấp phát động (dynamic allocate) tại thời điểm run time. Tương tự như stack segment, heap segment cũng có thể giãn ra hoặc co vào khi một biến được cấp phát hoặc free.

#### Stack và stack frame
Stack segment của tiến trình có thể giãn ra hoặc co lại mỗi khi một hàm được gọi hoặc return. Trong kiến trúc x86-32 có một thanh ghi đặc biệt được gọi là stack pointer (sp) lưu thông tin địa chỉ top của stack. Khi một hàm được gọi, một stack frame của hàm đó được cấp phát và push vào stack; và stack frame này sẽ được thu hồi khi hàm đó return.

Mỗi stack frame chứa các thông tin sau:
- Đối số (argument) của hàm và các biến cục bộ (local variable): Các biến này còn được gọi là biến tự động (automatic variable) vì chúng sẽ tự động được tạo ra khi hàm được gọi và tự động biến mất khi hàm đó return (vì stack frame cũng biến mất).
- Call linkage information: Các hàm khi chạy sẽ sử dụng các thanh ghi của CPU, ví dụ như thanh ghi program counter (pc) lưu chỉ lệnh tiếp theo được thực thi. Mỗi lần một hàm (ví dụ hàm X) gọi hàm khác (ví dụ hàm Y), giá trị các thanh ghi mà hàm X đang dùng sẽ được lưu vào stack frame của hàm Y; và sau khi hàm Y return, các giá trị thanh ghi này sẽ được phục hồi cho hàm X tiếp tục chạy.

#### Biến môi trường
Mỗi tiến trình có 1 danh sách các biến ở dạng string gắn với nó được gọi là các biến môi trường (environment list). Mỗi chuỗi trong số này được định nghĩa dưới dạng name=value, các biến này được dùng để lưu trữ 1 thông tin bất kỳ mà tiến trình muốn giữ. Hiểu một cách đơn giản, biến môi trường là biến của tiến trình đang chạy đó, không phải là biến của một hàm nào cả và được lưu trữ trong không gian bộ nhớ của tiến trình đó.

Khi 1 chương trình được tạo ra, nó sẽ kế thừa các biến môi trường của tiến trình cha. Vì đặc điểm này, sử dụng biến môi trường cũng có thể coi là 1 cách rất đơn giản cho giao tiếp liên tiến trình (IPC) giữa tiến trình cha và con. Ví dụ khi bạn tạo 1 biến môi trường từ tiến trình cha rồi tạo ra một tiến trình con, tiến trình con này sẽ lưu giữ giá trị của biến đó. Lưu ý là việc lưu giữ này chỉ là một chiều và một lần, nghĩa là sau đó nếu tiến trình cha hoặc con thay đổi giá trị của biến đó thì nó sẽ không được cập nhật sang cho tiến trình còn lại.

#### Kết luận
Phần trình bày này đã giúp chúng ta hiểu rõ khái niệm về tiến trình, cũng như các thành phần và tổ chức bộ nhớ của tiến trình. Qua đó ta có thể ứng dụng vào lập trình tốt hơn.

### 3.2.2 Thread
#### Tổng quan về thread
Thread là một cơ chế cho phép một ứng dụng thực thi đồng thời nhiều công việc (multi-task). Ví dụ một trường hợp đòi hỏi multi-task sau: một tiến trình web server của một trang web giải trí phải phục vụ hàng trăm hoặc hàng nghìn client cùng một lúc. Công việc đầu tiên của tiến trình là lắng nghe xem có client nào yêu cầu dịch vụ không. Giả sử có client A kết nối yêu cầu nghe một bài nhạc, server phải xử lý chạy bài hát client A yêu cầu; nếu trong lúc đó client B kết nối yêu cầu tải một bức ảnh thì server lúc đó không thể phục vụ vì đang bận phục vụ client A. Đây chính là kịch bản yêu cầu một tiến trình cần thực thi multi-task. Qua các bài học về process, ta thấy tiến trình server nói trên có thể giải quyết bài toán này như sau: Server chỉ làm công việc chính là lắng nghe xem có client nào yêu cầu dịch vụ không; khi tiến trình A kết nối, server dùng system call fork() để tạo ra một tiến trình con chỉ làm công việc client A yêu cầu, trong khi nó lại tiếp tục lắng nghe các yêu cầu từ các client khác. Tương tự, khi client B kết nối, server lại tạo ra một tiến trình con khác phục vụ yêu cầu của client B. Trong bài học này, chúng ta sẽ thấy việc xử lý đa tác vụ của server như trên có thể dùng thread. Thậm chí trong trường hợp này thread còn tỏ ra thích hợp hơn so với việc sử dụng tiến trình con như đã giải thích ở trên, chúng ta sẽ tìm hiểu rõ hơn trong bài này.

Thread là một thành phần của tiến trình, một tiến trình có thể chứa một hoặc nhiều thread. Hệ điều hành Unix quan niệm rằng mỗi tiến trình khi bắt đầu chạy luôn có một thread chính (main thread); nếu không có thread nào được tạo thêm thì tiến trình đó được gọi là đơn luồng (single-thread), ngược lại nếu có thêm thread thì được gọi là đa luồng (multi-thread). Các thread trong tiến trình chia sẻ các vùng nhớ toàn cục (global memory) của tiến trình bao gồm initialized data, uninitialized data và vùng nhớ heap:

Hình vẽ dưới đây mô tả về một tiến trình đơn luồng (single-thread) và đa luồng (multi-thread):
![Tiến trình single-thread và multi-thread](https://vimentor.com/vi/image/2/n/n/editors/single%20thread-multi%20thread.jpg)

Trong hình vẽ trên, một tiến trình có 4 thread, bao gồm 1 main thread (T0) được tạo ra khi tiến trình chạy hàm main(), và 3 thread lần lượt là T1, T2 và T3 được tạo mới trong hàm main(). Bốn thread sử dụng chung vùng nhớ toàn cục (global memory) nhưng mỗi thread có phân vùng stack riêng của mình, cụ thể như hình vẽ dưới đây:

![Tổ chức bộ nhớ của tiến trình có 4 thread](https://vimentor.com/vi/image/2/n/n/editors/thread.jpg)

Các thread trong tiến trình có thể thực thi đồng thời và độc lập với nhau. Nghĩa là nếu một thread bị block do đang chờ I/O thì các thread khác vẫn được lập lịch và thực thi thay vì cả tiến trình bị block.

#### So sánh Process và Thread
Quay trở lại ví dụ trên, tiến trình server tạo ra các tiến trình con để phục vụ yêu cầu multi-task. Cách này tuy giải quyết được yêu cầu nhưng tồn tại các hạn chế sau đây:
- Việc chia sẻ dữ liệu giữa các tiến trình khá khó khăn. Vì mỗi tiến trình trong Linux có không gian bộ nhớ riêng biệt nên chúng ta phải sử dụng các phương pháp giao tiếp liên tiến trình (IPC) như shared memory, message queue, socket... để chia sẻ dữ liệu.
- Tạo ra một tiến trình mới bằng system call fork() khá "tốn kém" về mặt tài nguyên cũng như thời gian vì phải tạo ra các vùng nhớ riêng biệt cho tiến trình con. Điều này khá quan trọng trong các hệ thống embedded có phần cứng bị hạn chế.

Thread có thể giải quyết được 2 vấn đề trên vì có các ưu điểm sau:
- Chia sẻ dữ liệu giữa các thread trong tiến trình rất đơn giản vì chúng sử có chung không gian bộ nhớ toàn cục. Do vậy, chỉ cần tạo dữ liệu ở trong các vùng nhớ toàn cục này thì các thread đều có thể truy xuất được.
- Việc tạo ra một thread mới nhanh hơn đáng kể so với việc tạo ra một tiến trình mới vì các thread dùng chung nhiều phần không gian bộ nhớ nên chỉ cần tạo không gian bộ nhớ cho những phần riêng thay vì phải nhân bản toàn bộ các vùng nhớ như khi tạo tiến trình con.

Hiển nhiên thread cũng không phải là chìa khóa vạn năng. Việc sử dụng thread cũng có các nhược điểm sau:
- Vì các thread dùng chung vùng nhớ toàn cục nên việc lập trình trên các thread "nguy hiểm" hơn trên process. Nếu một thread gây ra lỗi trên vùng nhớ toàn cục thì sẽ kéo theo các thread khác cũng bị lỗi theo.
- Các thread cùng chia sẻ vùng nhớ toàn cục của một tiến trình (3 GB với hệ thống 32 bit), cụ thể mỗi thread sẽ được cung cấp một vùng nhớ riêng trong tổng thể bộ nhớ của tiến trình. Bộ nhớ của tiến trình tuy là lớn nhưng cũng là một số hữu hạn. Nên một tiến trình cũng bị giới hạn bởi số lượng thread có thể tạo ra hoặc tạo ra các thread cần bộ nhớ lớn.

Cả hai nhược điểm trên không xảy ra trên tiến trình vì mỗi tiến trình có không gian bộ nhớ riêng.

## 3.3 Permission
#### Tại sao cần có phân quyền?
Ngày này hầu hết các hệ điều hành phổ biến đều hỗ trợ nhiều người cùng sử dụng đồng thời. Có thể sẽ có người thắc mắc rằng tại sao lại cần hỗ trợ tính năng này khi máy tính của họ chỉ có một người sử dụng. Lý do là vì một hệ điều hành được tạo ra với mục đích càng có nhiều người sử dụng càng tốt. Do vậy có thể một người nào đó không cần hệ điều hành phải hỗ trợ nhiều người sử dụng nhưng những người khác lại cần đến. Tính năng phân quyền được tạo ra là để hỗ trợ cho việc sử dụng đồng thời của hệ điều hành.
#### Hệ thống file trong Linux
Có một câu nói phổ biến là:
>Trong Linux mọi thứ đều là file

Khái niệm file trong Linux khá mềm dẻo, kể cả thư mục hay thiết bị cũng được coi như file đặc biệt. Và vì vậy, hệ thống phân quyền sẽ hoạt động trên cả thư mục, thiết bị ....

#### Nhóm phân quyền
Trong Linux có 3 nhóm phân quyền chính, không bao nhau:

- Owner: chỉ cấp quyền cho chủ sở hữu của file.
- Group: chỉ cấp quyền cho nhóm sở hữu của file.
- Other: cấp quyền cho những người dùng và nhóm không thuộc 2 nhóm trên.

Tuy nhiên vẫn có 1 ngoại lệ, đó là người dùng root (siêu người dùng). Người dùng này có mọi quyền hạn trên mọi file trong hệ thống, không bị ràng buộc bởi bất cứ sự phân quyền nào. Như mình đã nói trong phần mở đầu, hệ thống phân quyền sẽ rất nguy hiểm nếu bạn không hiểu về nó. VD câu lệnh sau sẽ xóa tất cả file trên hệ thống nếu người thực thi là root

>rm -rf /

Do vậy với bất cứ một tác vụ gì phải dùng tới quyền hạn của người dùng root ta nên chắc chắn sẽ không gây hại cho hệ thống thì mới nên sử dụng, hoặc tìm một phương án thay thế mà không cần tới quyền hạn nguy hiểm này.

#### Loại phân quyền
Với một file, có 3 loại phân quyền cơ bản như trong bảng sau:

| Tên quyền | Ký hiệu | Dạng số | Mô tả          |
| --------- | ------- | ------- | -------------- |
| Read      | r       | 4       | Quyền đọc      | file |
| Write     | w       | 2       | Quyền ghi      | file |
| Execute   | e       | 1       | Quyền thực thi | file |

Ngoài ra có một vài phân quyền đặc biệt như sau:

| Tên quyền      | Ký hiệu | Dạng số | Mô tả                                                                                        |
| -------------- | ------- | ------- | -------------------------------------------------------------------------------------------- |
| Setuid/Setguid | s       | 1       | Nếu file được thực thi, người thực thi sẽ là chủ sở hữu                                      |
| Sticky bit     | t       | 1       | Chỉ chủ sở hữu mới được xóa hoặc thay đổi tên file kể cả khi Other có toàn quyền với file đó |

#### Cách xem phân quyền của một file
Trong Linux, cách đơn giản nhất để xem phân quyền của một file là sử dụng lệnh ls -l Output của câu lệnh trên sẽ có dạng như sau: -rwxr-x-r-x 1 user group
- Kí tự - đầu tiên là một cờ đặc biệt để chỉ loại file, - với file thông thường, d với thư mục, c với thiết bị, l với liên kết (liên kết tới một file khác).
- 3 kí tự rwx đầu tiên là quyền hạn của Owner, ở đây Owner sẽ có mọi quyền với file
- 3 kí tự r-x ở giữa là quyền hạn của Group, ở đây Group sẽ có quyền đọc và quyền dùng lệnh cd
- 3 kí tự r-x cuối cùng là quyền hạn của Other, tương tự như Group ở trên sẽ có quyền đọc và dùng lệnh cd
- Số nguyên đi sau quyền hạn để chỉ số lượng liên kết tới file, ở đây 1 có nghĩa là file này không có liên kết tượng trưng mà chỉ có một liên kết cứng duy nhất trỏ tới chính nó.
- Cuối cùng là 2 thông tin nói về chủ sở hữu và nhóm sở hữu, ở đây là người dùng user và nhóm group

## 3.4 Network configuration 
1) Thay đổi địa chỉ IP của máy (chỉ có hiệu lực trong tức thời). Giả sự ta muốn thay đổi IP của eth0 thành 10.30.255.100 với netmask là 255.255.0.0, thực hiện như sau:
>[root@son]# ifconfig eth0 10.30.255.100 netmask 255.255.0.0 up

Lưu ý: Cũng thể sử dung câu lệnh sau để cấu hình lại mạng:
>[root@son]# setup

2) Start và Stop 1 card mạng:
>[root@son]# ifup eth0

>[root@son]# ifdown eth0

3) Thay đổi default gateway:

>[root@son]# route add default gw 10.30.0.1 eth0

>[root@son]# route delete default gw 10.30.0.1 eth0

Nếu muốn save lại default gw cho eth0 thì cần cập nhật nội dung của 1 trong 2 file hoặc cả 2:
>[root@son]# vi /etc/sysconfig/network

Hoặc:
>[root@son]# vi /etc/sysconfig/network-scripts/ifcfg-eth0

Với nội dung sau:
NETWORK=yes
HOSTNAME=localhost.localdomain
GATEWAYE=10.30.0.1

4) Thêm mạng con:
>[root@son]# route add -net 10.30.1.0 netmask 255.255.255.0 eth0

5) Tạo một alias (bí danh) cho eth0:
>[root@son]# ifconfig eth0:0 192.168.1.100

6) Xem thông tin mạng:
>[root@son]# ifconfig -a
>[root@son]# arp -a
>[root@son]# netstat -nr (hoặc -i hoặc -an)
Ví dụ: Tìm ra số lượng các Connection được thiết lập:
[root@son]# netstat -an|grep tcp|egrep -i 'established|timewait' | wc -l

[root@son]# route
[root@son]# ip addr show
[root@son]# ip route show 0/0
[root@son]# ipvsadm -L -n #Nếu đã cài ipvsadm (yum install ipvsadm)

## 3.5 Basic Terminal Command
### 3.5.1 cat 
cat là một trong các lệnh cơ bản trong Linux được sử dụng thường xuyên nhất trong Linux. Nó được dùng để xem nội dung file trên output tiêu chuẩn (sdout). Để chạy lệnh này, gõ cat theo sau là tên file và phần mở rộng. Ví dụ: cat file.txt.

Có nhiều cách để sử dụng cat command linux:
- cat > filename tạo ra file mới
- cat filename1 filename2>filename3 nhập 2 files (1 và 2) để lưu kết quả vào file (3)
- để chuyển một file từ in thường tới in hoa hoặc ngược lại, cat filename | tr a-z A-Z >output.txt

### 3.5.2 echo
Lệnh này được dùng để chuyển dữ liệu vào một file. Ví dụ, nếu bạn muốn thêm text “Hello, my name is John” vào trong file  name.txt, bạn nhập ệnh echo Hello, my name is John >> name.txt

### 3.5.3 ls
Command ls được dùng để xem nội dung thư mục. Mặc định là command này sẽ hiển thị danh sách file trong thư mục hiện tại.

Nếu bạn muốn xem nội dung thư mục khác, hãy nhập ls và sau đó là đường dẫn thư mục. Ví dụ: nhập ls /home/username/Document để xem nội dung của Documents.

Có nhiều phiên bản để dùng với lệnh ls như sau:
- ls -R liệt kê các file bao gồm cả các thư mục phụ bên trong
- ls -a liệt kê những file ẩn
- ls -al liệt kê tất cả file và thư mục với thông tin chi tiết như phân quyền, kích thước, chủ sở hữu, vân vân.
  
### 3.5.4 nano
ano và vi là những trình soạn thảo được cài đặt sẵn trong dòng lệnh Linux. Lệnh nano là một trình soạn thảo văn bản khá tố biểu thị các từ khóa với màu sắc và có thể nhận dạng hầu hết các ngôn ngữ. Và vi là trình soạn thảo đơn giản hơn nano. Bạn có thể tạo một file hoặc chỉnh sửa một file bằng cách sử dụng những trình soạn thảo này. Ví dụ, nếu bạn cần tạo một file mới tên “checked.txt”, bạn có thể tạo nó bằng cách dùng lệnh “nano checked.txt”. Bạn có thể lưu file sau khi chỉnh sửa bằng cách dùng tổ hợp phím Ctr+X, sau đó Y (hoặc N). Dựa theo kinh nghiệm của mình, dùng nano cho việc chỉnh sửa code HTML không được tốt lắm, bởi vì màu sắc của nó.

### 3.5.5 ps
Gọi ra một bản log của các tiến trình hiện tại chạy trên hệ thống của linux, cho phép kiểm tra tiến trình nào đang chạy, chạy với thời gian bao nhiều. Tuy nhiên lại không có tính thời gian thực vì đây chỉ là một câu lệnh để log. Một số tùy chọn của lệnh này như
- ps –A: Kiểm tra mọi tiến trình trong hệ thống.
- ps -U root -u root –N: Kiểm tra mọi tiến trình ngoại trừ những tiến trình hệ thống.
- ps -u username: Kiểm tra những tiến trình được thực hiện bởi một người dùng nhất định.

### 3.5.6 top
Là terminal tương đương với Task Manager trong Windows, command top sẽ hiển thị danh sách tiến trình đang chạy và lượng CPU mà tiến trình đó sử dụng theo thời gian thực. Rất có ích khi bạn giám sát dung lượng lưu trữ tài nguyên trên hệ thống, đặc biệt là biết được quá trình nào cần chấm dứt vì tiêu thụ quá nhiều tài nguyên.

### 3.5.7 grep
Đây là một trong số các lệnh cơ bản trong Linux hữu ích được dùng hằng ngày. Command grep cho phép bạn tìm kiếm tất cả text thông qua tập tin nhất định.

Để minh họa, grep blue notepad.txt sẽ tìm từ blue trong file notepad. Các dòng có chứa từ được tìm sẽ hiển thị đầy đủ.
### 3.5.8 awk
Awk là một ngôn ngữ lập trình hỗ trợ thao tác dễ dàng đối với kiểu dữ liệu có cấu trúc và tạo ra những kết quả được định dạng. Nó được đặt tên bằng cách viết tắt các chữ cái đầu tiên của các tác giả: Aho, Weinberger và Kernighan.

Awk thường được sử dụng cho việc tìm kiếm và xử lý text. Nó sẽ tìm kiếm một hoặc nhiều file để xem xem trong các file đó có dòng nào bao gồm những pattern (khuôn mẫu) cho trước hay không, sau đó thực hiện những action (hành động) tương ứng.

Một số đặc điểm nổi bật của Awk:
- nó xem 1 file text giống như bảng dữ liệu, bao gồm các bản ghi và các trường
- tương tự những ngôn ngữ lập trình phổ biến, Awk cũng có những khái niệm như biến, điều kiện, vòng lặp
- Awk có những toán tử số học và toán tử thao tác chuỗi

Câu lệnh được viết với Awk sẽ nhận đầu vào là một file hoặc một input có dạng chuẩn, rồi tạo ra output theo chuẩn của nó. Awk chỉ làm việc với các file text.

Cú pháp cơ bản của câu lệnh được viết với Awk sẽ như sau:
~~~~~~~
awk '/search pattern 1/ {Actions}
         /search pattern 2/ {Actions}' file
~~~~~~~
Giải thích:
- search pattern là một biểu thức chính quy
- Actions là những câu lệnh sẽ được thực hiện
- file: file đầu vào Awk sẽ chấp nhận một vài kiểu pattern và action.

### 3.5.9 curl
Lệnh Curl là gì? Nó là chữ viết tắt của “Client URL”, dùng để kiểm tra kết nối tới URL và curl command thường dùng để truyền tải dữ liệu. Curl được cung cấp bởi libcurl, là một bộ thư viện truyền URL phía máy khách. 

Curl command hỗ trợ các giao thức bên dưới:
- HTTP và HTTPS
- FTP và FTPS
- IMAP và IMAPS
- POP3 và POP3S
- SMB và SMBS
- SFTP
- SCP
- TELNET
- GOPHER
- LDAP và LDAPS
- SMTP và SMTPS
Đây là những protocol được hỗ trợ quan trọng nhất, cũng có các giao thức khác ít phổ biến hơn.

### 3.5.10 wget
Wget là một công cụ máy tính tạo ra từ GNU Project. Bạn có thể dùng nó để trích xuất dữ liệu và nội dung từ nhiều web servers khác nhau. Tên của nó là kết hợp của World Wide Web và từ get. Nó hỗ trợ download qua FTP, SFTP, HTTP, và HTTPS.

Wget được tạo ra từ portable C và có thể dùng trên bất kỳ Unix system nào. Nó cũng có thể được triển khai trên Mac OS X, Microsoft Windows, AmigaOS và các nền tảng phổ biến khác.

### 3.5.11 netstat
Lệnh netstat trên linux là một lệnh nằm trong số các tập lệnh để giám sát hệ thống trên linux. netstat giám sát cả chiều in và chiều out kết nối vào server, hoặc các tuyến đường route, trạng thái của card mạng. lệnh netstat rất hữu dụng trong việc giải quyết các vấn đề về sự cố liên quan đến network như là lượng connect kết nối, traffic, tốc độ, trạng thái của từng port, Ip …

Một số trường hợp sử dụng netstat:
- netstat -a:   Kiểm tra tổng quát 
- netstat -at:  Kiểm tra các port đang sử dụng phương phức TCP
- netstat -au: Kiểm tra các port đang sử dụng phương phức UDP3
- netstat -l: Kiểm tra các port đang ở trạng thái listening
- netstat -lt: Kiểm tra các port đang listen dùng phương thức TCP
- netstat -lu: Kiểm tra các port đang listen dùng phương thức UDP
- netstat -plnt: Kiểm tra được port đang lắng nghe sử dụng dịch vụ gì ?
- netstat -rn: Hiển thị bảng định tuyến
- netstat -an | grep :80 | sort: Kiểm tra những kết nối thông qua port 80
- netstat -np | grep SYN_REC | wc -l: Kiểm tra có bao nhiêu gói SYN_REC trên server. Đối với con số thì tùy thuộc vào server của bạn, ví dụ nếu mỏi ngày có tầm 20 đến 30 kết nối, bổng dưng một ngày lên cả 100 -> 1000 kết nối thì bạn hiểu rồi đó -> server bị ddos.

### 3.5.12 sed
sed là một trong những công cụ mạnh mẽ trong Linux giúp chúng ta có thể thực hiện các thao tác với văn bản như tìm kiếm, chỉnh sửa, xóa.. Khác với các text editor thông thường, sed chấp nhận văn bản đầu vào có thể là nội dung từ một file có trên hệ thống hoặc từ standard input hay stdin. Chính vì vậy sed còn được gọi là một stream editor.
### 3.5.13 head
Lệnh head dùng để xem những dòng đầu của tệp tin (theo mặc định là 10 dòng đầu tiên). Chúng ta có thể thay đổi số dòng bằng cách thêm -n vào sau lệnh head. Cách dùng lệnh head:

>head [options] file

Trong đó tùy chọn có thể là:
- -n, --lines=[-]n: In số dòng n đầu tiên của mỗi tệp
- -c, --byte=[-]n: In số byte n đầu tiên của mỗi tệp
- -q: Không in tiêu đề xác định tên tệp
- -v: Luôn in tiêu đề xác định tên tệp
- --help: Hiển thị các trợ giúp
- --version: Thông tin về phiên bản và thoát

Ngoài ra chúng ta có thể tìm hiểu thêm các tùy chọn sử dụng lệnh: man head

### 3.5.13 tail
Lệnh tail dùng để xem những dòng cuối của tệp tin (theo mặc định là 10 dòng). Lệnh tail rất hữu ích khi khắc phục sự cố tệp nhật ký. Cách dùng lệnh tail:

> tail [tuỳ chọn] file

Trong đó tùy chọn có thể là:
- -n, --lines=[-]n: In số dòng n cuối cùng của mỗi tệp
- -n, --lines=[+]n: In tất cả các dòng từ n về sau
- -c, --byte=[-]n: In số byte n đầu cuối cùng mỗi tệp
- -q: Không in tiêu đề đầu ra
- -f: Tiếp tục đọc tập tin cho đến khi CTRL + C
- --help: Hiển thị các trợ giúp
- --version: Thông tin về phiên bản và thoát

> Chú ý: Lệnh tail -f này rất hữu ích khi dùng để theo dõi trực tiếp các file log.

Ngoài ra chúng ta có thể tìm hiểu thêm các tùy chọn sử dụng lệnh: man head
### 3.5.14 find
Câu lệnh find trong Linux được dùng để tìm kiếm tập tin và thư mục dựa trên các điều kiện đầu vào khác nhau. Tương tự như câu lệnh ls thì câu lệnh find cũng là một trong những câu lệnh được sử dụng phổ biến trên Linux.

Cú Pháp

> $ find [-H] [-L] [-P] [-D debugopts] [-Olevel] [path...] [expression]

Trong đó:
- Tuỳ chọn -P (là tuỳ chọn mặc định): không tìm kiếm các tập tin hay thư mục lối tắt liên kết tới tập tin hay thư mục khác (hay còn gọi là symbolic link).
- Tuỳ chọn -L (là tuỳ chọn mặc định): tìm kiếm tất cả các tập tin hay thư mục bao gồm các lối tắt liên kết tới tập tin hay thư mục khác.
- Tuỳ chọn -H: không tìm kiếm các tập tin hay thư mục lối tắt liên kết tới tập tin hay thư mục khác trừ khi xử lý đối số truyền vào câu lệnh.
- Tuỳ chọn -D (Debug Options): quy định các tuỳ chọn giúp gỡ rối.
- Tuỳ chọn -O (level): sử dụng khi muốn tối ưu tốc độ của câu lệnh.

### 3.5.15 ssh
SSH, hoặc được gọi là Secure Shell, là một giao thức điều khiển từ xa cho phép người dùng kiểm soát và chỉnh sửa server từ xa qua Internet. Dịch vụ được tạo ra nhằm thay thế cho trình Telnet vốn không có mã hóa và sử dụng kỹ thuật cryptographic để đảm bảo tất cả giao tiếp gửi tới và gửi từ server từ xa diễn ra trong tình trạng mã hóa. Nó cung cấp thuật toán để chứng thực người dùng từ xa, chuyển input từ client tới host, và relay kết quả trả về tới khách hàng.

Lệnh SSH có 3 phần:
> ssh {user}@{host}

SSH key command cho hệ thống biết là bạn muốn mở một kết nối được mã hóa Secure Shell Connection. {user} đại diện cho tài khoản người dùng bạn muốn dùng để truy cập. Ví dụ, bạn muốn truy cập user root, thì thay root tại đây. User root là user quản trị hệ thống với toàn quyền để chỉnh sửa bất kỳ điều gì trên hệ thống. {host} đại diện cho máy tính bạn muốn dùng để truy cập. Nó có thể là một địa chỉ IP (ví dụ 244.235.23.19) hoặc một tên miền (ví dụ, www.xyzdomain.com).

Khi bạn nhấn enter, nó sẽ hỏi bạn nhập mật khẩu tương ứng cho tài khoản. Khi bạn gõ, bạn sẽ không thấy bất kỳ dấu hiệu nào trên màn hình, nhưng nếu bạn gõ đúng mật khẩu và nhấn enter, bạn sẽ vào được hệ thống và nhận thông báo đăng nhập thành công.

### 3.5.16 kill
Khi thực hiện lệnh Kill, điều đó có nghĩa bạn đang gửi tín hiệu đến hệ thống để nó chấm dứt ứng dụng bị lỗi. 

Trong đó: 
- Sigterm: yêu cầu một quá trình ngừng chạy. Sử dụng tín hiệu này tức quá trình được lệnh ngừng đúng cách, có nghĩa là nó lưu lại các tiến trình rồi mới tắt.
- Sigkill: yêu cầu quá trình ngừng ngay lập tức. Ngược lại với Sigterm, tiến trình chưa được lưu và bị tắt đột ngột nên dữ liệu sẽ bị mất.

Cú pháp của kill: 
> kill[tín hiệu hoặc tùy chọn] PID(s)

Thông thường tín hiệu mặc định là Sigterm. Nếu vì lý do gì đó mà cấu hình mặc định này bị thay đổi thì bạn dùng cú pháp sau:
> Kill Sigkill PID hoặc kill -9 PID

(-9: là tín hiệu Sigkill)

Còn nếu không nhận thức được PID của ứng dụng thì chạy lệnh: 
> ps ux

Bạn cũng có thể dùng lệnh kill cho nhiều quy trình cùng một lúc.
> kill -9 PID1 PID2 PID3

Khi muốn kill biểu thức chính quy mở rộng, bạn dùng lệnh Pkill thay cho PID. 
> pkill firefox 

Khi muốn tiêu diệt tất cả các cá thể của một tiến trình, bạn sử dụng killall.
> killall firefox.

xkill dùng để giết một ứng dụng được chỉ định đích danh. Khi gõ xkill vào terminal, con trỏ chuột sẽ trở thành cross. Lúc này việc của bạn là nhấp vào ứng dụng bị lỗi và diệt chúng ngay lập tức. Ngoài ra, khi muốn kích hoạt chức năng xkill, bạn cũng có thể thêm lối tắt bàn phím.
### 3.5.17 dpkg
dpkg (Debian Package) là một công cụ cấp thấp. APT (Advanced Packaging Tool), một công cụ cấp cao hơn,được sử dụng nhiều hơn dpkg vì nó có thể tim các gói gói tin từ xa và giải quyết các gói tin có quan hệ phức tạp, chẳng hạn như các gói tin phụ thuộc. Frontends cho APT như aptitude (ncurses) và synaptic (GTK+) được sử dụng các giao diện thân thiện của chúng.
### 3.5.18 apt-get
Advanced Packaging Tool, hay APT, là phần mềm miễn phí dùng để quản lý việc cài đặt phần mềm trên Linux[1]. APT làm đơn giản các thủ tục quản lý phần mềm trên các máy tính tựa Unix bằng cách tự động hóa việc tải về, cấu hình và cài đặt các gói phần mềm, cả ở dạng biên dịch sẵn (dạng binary) hoặc biên dịch mã nguồn.[1]

APT ban đầu được thiết kế như là một giao diện cho dpkg để làm việc với các gói .deb của Debian, nhưng nó cũng đã được thay đổi để có thể làm việc với hệ thống RPM thông qua apt-rpm[2]. Dự án Fink đã chuyển APT lên hệ điều hành Mac OS X với một số chức năng quản lý gói, và APT cũng đã có trên OpenSolaris (bao gồm bản phân phố Nexenta OS).

"apt" không phải là một chương trình, nó chỉ là tên của một gói chứa tập hợp các công cụ (và các thư viện cần thiết) để thực hiện chức năng của nó. Phần quan trọng nhất của APT là các hàm thư viện C++ (nằm trong một gói khác có tên là libapt) dùng bởi các phần mềm giao diện người dùng cuối có liên quan để xử lý các gói, thí dụ như apt-get và apt-cache. Nó được dùng phổ biến như là một thí dụ về sự đơn giản và thống nhất; apt-get và apt-cache là phần "quan trọng" trong tất cả các phiên bản Debian, do đó nó được cài đặt mặc định trong Debian. APT về mặt chức năng có thể xem như là một giao diện đầu cuối của dpkg, và thân thiện hơn so với dselect. Trong khi dpkg chỉ thực hiện trên từng gói, APT quản lý các mối quan hệ (đặc biệt là sự phụ thuộc) giữa các gói, cũng như là quản lý nguồn phần mềm và quản lý phiên bản (theo dõi các bản phát hành và version pinning).

APT thường được ca ngợi như là một chức năng tốt nhất của Debian[4].

Một chức năng chính của APT là các nó gọi dpkg - nó thực hiện sắp xếp phân cấp danh sách các gói để có thể cài đặt hoặc gỡ và gọi dpkg theo một trình tự tốt nhất. Một số trường hợp nó sử dụng tham số --force của dpkg. Tuy nhiên, nó chỉ làm vậy khi không thể tìm được cách buộc dpkg phải thực thi.

Lệnh apt-get được sử dụng nhiều nhất là apt-get install tên gói (thông thường tên gói chỉ đơn giản là tên của một ứng dụng có thể chạy được, apt-get update, apt-get upgrade và apt-get dist-upgrade.
### 3.5.19 free
Kiểm tra trạng thái của bộ nhớ
### 3.5.20 df
Nếu muốn biết có bao nhiêu dung lượng ổ đĩa trống trên PC Linux, điều đầu tiên bạn làm là gì? Nhiều người có thể kiểm tra công cụ quản lý ổ đĩa của bản phân phối Linux đang dùng, nhưng đây không phải là tùy chọn tốt nhất.

Cú pháp:
>df [OPTION]... [FILE]...

Đây là lệnh df, được nhập trong terminal bash. Nó có thể được sử dụng theo nhiều cách:
- Hiển thị việc sử dụng không gian ổ đĩa
- Hiển thị việc sử dụng trên một thiết bị cụ thể
- Hiển thị tất cả các hệ thống file đang sử dụng
- Tìm thông tin ổ đĩa trên một file cụ thể
- Hiển thị các inode
- Hiển thị hệ thống file trên một file cụ thể
- Bao gồm các loại hệ thống file
- Loại trừ các loại hệ thống file
- Hiển thị Help File (phần giải thích) cho lệnh df
- Hiển thị việc sử dụng tổng dung lượng ổ đĩa và dung lượng trống
- Lệnh df cũng có thể được điều chỉnh để hiển thị kích thước file và không gian trống tính bằng kilobyte, megabyte và gigabyte.

