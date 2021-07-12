# Tìm hiểu về viết Unit Test trong project với Golang
# 1. Tổng quan
Golang cung cấp cho ta 2 thư viện mặc định để thực hiện việc viết test:
- [testing](https://godoc.org/testing): Cung cấp các chức năng kiểm thử tự động cơ bản
- [httptest](https://godoc.org/net/http/httptest): Cung cấp chức năng kiểm thử đối với giao thức http

Ngoài ra còn có nhiều thư viện của bên thứ 3 cho phép ta thực hiện việc viết test như: [gocheck](https://github.com/go-check/check), [ginkgo](https://github.com/onsi/ginkgo), ...

Các file test khi dùng package `testing` có phần đuôi là `_test.go`. VD nếu ta muốn test file `hello.go` thì file test tương ứng ta cần phải đặt tên là `hello_test.go`. Hàm test trong file này sẽ có dạng:
```golang
// Xxx không được bắt đầu bằng chữ thường
func TestXxx(*testing.T)
```
Hàm Test truyền vào một tham số kiểu con trỏ kiểu `T` của thư viện `testing`. Kiểu `T` chứa các phương thức hỗ trợ việc quản lý trạng thái của testcase cũng như hỗ trợ việc format logs in ra của testcase thông qua các phương thức `Error`, `Errorf`, `Fail`, `Log`, v.v. Chi tiết tất cả phương thức của kiểu `T` và cách sử dụng các bạn có thể xem chi tiết ở [đây](https://godoc.org/testing#T).
```golang
type T struct {
	common
	isParallel bool
	context    *testContext // For running tests and subtests.
}
```
# 2. Viết testcase với thư viện testing
Xét ví dụ:
```golang
// hello.go
package main

import "fmt"

func hello(name string) string {
	if name == "" {
		return fmt.Sprintf("What is your name ?")
	} else {
		return fmt.Sprintf("Hello %s", name)
	}
}
```
```golang
// hello_test.go
package main

import "testing"

func TestHello(t *testing.T) {
	emptyNameResult := hello("")

	if emptyNameResult != "What is your name ?" {
		t.Errorf("Output expect What is your name ? instead of %v", emptyNameResult)
	}

	result := hello("Gopher")

	if result != "Hello Gopher" {
		t.Errorf("Output expect Hello Gopher instead of %v", result)
	}
}
```
- Hàm `hello` thực hiện chức năng cơ bản là trả về `Hello` + `tham số name` truyền vào. Nếu `name` rỗng thì trả về chuỗi `What is your name?`.
- Hàm `TestHello` kiểm tra 2 khả năng đầu ra của hàm `hello`. Nếu giá trị trả về khác với đầu ra kỳ vọng thì test sẽ failed.
- Phương thức `t.Errorf` cũng như `t.Error` có chức năng đánh dấu các test failed nếu đầu ra không đúng kỳ vọng. Sau khi thực thi xong, code trong hàm `TestHello` vẫn tiếp tục được thực thi. `t.Errorf` khác `t.Error` giống như `fmt.Printlnf` khác `fmt.Println`.

Chạy test case, hàm trả về giá trị giống với kì vọng:
![giá trị trả về](https://images.viblo.asia/1aa35e89-6095-4b69-9e0a-b3a2b0190b9c.png)

Khi ta sửa lại hàm `hello` và giữ nguyên hàm `TestHello`:
```golang
func hello(name string) string {
	if name == "" {
		return fmt.Sprintf("You do not have name !")
	} else {
		return fmt.Sprintf("Hello %s", name)
	}
}
```

Khi đó giá trị của hàm testcase là:
![Giá trị của testcase khi thay đổi hàm hello](https://images.viblo.asia/d2905698-2e79-4729-bfa3-9f664f53ffd7.png)

Output thực tế là `You do not have name !` thay vì `What is your name ?` như kỳ vọng.

**Kiểm tra test coverage**

![Kiểm tra test coverage](https://images.viblo.asia/ce169a60-b96c-4a59-8482-f2f15e502b79.png)

# 3. Viết testcase với thư viện httptest

```golang
// server.go
package main

import (
	"fmt"
	_ "encoding/json"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "Welcome to our website")
	return
}

func main() {
	http.HandleFunc("/", welcome)
	
	http.ListenAndServe(":3000", nil)
}
```

```golang
// server_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWelcome(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", welcome)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
```
- Hàm `Welcome` đơn giản thực hiện trả về `client` thông điệp `Welcome to our website` khi người dùng truy cập vào `http//localhost:3000`.
- Hàm `TestWelcome` sẽ cần tạo một webserver riêng (với multiplexer mặc định của thư viện `net/http`). Kết quả status code trả về được kiểm tra, nếu khác `200` thì test sẽ failed.

# 3. Viết testcase bằng thư viện của bên thứ 3

# 4. Kiểm tra sự đầy đủ của các testcase
## 4.1. Test cover
Golang cung cấp cho chúng ta một công cụ là test cover để có thể tìm hiểu xem trong một file `*_test.go` ta đã có thể thực thi được hết tất cả các dòng code của file `*.go` tương ứng chưa.
### B1: Kiểm tra tổng quát bằng test coverage.

```terminal
go test -cover
```
Nếu kết quả trả về không đạt 100% như dưới đây, ta chuyển sang B2

![Kết quả thực hiện <100%](/Unitest/image/Screenshot%20from%202020-09-30%2011-38-27.png)

### B2: Ghi lại thông tin trong test vào một file `cover.out`
```terminal
go test -coverprofile cover.out
```
File này sẽ chứa thông tin về lần mà ta test được, nhưng ta khó có thể đọc trực tiếp file này, cũng như khó có cái nhìn trực quan đối với code của mình. Vì thế ta sẽ sử dụng 2 công cụ của bước 3 và bước 4 như sau.

### B3: Thống kê theo % mức độ thực hiện đầy đủ của các hàm

```terminal
go tool cover -func cover.out
```
 
![Thống kế thực hiện](/Unitest/image/Screenshot%20from%202020-09-30%2011-46-30.png)

Như ta có thể thấy hàm `dbOpen()` được đánh số 100%, tức là 
- _**tất cả các câu lệnh trong hàm**_ được thực hiện đầy đủ.
- _**tất cả những lần hàm được gọi**_ được thực hiện đầy đủ.

Đối với các hàm không đạt 100% thì một trong 2 khả năng chúng sẽ không được thực hiện đầy đủ như 2 lần trên, VD hàm `checkError()` chỉ đạt 50%, tức là chỉ 1 nửa các câu lệnh của hàm được thực hiện trong `*_test.go`. Có thể là do file test không thực hiện được hết tất cả các rẽ nhánh `if` của logic hàm `checkError()` trong file `*.go` tương ứng. Hoặc cũng có thể ta gọi hàm `checkError()` 2 lần nhưng chỉ 1 lần được thực hiện được gọi.

### B4: Thực hiện truy tìm những lệnh được gọi trong golang.
Sau khi biết được lượng %, ta có thể truy vết kĩ hơn bằng các sử dụng go tool để truy vấn đến file `cover.out` ta tạo ra ở B2, để cho ta một cái nhìn tổng quát về những dòng lệnh được chạy và không được chạy.

```Terminal
go tool cover -html=cover.out
```

![](image/Screenshot%20from%202020-09-30%2014-04-59.png)

Sau khi chờ một chút, browser sẽ tự hiện lên một file html được lưu trong hệ thống của ta như trong hình. Ta có thể thấy một cách trực quan rằng file test của ta đã sử dụng những dòng code màu xanh, chưa dùng đến dòng code màu đỏ và không dùng đến đoạn code màu xám. Từ đó ta có thể thực hiện việc điều chỉnh file test một cách dễ dàng hơn.

## 4.2. Test đối với websocket

 