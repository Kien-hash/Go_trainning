# Trong thư mục này gồm file code sau:
- config.js : chứa thông tin về cấu hình database được kết nối đến
- controller.js: mô tả tương tác giữa các phương thức của restfulAPI và database
- routes.js: mô tả thông tin về các đường dẫn cung cấp các restfulAPI
- server.js: file chạy của hệ thống nodejs

# Mô tả code: 
Code mô tả một các khái quát khi 1 request đến thì ta sẽ tương tác với database thế nào:
- method GET: trả về một hoặc tất cả bản ghi có trong bảng
- method PUT: thực hiện việc thay đổi database, có thể là thêm vào một bản ghi mới hay update một bản ghi mới trong bảng
- method DELETE: thực hiện xoá một bản ghi trong bảng

# Các thư viện sử dụng trong chương trình
- express: Thiết lập các lớp trung gian để trả về các HTTP request
- util: Cung cấp các hàm tiện ích (Utility functions).
- mysql: Thực hiện việc tạo liên kết với DB
- body-parser: Đẩy dữ liệu đã lấy từ DB lên HTML