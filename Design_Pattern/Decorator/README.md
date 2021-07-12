# Decorator Pattern
## 1. Decorator Pattern là gì?
> Attach additional responsibilities to an object dynamically. Decorators provide a flexible alternative to subclassing for extending functionality.

Decorator pattern là một trong những Pattern thuộc nhóm cấu trúc (Structural Pattern). Nó cho phép người dùng thêm chức năng mới vào đối tượng hiện tại mà không muốn ảnh hưởng đến các đối tượng khác. Kiểu thiết kế này có cấu trúc hoạt động như một lớp bao bọc (wrap) cho lớp hiện có. Mỗi khi cần thêm tính năng mới, đối tượng hiện có được wrap trong một đối tượng mới (decorator class).

Decorator pattern sử dụng composition thay vì inheritance (thừa kế) để mở rộng đối tượng. Decorator pattern còn được gọi là Wrapper hay Smart Proxy.

## 2. Các thành phần trong Decorator
Decorator pattern hoạt động dựa trên một đối tượng đặc biệt, được gọi là decorator (hay wrapper). Nó có cùng một interface như một đối tượng mà nó cần bao bọc (wrap), vì vậy phía client sẽ không nhận thấy khi bạn đưa cho nó một wrapper thay vì đối tượng gốc.

Tất cả các wrapper có một trường để lưu trữ một giá trị của một đối tượng gốc. Hầu hết các wrapper khởi tạo trường đó với một đối tượng được truyền vào constructor của chúng.

Vậy làm thế nào để có thể thay đổi hành vi của đối tượng? Như đã đề cập, wrapper có cùng interface với các đối tượng đích. Khi bạn gọi một phương thức decorator, nó thực hiện cùng một phương thức trong một đối tượng được wrap và sau đó thêm một cái gì đó (tính năng mới) vào kết quả, công việc này tùy thuộc vào logic nghiệp vụ.

