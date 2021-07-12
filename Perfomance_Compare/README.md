# So sánh hiệu năng giữa các ngôn ngữ Java - C - Golang - Python - ....
Trong phần so sánh hiệu năng này từ phần 1 - 4 em sử dụng 3 ngôn ngữ C++, Java, Golang. Ở phần cuối cùng em so sánh giữa Golang và Nodejs để có thể thấy sự khác biệt về hiệu năng giữa các ngôn ngữ khác nhau.

## 1. Benchmark phép Cộng
Cách thực hiện:

- Hàm truyền vào một số N
- Sử dụng vòng for để tính 1+2+3+...+N
- Tính thời gian thực hiện
- So sánh thời gian thực hiện với N rất lớn của các ngôn ngữ
    - N = 500.000
    - N = 1.000.000
    - N = 1.000.000.000
    - N = 10.000.000.000
    - N = 100.000.000.000
- Có thể thực hiện M lần rất lớn lặp lại để tính thời gian trung bình

Kết quả

| Thời gian                   | GoLang       | Java       | C++        |
| --------------------------- | ------------ | ---------- | ---------- |
| N = 500.000 (M = 10000)     | 666.201µs    | 0.6043(ms) | 0.001447   |
| N = 1.000.000 (M = 10000)   | 1.361673ms   | 1.2054(ms) | 0.002847   |
| N = 1.000.000.000 (M = 100) | 1.366895417s | 1.23       | 3.0026     |
| N = 10.000.000.000 (M = 10) | 13.5687559s  | 14.9       | 28.714600  |
| N = 100.000.000.000 (M = 1) | 136.7627742s | 242.0      | 280.350000 |

=> Nhận xét: Với các phép toán đơn giản và có số lượng nhiều như trên, Golang có hiệu năng vượt trội hơn so với 2 ngôn ngữ còn lại là C++ và Java

## 2. Benchmark nhiều phép toán
Cách thực hiện:

- Hàm truyền vào một số N
- Sử dụng vòng for để tính 1/23+2/34-3/45+...+(-1)^NN/(N+1)*(N+2)
- Tính thời gian thực hiện
- So sánh thời gian thực hiện với N rất lớn của các ngôn ngữ
    - N = 500.000
    - N = 1.000.000
    - N = 1.000.000.000
    - N = 10.000.000.000
    - N = 100.000.000.000
- Có thể thực hiện M lần rất lớn lặp lại để tính thời gian trung bình

Kết quả

| Thời gian                   | GoLang       | Java   | C++        |
| --------------------------- | ------------ | ------ | ---------- |
| N = 500.000 (M = 10000)     | 923.819(µs)  | 8.0E-4 | 0.001704   |
| N = 1.000.000 (M = 10000)   | 1.851808(ms) | 0.0016 | 0.003416   |
| N = 1.000.000.000 (M = 100) | 1.80417625   | 1.9    | 3.379300   |
| N = 10.000.000.000 (M = 10) | 18.1744078   | 26.0   | 33.881000  |
| N = 100.000.000.000 (M = 1) | 183.594165   | 261.0  | 340.089000 |

=> Nhận xét: Với các phép toán đơn giản và có số lượng nhiều như trên, Golang có hiệu năng vượt trội hơn so với 2 ngôn ngữ còn lại là C++ và Java
## 3. Benchmark tạo chuỗi Fibonaci
Cách thực hiện:
- Hàm truyền vào một số N
- Sử dụng thuật toán đệ quy để tìm số thứ N của dãy Fibonaci
- So sánh thời gian thực hiện với N rất lớn của các ngôn ngữ 
    - N = 20
    - N = 30
    - N = 40
    - N = 50
    - N = 60

Kết quả

| Thời gian          | GoLang         | Java       | C++        |
| ------------------ | -------------- | ---------- | ---------- |
| N = 20 (M = 10000) | 25.258(µs)     | 0.0162(ms) | 0.000036   |
| N = 30 (M = 10000) | 3.05284(ms)    | 1.7883(ms) | 0.004240   |
| N = 40 (M = 100)   | 364.922887(ms) | 0.22       | 0.5187     |
| N = 50 (M = 1)     | 45.2798885s    | 27.905     | 67.184000  |
| N = 55 (M = 1)     | 528.6157012s   | 308.0      | 727.705000 |

=> Nhận xét: với phép toán đệ quy, Java thể hiện sự vượt trội hơn về hiệu năng của mình so với 2 ngôn ngữ còn lại
## 4. Benchmark so sánh 2 xâu
Cách thực hiện:

- Hàm truyền vào một số N
- Tạo xâu N kí tự ngẫu nhiên, sau đó thực hiện so sánh xem 2 xâu có bao nhiêu kí tự giống nhau
- Tính thời gian thực hiện so sánh & thời gian tạo chuỗi
- So sánh thời gian thực hiện với N rất lớn của các ngôn ngữ
    - N = 500.000
    - N = 1.000.000
    - N = 1.000.000.000
    - N = 10.000.000.000
    - N = 100.000.000.000
