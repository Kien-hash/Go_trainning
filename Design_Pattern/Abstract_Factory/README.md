# Abstract Factory Design Pattern
>Abstract Factory is a creational design pattern that provide an interface for creating families of  related or dependent objects without specifying their concrete classes.

Abstract Factory pattern là một trong những Creational pattern. Nó là phương pháp tạo ra một Super-factory dùng để tạo ra các Factory khác. Hay còn được gọi là Factory của các Factory. Abstract Factory Pattern là một Pattern cấp cao hơn so với Factory Method Pattern.

Trong Abstract Factory pattern, một interface có nhiệm vụ tạo ra một Factory của các object có liên quan tới nhau mà không cần phải chỉ ra trực tiếp các class của object. Mỗi Factory được tạo ra có thể tạo ra các object bằng phương pháp giống như Factory pattern.

## Các thành phần cơ bản
Một Abstract Factory Pattern bao gồm các thành phần cơ bản sau:
![Sơ đồ cấu trúc Abstract Factory](https://i.imgur.com/xtQ9wDK.png)

- **AbstractFactory**: Khai báo dạng interface hoặc abstract class chứa các phương thức để tạo ra các đối tượng abstract.
- **ConcreteFactory**: Xây dựng, cài đặt các phương thức tạo các đối tượng cụ thể.
- **AbstractProduct**: Khai báo dạng interface hoặc abstract class để định nghĩa đối tượng abstract.
- **Product**: Cài đặt của các đối tượng cụ thể, cài đặt các phương thức được quy định tại AbstractProduct.
- **Client**: là đối tượng sử dụng AbstractFactory và các AbstractProduct.

## Ví dụ 
Một công ty đồ nội thất chuyên sản xuất ghế (Chair): ghế nhựa (PlasticChair) và ghế gỗ (WoodChair). Với tình hình kinh doanh ngày càng thuận thợi nên công ty quyết định mở rộng thêm sản xuất bàn (Table).\
Với lợi thế là đã có kinh nghiệm từ sản xuất ghế nên công ty vẫn giữ chất liệu là nhựa (PlasticTable) và gỗ (WoodTable) cho sản xuất bàn. Tuy nhiên, quy trình sản xuất ghế/ bàn theo từng chất liệu (MaterialType) là khác nhau. Nên công ty tách ra là nhà máy (Factory): 1 cho sản xuất vật liệu bằng nhựa (PlasticFactory), 1 cho sản xuất vật liệu bằng gỗ (WoodFactory), nhưng cả 2 đều có thể sản xuất ghế và bàn (FunitureAbstractFactory). Khi khách hàng cần mua một món đồ nào, khách hàng (Client) chỉ cần đến cửa hàng để mua (FunitureFactory). Khi đó ứng với từng hàng hóa và vật liệu sẽ được chuyển về phân xưởng tương ứng để sản xuất (createXXX) ra bàn (Table) và ghế (Chair).

Hệ thống được minh họa như sau:

![Factory](https://gpcoder.com/wp-content/uploads/2018/09/design-patterns-abstract-factory-diagram.png)

Trong đó:
- **AbstractFactory**: tương ứng với FurnitureAbstractFactory
- **ConcreteFactory**: tương ứng với FurnitureFactory, Plastic Factory & WoodFactory
- **Product**: bao gồm Plastic Chair, Plastic Table, Wood Chair, Wood Table.
- **Client**: tương ứng với client trên hình.
- **Các thành phần khác:** Các interface mô tả các phuong thức mà product sẽ có, MaterialType(Enum) đơn giản là một trong các tham số truyền vào factory để tạo đối tượng

Khi hệ thống phát triển cần mở rộng thêm 1 nhà máy khác, chẳng hạn sản xuất hàng hóa bằng inox, thì đơn giản cần tạo thêm một class mới implement từ FurnitureAbstractFactory, và thêm vào logic khởi tạo Funiture trong FurnitureFactory. Nó không làm ảnh hưởng đến code ở phía Client.

## Lợi ích của Abstract Factory Pattern
- Các lợi ích của Factory Pattern cũng tương tự như Factory Method Pattern như: cung cấp hướng tiếp cận với Interface thay thì các implement, che giấu sự phức tạp của việc khởi tạo các đối tượng với người dùng (client), độc lập giữa việc khởi tạo đối tượng và hệ thống sử dụng, …
- Giúp tránh được việc sử dụng điều kiện logic bên trong Factory Pattern. Khi một Factory Method lớn (có quá nhiều sử lý if-else hay switch-case), chúng ta nên sử dụng theo mô hình Abstract Factory để dễ quản lý hơn (cách phân chia có thể là gom nhóm các sub-class cùng loại vào một Factory).
- Abstract Factory Pattern là factory của các factory, có thể dễ dạng mở rộng để chứa thêm các factory và các sub-class khác.
- Dễ dàng xây dựng một hệ thống đóng gói (encapsulate): sử dụng được với nhiều nhóm đối tượng (factory) và tạo nhiều product khác nhau.

## Tham khảo
[1] https://gpcoder.com/4365-huong-dan-java-design-pattern-abstract-factory/

[2] https://vi.fitwp.com/abstract-factory-pattern/#:~:text=Abstract%20Factory%20Pattern%20bao%20g%E1%BB%93m,Product%2C%20Concrete%20Product%20v%C3%A0%20Client.
