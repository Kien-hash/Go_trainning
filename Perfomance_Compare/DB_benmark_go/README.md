# Mô tả code
File main.go có chứa 4 hàm mô tả cách xử lý trên database đối với các request được gửi đến.
- deleteTable(response, request) thực hiện xoá đi hàng được chọn trong bảng users. Dữ liệu về hàng được chọn để xoá được giải mã từ request.Body 
- updateTable(response, request) thực hiện update đi hàng được chọn trong bảng users. Dữ liệu về hàng được chọn để update được giải mã từ request.Body 
- readTable(response, request) thực hiện đọc tất cả các hàng trong bảng users. Dữ liệu được mả hoá thành json và gửi về thông qua respond
- insertTable(respond, request) thực hiện việc chèn một hàng mới vào bảng. 

# Thư viện sử dụng
- database/sql: Thư viện thực hiện kết nối với database kiểu mysql
- github.com/gorilla/mux: Thư viện thực hiện đơn giản hoá việc map giữa các API và hàm xử lý tương ứng
- encoding/json: Xử lý mã hoá dữ liệu của sql thành json và ngược lại
