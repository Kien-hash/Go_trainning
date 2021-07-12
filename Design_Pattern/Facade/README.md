# Design Pattern – Facade
## 1. Facade Pattern là gì?
>Provide a unified interface to a set of interfaces in a subsystem. Facade defines a higher-level interface that makes the subsystem easier to use.

Facade Pattern là một trong những Pattern thuộc nhóm cấu trúc (Structural Pattern). Pattern này cung cấp một giao diện chung đơn giản thay cho một nhóm các giao diện có trong một hệ thống con (subsystem). Facade Pattern định nghĩa một giao diện ở một cấp độ cao hơn để giúp cho người dùng có thể dễ dàng sử dụng hệ thống con này.

Facade Pattern cho phép các đối tượng truy cập trực tiếp giao diện chung này để giao tiếp với các giao diện có trong hệ thống con. Mục tiêu là che giấu các hoạt động phức tạp bên trong hệ thống con, làm cho hệ thống con dễ sử dụng hơn.

Ví dụ khi không có Facade, hệ thống của chúng ta trông giống như sau:
![Hê thống chưa dùng Facade](https://gpcoder.com/wp-content/uploads/2018/11/facade-1.png)

Khi có Facade hệ thống sẽ trông như sau:
![Hệ thống có sử dụng Facade](https://gpcoder.com/wp-content/uploads/2018/11/facade-2.png)

## 2. Các thành phần của Facade
Các thành phần cơ bản của một Facade Pattern:
- Facade: biết rõ lớp của hệ thống con nào đảm nhận việc đáp ứng yêu cầu của client, sẽ chuyển yêu cầu của client đến các đối tượng của hệ thống con tương ứng.
- Subsystems: cài đặt các chức năng của hệ thống con, xử lý công việc được gọi bởi Facade. Các lớp này không cần biết Facade và không tham chiếu đến nó.
- Client: đối tượng sử dụng Facade để tương tác với các subsystem.

Các đối tượng Facade thường là Singleton bởi vì chỉ cần duy nhất một đối tượng Facade.
![Các thành phân của Facade](https://gpcoder.com/wp-content/uploads/2018/11/design-patterns-facade-diagram.png)

## 3. Ví dụ ứng dụng Facade vào một hệ thống bán hàng
Một công ty bán hàng online, chẳng hạn Tiki cung cấp nhiều lựa chọn cho khách hàng khi mua sản phẩm. Khi một sản phẩm được mua, nó sẽ qua các bước xử lý: lấy thông tin về tài khoản mua hàng, thanh toán, vận chuyển, gửi Email/ SMS thông báo.

![Hệ thống bán hàng sử dụng facade](https://gpcoder.com/wp-content/uploads/2018/11/design-patterns-facade-example-1024x472.png)

Ứng dụng của chúng ta được thiết kế với Facade Pattern, bao gồm các class như sau:
- Thông tin về tài khoản (AccountService) : lấy thông tin cơ bản của khách hàng thông qua email được cung cấp.
- Dịch vụ thanh toán (PaymentService) : có thể thanh toán thông qua Paypal, thẻ tín dụng (Credit Card), tài khoản ngân hàng trực tuyến (E-banking), Tiền mặt (cash).
- Dịch vụ vận chuyển (ShippingService) : có thể chọn Free Shipping, Standard Shipping, Express Shipping.
- Dịch vụ email (EmailService) : có thể gửi mail cho khách hàng về tình hình đặt hàng, thanh toán, vận chuyển, nhận hàng.
- Dịch vụ tin nhắn (SMS) : có thể gửi thông báo SMS cho khách hàng khi thanh toán online.
- ShopFacade : là một Facade Pattern, class này bao gồm các dịch vụ có bên trong hệ thống. Nó cung cấp một vài phương thức để Client có thể dễ dàng mua hàng. Tùy vào nghiệp vụ mà nó sẽ sử dụng những dịch tương ứng, chẳng hạn dịch vụ SMS chỉ được sử dụng nếu khách hàng đăng ký mua hàng thông qua hình thức thanh toán online (Paypal, E-banking, …).
- Client : là người dùng cuối sử dụng ShopFacade để mua hàng.

Như ta thấy phía Client chỉ cần gửi thông tin về ShopFacade là có thể sử dụng được dịch vụ của các hệ thống con. Nếu không có Facade, phía Client sẽ không biết sử dụng những dịch vụ nào để có thể mua được sản phẩm. Khi phát sinh thêm một dịch vụ sẽ rất khó khăn khi sửa đổi và code phía Client cũng sẽ bị ảnh hưởng.

## 4. Ưu điểm của Facade
- Giúp cho hệ thống của bạn trở nên đơn giản hơn trong việc sử dụng và trong việc hiểu nó, vì một mẫu Facade có các phương thức tiện lợi cho các tác vụ chung.
- Giảm sự phụ thuộc của các mã code bên ngoài với hiện thực bên trong của thư viện, vì hầu hết các code đều dùng Facade, vì thế cho phép sự linh động trong phát triển các hệ thống.
- Đóng gói tập nhiều hàm API được thiết kế không tốt bằng một hàm API đơn có thiết kế tốt hơn.

## 5. Sử dụng Facade Pattern khi nào?
- Khi hệ thống có rất nhiều lớp làm người sử dụng rất khó để có thể hiểu được quy trình xử lý của chương trình. Và khi có rất nhiều hệ thống con mà mỗi hệ thống con đó lại có những giao diện riêng lẻ của nó nên rất khó cho việc sử dụng phối hợp. Khi đó có thể sử dụng Facade Pattern để tạo ra một giao diện đơn giản cho người sử dụng một hệ thống phức tạp.
- Khi người sử dụng phụ thuộc nhiều vào các lớp cài đặt. Việc áp dụng Façade Pattern sẽ tách biệt hệ thống con của người dùng và các hệ thống con khác, do đó tăng khả năng độc lập và khả chuyển của hệ thống con, dễ chuyển đổi nâng cấp trong tương lai.
- Khi bạn muốn phân lớp các hệ thống con. Dùng Façade Pattern để định nghĩa cổng giao tiếp chung cho mỗi hệ thống con, do đó giúp giảm sự phụ thuộc của các hệ thống con vì các hệ thống này chỉ giao tiếp với nhau thông qua các cổng giao diện chung đó.
- Khi bạn muốn bao bọc, che giấu tính phức tạp trong các hệ thống con đối với phía Client.

## Tham khảo
[1] https://gpcoder.com/4604-huong-dan-java-design-pattern-facade/
