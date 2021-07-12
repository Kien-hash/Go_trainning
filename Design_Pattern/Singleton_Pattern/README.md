## Singleton Pattern
Đôi khi, trong quá trình phân tích thiết kế một hệ thống, chúng ta mong muốn có những đối tượng cần tồn tại duy nhất và có thể truy xuất mọi lúc mọi nơi. Làm thế nào để hiện thực được một đối tượng như thế khi xây dựng mã nguồn? Chúng ta có thể nghĩ tới việc sử dụng một biến toàn cục (global variable : public static final). Tuy nhiên, việc sử dụng biến toàn cục nó phá vỡ quy tắc của OOP (encapsulation). Để giải bài toán trên, người ta hướng đến một giải pháp là sử dụng Singleton pattern.
>![](https://gpcoder.com/wp-content/uploads/2018/09/singleton-pattern.png)\
"Singleton is a creational design pattern that lets you ensure that a class has only one instance and provide a global access point to this instance."

Singleton là design pattern trong nhóm khởi tạo. Mang trong mình chức năng đảm bảo việc tạo một thể hiện (instance) duy nhất trong chương trình và cung cấp method để có thể truy cập thể hiện đó ở mọi nơi trong chương trình.\
Sử dụng Singleton khi chúng ta muốn:
- Đảm bảo rằng chỉ có một instance của lớp.
- Việc quản lý việc truy cập tốt hơn vì chỉ có một thể hiện duy nhất.
- Có thể quản lý số lượng thể hiện của một lớp trong giớn hạn chỉ định.

Các nguyên tắc trong implement một Singleton 
- private constructor để hạn chế truy cập từ class bên ngoài.
- Đặt private static final variable đảm bảo biến chỉ được khởi tạo trong class.
- Có một method public static để return instance được khởi tạo ở trên.

## Về code
Trong folder singleton là những file code của nhiều các tạo nên các singleton khác nhau và file main là file thực thi việc kiểm tra tính duy nhất của các cách thực hiện đấy.

Những cách khác nhau tạo một singleton : 
- Eager initialization: Singleton Class được khởi tạo ngay khi được gọi đến. Đây là cách dễ nhất nhưng nó có một nhược điểm mặc dù instance đã được khởi tạo mà có thể sẽ không dùng tới
- Static block initialization: Cách làm tương tự như Eager initialization chỉ khác phần static block cung cấp thêm lựa chọn cho việc handle exception hay các xử lý khác.
- Lazy Initialization: Là một cách làm mang tính mở rộng hơn so với 2 cách làm trên và hoạt động tốt trong môi trường đơn luồng (single-thread).
- Thread Safe Singleton: nhược điểm là một phương thức synchronized sẽ chạy rất chậm và tốn hiệu năng, bất kỳ Thread nào gọi đến đều phải chờ nếu có một Thread khác đang sử dụng. Có những tác vụ xử lý trước và sau khi tạo thể hiện không cần thiết phải block.
- Double Check Locking Singleton: Để implement theo cách này, chúng ta sẽ kiểm tra sự tồn tại thể hiện của lớp, với sự hổ trợ của đồng bộ hóa, hai lần trước khi khởi tạo. Phải khai báo volatile cho instance để tránh lớp làm việc không chính xác do quá trình tối ưu hóa của trình biên dịch.
- Bill Pugh Singleton Implementation: Với cách làm này bạn sẽ tạo ra static nested class với vai trò 1 Helper khi muốn tách biệt chức năng cho 1 class function rõ ràng hơn. Đây là cách thường hay được sử dụng và có hiệu suất tốt 
- Enum Singleton: Khi dùng enum thì các params chỉ được khởi tạo 1 lần duy nhất, đây cũng là cách giúp bạn tạo ra Singleton instance.
- Serialization and Singleton: Đôi khi trong các hệ thống phân tán (distributed system), chúng ta cần implement interface Serializable trong lớp Singleton để chúng ta có thể lưu trữ trạng thái của nó trong file hệ thống và truy xuất lại nó sau

## Sử dụng Singleton Pattern khi nào?
Dưới đây là một số trường hợp sử dụng của Singleton Pattern thường gặp:
- Vì class dùng Singleton chỉ tồn tại 1 Instance (thể hiện) nên nó thường được dùng cho các trường hợp giải quyết các bài toán cần truy cập vào các ứng dụng như: Shared resource, Logger, Configuration, Caching, Thread pool, …
- Một số design pattern khác cũng sử dụng Singleton để triển khai: Abstract Factory, Builder, Prototype, Facade,…
- Đã được sử dụng trong một số class của core java như: java.lang.Runtime, java.awt.Desktop.

## Một số mẹo
- BillPughSingleton hay đươc dùng nhất vì có hiệu suất cao.
- LazyInitializedSingleton hay dùng cho những ứng dụng chỉ làm việc với single-thread 
- DoubleCheckLockingSingleton hay dùng khi làm việc với ứng dụng multi-thread

## Thảm khảo
[1]: <https://gpcoder.com/4190-huong-dan-java-design-pattern-singleton/>(Hướng dẫn Java Design Pattern - Singleton - GPCoder )