- Có thể thực hiện M lần rất lớn lặp lại để tính thời gian trung bình

Kết quả

| Thời gian         | GoLang     | Java           | C++              |
| ----------------- | ---------- | -------------- | ---------------- |
| N = 500.000       | 13.9635ms  | 38ms           | 0.069000         |
| N = 1.000.000     | 26.9046ms  | 62.0ms         | 0.141000         |
| N = 1.000.000.000 | 26.048791s | std::bad_alloc | OutOfMemoryError |

=> Nhận xét: Với phép toán so sánh 2 xâu, do việc kiểm soát tốt hơn các dữ liệu mà Golang có thời gian thực hiện nhanh hơn cũng như không bị lỗi vùng nhớ như 2 ngôn ngữ kia
## 5. Benchmark RestAPI
- Triển khai mô hình:

![](https://res.cloudinary.com/practicaldev/image/fetch/s--zmwOgeuO--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://cdn-images-1.medium.com/max/318/1%2ACRk1t9DP5TOQd9NTxhQTxw.png)

- Tạo và test các api liên quan đến:
    - Insert Tạo mới bản ghi trong db
    - Update bản ghi trong db
    - Read Bản ghi
    - Delete bản ghi

- Thực hiện kiểm tra:
  - Tool sử dụng: Postman, ApacheBench
  - Kịch bản 1: thực hiện triển khai truy cập lần lượt: 10, 100, 1000, 2000 lần => đưa ra tổng thời gian truy cập
  - Kịch bản 2: Thực hiện nhiều truy cập cùng lúc: 500, 1000, 2000 => xét lấy bản ghi của 10000 request để thống kê lại

Kết quả kịch bản 1: (Đơn vị của thời gian ở đây là ms)

| Thời gian | Read (Go) | Read (Nodejs) | Insert (Go) | Insert (Nodejs) | Update (Go) | Update (Nodejs) | Delete (Go) | Delete (Nodejs) |
| --------- | --------- | ------------- | ----------- | --------------- | ----------- | --------------- | ----------- | --------------- |
| N = 10    | 21        | 56            | 358         | 344             | 324         | 350             | 345         | 344             |
| N = 100   | 326       | 492           | 2834        | 2994            | 2784        | 3087            | 2942        | 3082            |
| N = 1000  | 1307      | 1789          | 28640       | 30088           | 28689       | 30036           | 28662       | 30606           |
| N = 2000  | 2558      | 2986          | 55907       | 56600           | 56406       | 56411           | 55943       | 56744           |

=> Nhận xét: Hiệu năng của golang không thay đổi nhiều so với js, do cả 2 đều là ngôn ngữ bậc cao hỗ trợ mạnh cho việc tạo ứng dụng server nên việc tối ưu hiệu năng của 2 ngôn ngữ là như nhau.


Kết quả kịch bản 2: (Lấy 10000 mẫu để thống kê lại), giá trị được ghi trong bảng là REQUEST PER SECOND.

| Số lượng truy cập đồng thời | Read (Go) | Read (Nodejs) | Insert (Go) | Insert (Nodejs) | Update (Go) | Update (Nodejs) | Delete (Go) | Delete (Nodejs) |
| --------- | --------- | ------------- | ----------- | --------------- | ----------- | --------------- | ----------- | --------------- |
| N = 500   | 22880.53  | 3635.86       |4970.16      | 32.04           |1496.28      | 9.39            |613.71       | 2592.44         |
| N = 1000  | 17687.91  | 3635.86       |609.96       |  33.12          |579.20       | 10.99           |653.11       | 2729.40         |
| N = 2000  | 22127.52  | 3510.37       |304.91       |  32.85          |532.75       | 16.59           |573.21       | 1290.10        |
  
Thời gian để thực hiện được 95% số lượng request của từng ngôn ngữ (Đơn vị: ms)

| Số lượng truy cập đồng thời | Read (Go) | Read (Nodejs) | Insert (Go) | Insert (Nodejs) | Update (Go) | Update (Nodejs) | Delete (Go) | Delete (Nodejs) |
| --------- | --------- | ------------- | ----------- | --------------- | ----------- | --------------- | ----------- | --------------- |
| N = 500   | 19        | 165           |197          | 16481           |725          | 58970           |197          | 471             |
| N = 1000  | 42        | 296           |668          | 30630           |1226         | 92906           |726          | 741             |
| N = 2000  | 59        | 587           |16617        | 61626           |16405        | 121263          | 1332        | 3465         |

=> Nhận xét, với số lượng truy cập đồng thời lớn, golang cho thấy hiệu năng vượt trội so với các ngôn ngữ khác, như ta thấy ở bảng trên của kịch bản 2 tất cả các giá trị REQUEST đáp ứng trên 1 giây của GO luôn lớn hơn so với Nodejs (trừ trường hợp lệnh xóa, do cách test ), còn trong bảng 2 thời gian hoàn thành 95% giá trị yêu cầu của GO luôn ở dưới so với Nodejs, cho thấy hiệu năng tốt hơn của GO.