![](https://gpcoder.com/wp-content/uploads/2018/11/design-patterns-decorator-diagram.png)

Các thành phần trong mẫu thiết kế Decorator:
- Component: là một interface quy định các method chung cần phải có cho tất cả các thành phần tham gia vào mẫu này.
- ConcreteComponent : là lớp hiện thực (implements) các phương thức của Component.
- Decorator : là một abstract class dùng để duy trì một tham chiếu của đối tượng Component và đồng thời cài đặt các phương thức của Component  interface.
- ConcreteDecorator : là lớp hiện thực (implements) các phương thức của Decorator, nó cài đặt thêm các tính năng mới cho Component.
- Client : đối tượng sử dụng Component.

## 3. Ví dụ đối với hệ thống quản lý nhân viên
Để đơn giản hơn, chúng ta xem ví dụ về một hệ thống quản lý dự án, nơi nhân viên đang làm việc với các vai trò khác nhau, chẳng hạn như thành viên nhóm (team member), trưởng nhóm (team lead) và người quản lý (manager). Một thành viên trong nhóm chịu trách nhiệm thực hiện các nhiệm vụ được giao và phối hợp với các thành viên khác để hoàn thành nhiệm vụ nhóm. Mặt khác, một trưởng nhóm phải quản lý và cộng tác với các thành viên trong nhóm của mình và lập kế hoạch nhiệm vụ của họ. Tương tự như vậy, một người quản lý có thêm một số trách nhiệm đối với một trưởng nhóm như quản lý yêu cầu dự án, tiến độ, phân công công việc.

Sau đây là các thành phần tham gia vào hệ thống và hành vi của chúng:
- Employee: thực hiện công việc (doTask), tham gia vào dự án (join), rời khỏi dự án (terminate).
- Team member: báo cáo task được giao (report task), cộng tác với các thành viên khác (coordinate with others).
- Team lead: lên kế hoạch (planning), hỗ trợ các thành viên phát triển (motivate), theo dõi chất lượng công việc và thời gian (monitor).
- Manager: tạo các yêu cầu dự án (create requirement), giao nhiệm vụ cho thành viên (assign task), quản lý tiến độ dự án (progress management).

Với cách làm thông thường, chúng ta có sơ đồ như sau:
![Sơ đồ quản lý nhân viên với cách làm thông thường](https://gpcoder.com/wp-content/uploads/2018/11/design-patterns-decorator-example1.png)

Bất cứ khi nào một thành viên trong nhóm trở thành một Team Lead, chúng ta phải tạo một đối tượng mới của Team Lead và đối tượng trước đó tham chiếu vào nhân viên đó (Team Member trong nhóm) có thể bị hủy hoặc lưu trữ. Đó không phải là cách tiếp cận được khuyến nghị khi nhân viên vẫn là một phần của tổ chức của bạn. Tương tự như trường hợp với Manager, khi một nhân viên trở thành người quản lý từ một Team Lead / Team Member.

Một trường hợp khác là khi một nhân viên có thể thực hiện trách nhiệm của một Team Member trong nhóm cũng như trách nhiệm của Team Lead hoặc một Manager. Trong trường hợp đó, bạn cần tạo hai đối tượng cho cùng một nhân viên là hoàn toàn sai.

Trong các kịch bản này, một Team Member/ Team Lead có thể có thêm trách nhiệm trong lúc thực hiện (run-time). Và trách nhiệm của họ có thể được chỉ định / thu hồi trong lúc run-time.

Hãy xem sơ đồ bên dưới để thấy được cách chúng ta áp dụng Decorator Pattern như thế nào trong trường hợp này.

![Sơ đồ quản lý nhân viên với Decorator Pattern](https://gpcoder.com/wp-content/uploads/2018/11/design-patterns-decorator-example2.png)

Như ta thấy, với Decorator hệ thống của chúng ta linh hoạt hơn rất nhiều. Chúng ta có thể dễ dàng gán một nhân viên sang vai trò TeamMember, TeamLeader, Manager.

## 4. Ưu điểm của Decorator Pattern
- Tăng cường khả năng mở rộng của đối tượng, bởi vì những thay đổi được thực hiện bằng cách implement trên các lớp mới.
- Client sẽ không nhận thấy sự khác biệt khi bạn đưa cho nó một wrapper thay vì đối tượng gốc.
- Một đối tượng có thể được bao bọc bởi nhiều wrapper cùng một lúc.
- Cho phép thêm hoặc xóa tính năng của một đối tượng lúc thực thi (run-time).

## 5. Sử dụng Decorator Pattern khi nào?
- Khi muốn thêm tính năng mới cho các đối tượng mà không ảnh hưởng đến các đối tượng này.
- Khi không thể mở rộng một đối tượng bằng cách thừa kế (inheritance). Chẳng hạn, một class sử dụng từ khóa final, muốn mở rộng class này chỉ còn cách duy nhất là sử dụng decorator.
- Trong một số nhiều trường hợp mà việc sử dụng kế thừa sẽ mất nhiều công sức trong việc viết code. Ví dụ trên là một trong những trường hợp như vậy.

## 6. So sánh Decorator và Adapter
Giống nhau:
- Cả hai đều là structural pattern như định nghĩa của GOF.
- Cả hai sử dụng cách composition để cài đặt.

Khác nhau:
- Decorator cho phép thêm một tính năng mới vào một object nhưng không được phép sử dụng thừa kế. Nó cho phép thay đổi lúc thực thi (run-time). Adapter được sử dụng khi bạn có một interface, và bạn muốn ánh xạ interface đó đến một đối tượng khác có vai trò chức năng tương tự, nhưng là một interface khác.
- Decorator có xu hướng hoạt động trên một đối tượng. Adapter có xu hướng hoạt động trên nhiều đối tượng (có thể wrap nhiều interface).

## Tham khảo
[1] https://gpcoder.com/4574-huong-dan-java-design-pattern-decorator/
