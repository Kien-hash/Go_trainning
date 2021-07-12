# Adapter Pattern
## 1.Adapter Pattern là gì?
>Adapter Pattern is a structural design pattern that convert the interface of a class into another interface clients expect. Adapter lets classes work together that couldn’t otherwise because of incompatible interfaces.

Adapter Pattern (Người chuyển đổi) là một trong những Pattern thuộc nhóm cấu trúc (Structural Pattern). Adapter Pattern cho phép các inteface (giao diện) không liên quan tới nhau có thể làm việc cùng nhau. Đối tượng giúp kết nối các interface gọi là Adapter.

Adapter Pattern giữ vai trò trung gian giữa hai lớp, chuyển đổi interface của một hay nhiều lớp có sẵn thành một interface khác, thích hợp cho lớp đang viết. Điều này cho phép các lớp có các interface khác nhau có thể dễ dàng giao tiếp tốt với nhau thông qua interface trung gian, không cần thay đổi code của lớp có sẵn cũng như lớp đang viết.

![](https://gpcoder.com/wp-content/uploads/2018/10/adapter-pattern-example.jpg)

Adapter Pattern còn gọi là Wrapper Pattern do cung cấp một interface “bọc ngoài” tương thích cho một hệ thống có sẵn, có dữ liệu và hành vi phù hợp nhưng có interface không tương thích với lớp đang viết.

## 2. Các thành phần trong Adapter Pattern
Một Adapter Pattern bao gồm các thành phần cơ bản sau:
- Adaptee: định nghĩa interface không tương thích, cần được tích hợp vào.
- Adapter: lớp tích hợp, giúp interface không tương thích tích hợp được với interface đang làm việc. Thực hiện việc chuyển đổi interface cho Adaptee và kết nối Adaptee với Client.
- Target: một interface chứa các chức năng được sử dụng bởi Client (domain specific).
- Client: lớp sử dụng các đối tượng có interface Target.

Có hai cách để thực hiện Adapter Pattern dựa theo cách cài đặt (implement) của chúng:
* Object Adapter – Composition (Chứa trong): trong mô hình này, một lớp mới (Adapter) sẽ tham chiếu đến một (hoặc nhiều) đối tượng của lớp có sẵn với interface không tương thích (Adaptee), đồng thời cài đặt interface mà người dùng mong muốn (Target). Trong lớp mới này, khi cài đặt các phương thức của interface người dùng mong muốn, sẽ gọi phương thức cần thiết thông qua đối tượng thuộc lớp có interface không tương thích.

![Object Adapter - Composition](https://gpcoder.com/wp-content/uploads/2018/10/design-patterns-object-adapter-diagram.png)

* Class Adapter – Inheritance (Kế thừa) : trong mô hình này, một lớp mới (Adapter) sẽ kế thừa lớp có sẵn với interface không tương thích (Adaptee), đồng thời cài đặt interface mà người dùng mong muốn (Target). Trong lớp mới, khi cài đặt các phương thức của interface người dùng mong muốn, phương thức này sẽ gọi các phương thức cần thiết mà nó thừa kế được từ lớp có interface không tương thích.

![Class Adapter - Inheritance](https://gpcoder.com/wp-content/uploads/2018/10/design-patterns-class-adapter-diagram.png)

**So sánh Class Adapter với Object Adapter:**
- Sự khác biệt chính là Class Adapter sử dụng Inheritance (kế thừa) để kết nối Adapter và Adaptee trong khi Object Adapter sử dụng Composition (chứa trong) để kết nối Adapter và Adaptee.
- Trong cách tiếp cận Class Adapter, nếu một Adaptee là một class và không phải là một interface thì Adapter sẽ là một lớp con của Adaptee. Do đó, nó sẽ không phục vụ tất cả các lớp con khác theo cùng một cách vì Adapter là một lớp phụ cụ thể của Adaptee.

**Tại sao Object Adapter lại tốt hơn?**
- Nó sử dụng Composition để giữ một thể hiện của Adaptee, cho phép một Adapter hoạt động với nhiều Adaptee nếu cần thiết.

## 3. Ví dụ về Adapter Pattern với ứng dụng Translation
Một người Việt muốn trao đổi với một người Nhật. Tuy nhiên, 2 người này không biết ngôn ngữ của nhau nên cần phải có một người để chuyển đổi từ ngôn ngữ tiếng Việt sang ngôn ngữ tiếng Nhật. Chúng ta sẽ mô hình hóa trường hợp này với Adapter Pattern như sau:
- Client: người Việt sẽ là Client trong ví dụ này,vì anh ta cần gửi một số message cho người Nhật.
- Target: đây là nội dung message được Client cung cấp cho thông dịch viên (Translator / Adapter).
- Adapter: thông dịch viên (Translator) sẽ là Adapter, nhận message tiếng Việt từ Client và chuyển đổi nó sang tiếng Nhật trước khi gởi cho người Nhật.
- Adaptee: đây là interface hoặc class được người Nhật sử dụng để nhận message được chuyển đổi từ thông dịch viên (Translator).

![Translator Adaptor](https://gpcoder.com/wp-content/uploads/2018/10/design-patterns-adapter-diagram-translator-example.png)

## 4. Ưu điểm của Adapter Pattern
Việc sử dụng Adapter Pattern đem lại các lợi ích sau:
- Cho phép nhiều đối tượng có interface giao tiếp khác nhau có thể tương tác và giao tiếp với nhau.
- Tăng khả năng sử dụng lại thư viện với interface không thay đổi do không có mã nguồn.

## 5. Nhược điểm của Adapter Pattern
- Tất cả các yêu cầu được chuyển tiếp, do đó làm tăng thêm một ít chi phí.
- Đôi khi có quá nhiều Adapter được thiết kế trong một chuỗi Adapter (chain) trước khi đến được yêu cầu thực sự.

## 6. Sử dụng Adapter Pattern khi nào?
- Adapter Pattern giúp nhiều lớp có thể làm việc với nhau dễ dàng mà bình thường không thể. Một trường hợp thường gặp phải và có thể áp dụng Adapter Pattern là khi không thể kế thừa lớp A, nhưng muốn một lớp B có những xử lý tương tự như lớp A. Khi đó chúng ta có thể cài đặt B theo Object Adapter, các xử lý của B sẽ gọi những xử lý của A khi cần.
- Khi muốn sử dụng một lớp đã tồn tại trước đó nhưng interface sử dụng không phù hợp như mong muốn.
- Khi muốn tạo ra những lớp có khả năng sử dụng lại, chúng phối hợp với các lớp không liên quan hay những lớp không thể đoán trước được và những lớp này không có những interface tương thích.
- Cần phải có sự chuyển đổi interface từ nhiều nguồn khác nhau.
- Khi cần đảm bảo nguyên tắc Open/ Close trong một ứng dụng.

## Tham khảo
[1] https://gpcoder.com/4483-huong-dan-java-design-pattern-adapter/
