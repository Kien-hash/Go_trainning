# Java Design Pattern – Builder
## 1. Builder Pattern là gì?
> Builder is a creational design pattern that separate the construction of a complex object from its representation so that the same construction process can create different representations.

Builder pattern là một trong những Creational pattern. Builder pattern là mẫu thiết kế đối tượng được tạo ra để xây dựng một đôi tượng phức tạp bằng cách sử dụng các đối tượng đơn giản và sử dụng tiếp cận từng bước, việc xây dựng các đối tượng đôc lập với các đối tượng khác.

Builder Pattern được xây dựng để khắc phục một số nhược điểm của Factory Pattern và Abstract Factory Pattern khi mà Object có nhiều thuộc tính.

Có ba vấn đề chính với  Factory Pattern và Abstract Factory Pattern khi Object có nhiều thuộc tính:
- Quá nhiều tham số phải truyền vào từ phía client tới Factory Class.
- Một số tham số có thể là tùy chọn nhưng trong Factory Pattern, chúng ta phải gửi tất cả tham số, với tham số tùy chọn nếu không nhập gì thì sẽ truyền là null.
- Nếu một Object có quá nhiều thuộc tính thì việc tạo sẽ phức tạp.

Chúng ta có thể xử lý những vấn đề này với một số lượng lớn các tham số bằng việc cung cấp một hàm khởi tạo với những tham số bắt buộc và các method getter/ setter để cài đặt các tham số tùy chọn. Vấn đề với hướng tiếp cận này là trạng thái của Object sẽ không nhất quán cho tới khi tất cả các thuộc tính được cài đặt một cách rõ ràng. Nếu cần xây dựng một đối tượng Immutable thì cách này cũng không thể thực hiện được.

## 2. Các thành phần trong Builder Pattern
![Các thành phần](https://images.viblo.asia/c6465c40-59f4-4800-b30b-fd060d3d92ab.gif)

Một builder gồm các thành phần cơ bản sau:
- Product : đại diện cho đối tượng cần tạo, đối tượng này phức tạp, có nhiều thuộc tính.
- Builder : là abstract class hoặc interface khai báo phương thức tạo đối tượng.
- ConcreteBuilder : kế thừa Builder và cài đặt chi tiết cách tạo ra đối tượng. Nó sẽ xác định và nắm giữ các thể hiện mà nó tạo ra, đồng thời nó cũng cung cấp phương thức để trả các các thể hiện mà nó đã tạo ra trước đó.
- Director/ Client: là nơi sẽ gọi tới Builder để tạo ra đối tượng.

Trường hợp đơn giản, chúng ta có thể gộp Builder và ConcreteBuilder thành static nested class bên trong Product.

## 3. Ví dụ sử dụng Builder trong việc đặt đồ ăn nhanh
![UML của code ví dụ](https://i.ibb.co/rZQ5bPd/Untitled-Diagram-1.png)

Trong đó:
- Builder: là interface OrderBuilder nhằm khai báo các phương thức cân thiết để khởi tạo
- ConcreteBuilder: là FastFoodOrderBuilder, là nơi triển khai cụ thể các phương thức khởi tạo đối với từng tham số của một Order như orderType(), orderBread(),...
- Product: là class Order
- Client: là nơi gọi tới Builder để tạo đối tượng, ở đây cho là hàm main().
- Các thành phần khác: Các enum đơn giản là những tham số truyền vào cho Builder


## 4. Ưu điểm của Builder Pattern 
- Hỗ trợ, loại bớt việc phải viết nhiều constructor.
- Code dễ đọc, dễ bảo trì hơn khi số lượng thuộc tính (propery) bắt buộc để tạo một object từ 4 hoặc 5 propery.
- Giảm bớt số lượng constructor, không cần truyền giá trị null cho các tham số không sử dụng.
- Ít bị lỗi do việc gán sai tham số khi mà có nhiều tham số trong constructor: bởi vì người dùng đã biết được chính xác giá trị gì khi gọi phương thức tương ứng.
- Đối tượng được xây dựng an toàn hơn: bởi vì nó đã được tạo hoàn chỉnh trước khi sử dụng.
- Cung cấp cho bạn kiểm soát tốt hơn quá trình xây dựng: chúng ta có thể thêm xử lý kiểm tra ràng buộc trước khi đối tượng được trả về người dùng.
- Có thể tạo đối tượng immutable.

## 5. Nhược điểm của Builder Pattern
Builder Pattern có nhược điểm là duplicate code khá nhiều: do cần phải copy tất cả các thuộc tính từ class Product sang class Builder.

Tăng độ phức tạp của code (tổng thể) do số lượng class tăng lên.

## 6. So sánh Builder Pattern với Factory/ Abstract Factory Pattern
Factory Pattern cũng có thể được sử dụng để xây dựng một đối tượng phức tạp, vậy sự khác biệt của nó với mô hình Builder Pattern là gì?

Sự khác biệt lớn duy nhất giữa Builder Pattern và Factory Pattern cung cấp cho bạn nhiều quyền kiểm soát hơn đối với quá trình tạo đối tượng.

>Factory/ Abstract Factory Pattern là câu trả lời cho “WHAT” và Builder Pattern là câu trả lời cho “HOW“.

Trong Builder Pattern, đối tượng được xây dựng từng bước (step by step). Builder Pattern có nhiều bước nhỏ, mỗi bước sẽ có các đơn vị logic nhỏ kèm theo trong đó. Cũng sẽ có một chuỗi (sequence) liên quan. Nó sẽ bắt đầu từ bước 1 và sẽ đi lên tối đa bước n và bước cuối cùng là trả về đối tượng. Nhưng trong Factory Pattern, bạn sẽ không thấy được đối tượng phức tạp được tạo như thế nào, nó không có từng bước xây dựng đối tượng.

## Tài liệu tham khảo
[1] https://gpcoder.com/4434-huong-dan-java-design-pattern-builder/
