# Java Design Pattern – Observer
## 1. Observer Pattern là gì?
> Define a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.

Observer Pattern là một trong những Pattern thuộc nhóm hành vi (Behavior Pattern). Nó định nghĩa mối phụ thuộc một – nhiều giữa các đối tượng để khi mà một đối tượng có sự thay đổi trạng thái, tất các thành phần phụ thuộc của nó sẽ được thông báo và cập nhật một cách tự động.

Observer có thể đăng ký với hệ thống. Khi hệ thống có sự thay đổi, hệ thống sẽ thông báo cho Observer biết. Khi không cần nữa, mẫu Observer sẽ được gỡ khỏi hệ thống.

![Một số ví dụ](https://gpcoder.com/wp-content/uploads/2018/12/design-patterns-observer-example-1.png)

- Hình 1-8, cho phép observer thứ 1 đăng ký với hệ thống.
- Hình 1-9, cho phép observer thứ 2 đăng ký với hệ thống.
- Hiện tại hệ thống đang liên lạc với 2 observer: Observer 1 và Observer 2. Khi hệ thống phát sinh một sự kiện cụ thể nào đó, nó sẽ thông báo (notification) với cả 2 observer như hình số 1-10.

Observer Pattern còn gọi là Dependents, Publish/Subcribe hoặc Source/Listener.

## 2. Các thành phần trong Observer Pattern
![Các thành phần trong Observer Pattern](https://viblo.asia/uploads/99269945-2bd6-43ee-be8a-75fcb6c753cc.jpg)

**Subject**
- Biết danh sách không giới hạn các observers của nó.
- Cung cấp một giao diện để có thể thêm và loại bỏ observer.

**Observer**
- Định nghĩa một phương thức update() cho các đối tượng sẽ được subject thông báo đến khi có sự thay đổi trạng thái.

**ConcreteSubject**
- Cài đặt các phương thức của Subject
- Lưu trữ trạng thái danh sách các ConcreateObserver.
- Gửi thông báo đến các observer của nó khi có sự thay đổi trạng thái.

**ConcreteObserver**
- Cài đặt các phương thức của Observer
- Có thể duy trì một liên kết đến đối tượng ConcreteSubject.
- Lưu trữ trạng thái của subject.
- Thực thi việc cập nhật để giữ cho trạng thái đồng nhất với subject gửi thông báo đến.

Ta có thể hình dung sự tương tác giữa subject và các observer như sau:
- Khi subject có sự thay đổi trạng thái, nó sẽ duyệt qua danh sách các observer của nó và gọi phương thức cập nhật trạng thái ở từng observer, có thể truyền chính nó vào phương thức để các observer có thể lấy ra trạng thái của nó và xử lý.
![Hoạt động của các thành phần](https://viblo.asia/uploads/92977200-658d-48e7-b2f5-ae1e262add02.jpg)

## 3. Ví dụ Observer Pattern với ứng dụng Tracking thao tác một Account
Giả sử hệ thống của chúng ta cần theo dõi về tài khoản của người dùng. Mọi thao tác của người dùng đều cần được ghi log lại, sẽ thực hiện gửi mail thông báo khi tài khoản hết hạn, thực hiện chặn người dùng nếu truy cập không hợp lệ.
![Ví dụ](https://gpcoder.com/wp-content/uploads/2018/12/design-patterns-observer-example-768x485.png)

Chương trình của chúng ta như sau:

- Subject : cung cấp các phương thức để thêm, loại bỏ, thông báo observer.
- AccountService : đóng vai trò là ConcreteSubject, sẽ thông báo tới tất cả các observers bất cứ khi nào có thao tác của người dùng liên quan đến đăng nhập, tài khoản hết hạn.
- Observer : định nghĩa một phương thức update() cho các đối tượng sẽ được subject thông báo đến khi có sự thay đổi trạng thái. Phương thức này chấp nhận đối số là SubjectState, cho phép các ConcreteObserver sử dụng dữ liệu của nó.
- Logger, Mailer và Protector là các ConcreteObserver. Sau khi nhận được thông báo rằng có thao tác với user và gọi tới phương thức update(), các ConcreteObserver sẽ sử dụng dữ liệu SubjectState để xử lý.

## 4. Ưu điểm của Observer Pattern
- Dễ dàng mở rộng với ít sự thay đổi : mẫu này cho phép thay đổi Subject và Observer một cách độc lập. Chúng ta có thể tái sử dụng các Subject mà không cần tái sử dụng các Observer và ngược lại. Nó cho phép thêm Observer mà không sửa đổi Subject hoặc Observer khác. Vì vậy, nó đảm bảo nguyên tắc Open/Closed Principle (OCP).
- Sự thay đổi trạng thái ở 1 đối tượng có thể được thông báo đến các đối tượng khác mà không phải giữ chúng liên kết quá chặt chẽ.
- Một đối tượng có thể thông báo đến một số lượng không giới hạn các đối tượng khác.

## 5. Nhược điểm của Observer Pattern
Bởi vì các Observer không biết về sự hiện diện của nhau, nó có thể gây tốn nhiều chi phí của việc thay đổi Subject, hoặc có thể gây ra trường hợp không mong muốn (Unexpected update) của Subject.

## 6. Sử dụng Observer Pattern khi nào?
- Thường được sử dụng trong mối quan hệ 1-n giữa các object với nhau. Trong đó một đối tượng thay đổi và muốn thông báo cho tất cả các object liên quan biết về sự thay đổi đó.
- Khi thay đổi một đối tượng, yêu cầu thay đổi đối tượng khác và chúng ta không biết có bao nhiêu đối tượng cần thay đổi, những đối tượng này là ai.
- Sử dụng trong ứng dụng broadcast-type communication.
- Sử dụng để quản lý sự kiện (Event management).
- Sử dụng trong mẫu mô hình MVC (Model View Controller Pattern) : trong MVC, mẫu này được sử dụng để tách Model khỏi View. View đại diện cho Observer và Model là đối tượng Observable.

## Tham khảo
[1] https://viblo.asia/p/design-pattern-observer-V3m5WO8W5O7

[2] https://gpcoder.com/4747-huong-dan-java-design-pattern-observer/