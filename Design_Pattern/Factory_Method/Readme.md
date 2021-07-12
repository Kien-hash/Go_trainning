# Factory Pattern
## Factory Method Pattern là gì?
>Factory Method is a creational design pattern that Define an interface for creating an object, but let subclasses decide which class to instantiate. Factory Method lets a class defer instantiation to subclasses

Factory Method Design Pattern hay gọi ngắn gọn là Factory Pattern là một trong những Pattern thuộc nhóm Creational Design Pattern. Nhiệm vụ của Factory Pattern là quản lý và trả về các đối tượng theo yêu cầu, giúp cho việc khởi tạo đổi tượng một cách linh hoạt hơn.

Factory Pattern đúng nghĩa là một nhà máy, và nhà máy này sẽ “sản xuất” các đối tượng theo yêu cầu của chúng ta.

Trong Factory Pattern, chúng ta tạo đối tượng mà không để lộ logic tạo đối tượng ở phía người dùng và tham chiếu đến đối tượng mới được tạo ra bằng cách sử dụng một interface chung.

Factory Pattern được sử dụng khi có một class cha (super-class) với nhiều class con (sub-class), dựa trên đầu vào và phải trả về 1 trong những class con đó.
## Các thành phần cơ bản
Một Factory Pattern bao gồm các thành phần cơ bản sau:
![Thành phần cơ bản Factory Method](https://i.imgur.com/pugRJDC.png)

Trong đó:
- Product là một interface chung mà tất cả các product con đều phải implements.
- Mỗi concrete product là một product con cụ thể implements từ interface Product.
- Class Creator là nơi khai báo factory method có kiểu dữ liệu trả về là Product cho phép trả về các product mới.
- Mỗi Concrete Creator là một class kế thừa từ Creator có nhiệm vụ override factory method và trả về loại product tương ứng.

Một lưu ý nhỏ, factory method không nhất thiết phải luôn luôn tạo ra các đối tượng mới. Đôi khi nó cũng có thể trả về các đối tượng đã được tạo từ trước đó và đang được lưu trữ trong cache, object pool hoặc từ các nguồn khác.

## Ví dụ
Ví dụ: Tất cả hệ thống ngân hàng có cung cấp API để truy cập đến hệ thống của họ. Team được giao nhiệm vụ thiết kế một API để client có thể sử dụng dịch vụ của một ngân hàng bất kỳ. Hiện tại, phía client chỉ cần sử dụng dịch vụ của 2 ngân hàng là VietcomBank và TPBank. Tuy nhiên để dễ mở rộng sau này, và phía client mong muốn không cần phải thay đổi code của họ khi cần sử dụng thêm dịch vụ của ngân hàng khác. Với yêu cầu như vậy, chúng ta có thể sử dụng một Pattern phù hợp là Factory Method Pattern.

Hệ thống được minh họa như sau:

![UML](https://gpcoder.com/wp-content/uploads/2018/09/design-patterns-factory-method-diagram.png)

Trong đó:
- Product tương ứng ở đây là Bank.
- ConcreteProduct là TPBank và VietcomBank
- Creator ở đây là BankFactory
- Không có Concrete Creator vì ở đây ta làm một cấu trúc đơn giản nên Creator đóng vai trò làm Concrete Creator luôn
- Các thành phần khác: Client là nơi truy vấn Factory, BankType là một enum đóng vai trò là tham số truyền vào trong Creator để tạo ra các product cụ thể.

## Sử dụng Factory Pattern khi nào?
Factory Method Pattern phát huy được ưu điểm của nó trong một số trường hợp sau:
- Khi bạn chưa biết nên khởi tạo đối tượng mới từ class nào.
- Khi bạn muốn tập trung các đoạn code liên quan đến việc khởi tạo các đối tượng mới về cùng một nơi để dễ thao tác và xử lý.
- Khi bạn không muốn người dùng phải biết hết tên của các class có liên quan đến quá trình khởi tạo cũng như muốn che giấu, đóng gói toàn bộ logic của quá trình khởi tạo một đối tượng mới nào đó khỏi phía người dùng.

## Ưu điểm 
- Factory Method Pattern giúp hạn chế sự phụ thuộc giữa creator và concrete products.
- Factory Method Pattern giúp gom các đoạn code tạo ra product vào một nơi trong chương trình, nhờ đó giúp dễ theo dõi và thao tác.
- Với Factory Method Pattern, chúng ta có thể thoải mái thêm nhiều loại product mới vào chương trình mà không làm thay đổi các đoạn code nền tảng đã có trước đó.

## Nhược điểm
- Code có thể trở nên nhiều hơn và phức tạp hơn do đòi hỏi phải sử dụng nhiều class mới có thể cài đặt được pattern này.
