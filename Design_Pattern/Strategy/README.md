# Strategy
## 1. Strategy Pattern 
>Define a family of algorithms, encapsulate each one, and make them interchangeable. Strategy lets the algorithm vary independently from the clients that use it.

Strategy Pattern là một trong những Pattern thuộc nhóm hành vi (Behavior Pattern). Nó cho phép định nghĩa tập hợp các thuật toán, đóng gói từng thuật toán lại, và dễ dàng thay đổi linh hoạt các thuật toán bên trong object. Strategy cho phép thuật toán biến đổi độc lập khi người dùng sử dụng chúng.

Ý nghĩa thực sự của Strategy Pattern là giúp tách rời phần xử lý một chức năng cụ thể ra khỏi đối tượng. Sau đó tạo ra một tập hợp các thuật toán để xử lý chức năng đó và lựa chọn thuật toán nào mà chúng ta thấy đúng đắn nhất khi thực thi chương trình. Mẫu thiết kế này thường được sử dụng để thay thế cho sự kế thừa, khi muốn chấm dứt việc theo dõi và chỉnh sửa một chức năng qua nhiều lớp con.

### **Bài toán thực tế.** 
Bạn nhận được một hợp đồng thiết kế ô tô. Có rất nhiều mẫu ô tô để bạn có thể làm. Bạn nghĩ ngay đến việc sử dụng OOP vào trong thiết kế ô tô của mình. Đầu tiên, bạn tạo ra 1 lớp cơ sở có tên là Vehicle với một phương thức có tên là go, phương thức này xuất hiện lên dòng chữ Now I’m driving.

~~~
abstract class Vehicle
{
    public function go()
    {
        echo "Now I'm driving";
    }
}
~~~
Sau đó, bạn tạo tiếp 1 lớp mới là lớp StreetRacer thừa kế từ lớp Vehicle như sau:
~~~
public class StreetRacer extends Vehicle
{

}
~~~
Tới đây, chương trình của bạn vẫn hoàn toàn tốt đẹp. Bạn có thể khai báo 1 đối tượng StreetRacer và gọi tới hàm go:
~~~
$streetRacer = new StreetRacer;
$streetRacer->go();
~~~
Và kết quả trả về là: ***Now I’m driving***. Kết quả hoàn toàn chính xác. Nhưng sau đó, bạn nhận thêm 1 hợp đồng sản xuất máy bay trực thăng Helicopter. Bạn nhận thấy máy bay trực thăng thì cũng là 1 phương tiện vận chuyển. Vì vậy bạn quyết định tạo ra 1 lớp Helicopter thừa kế từ lớp Vehicle:
~~~
public class Helicopter extends Vehicle
{

}
~~~
Nhưng bạn chợt nhận ra vấn đề là khi sử dụng hàm go cho Helicopter, thì kết quả trả về có vẻ không chính xác. ***Now I’m driving?*** Tại sao máy bay lại là ***driving***? Máy bay thì phải bay chứ nhỉ? Nhưng nó không phải là vấn đề lớn. Bạn quyết định sẽ override hàm go cho lớp Helicopter như sau:
~~~
public class Helicopter extends Vehicle
{
    public function go()
    {
        echo "Now I'm flying";
    }
}
~~~
Có vẻ vấn đề đã được giải quyết. Giờ thì máy bay đã là ***flying*** rồi. Nhưng vài tuần sau, khách hàng yêu cầu phải chuyển từ ***Now, I’m flying*** sang ***Now, I’m flying 200mph*** và nhiều sự thay đổi kế tiếp. Có một vấn đề nảy sinh ở đây. Đó chưa phải là một vấn đề lớn, nhưng nếu bạn phải xử lý các công việc này một cách khá thường xuyên, thì việc cứ phải chỉnh sửa các lớp con như thế này trờ thành 1 vấn đề bảo trì khá nghiêm trọng.

Bạn bắt đầu suy nghĩ. Có lẽ sự thừa kế không phải là cách giải quyết tốt cho tình huống này. Nơi mà bạn cần phải thay đổi các chức năng thường xuyên ở các lớp con. Và khi có càng nhiều lớp kế thừa liên quan, chúng cũng cần được bảo trì khi có sự thay đổi và khi đó, bạn sẽ phải cập nhất phương thức go nhiều lần.

Vấn đề bạn cần phải giải quyết ở đây là làm sao để tránh được việc thay đổi ở các lớp con, nếu không, bạn sẽ phải thay đổi code ở rất nhiều file để cập nhật được yêu cầu của khách hàng.

Có lẽ là bạn cần một cách khác tốt hơn để xử lý vấn đề này thay vì sử dụng thừa kế. Ở đây ta dùng Strategy Pattern.

## 2. Các thành phần trong Strategy Pattern
![Các thành phần trong Strategy Pattern](https://gpcoder.com/wp-content/uploads/2019/01/design-patterns-strategy-diagram.png)

Các thành phần tham gia Strategy Pattern:
- **Strategy** : định nghĩa các hành vi có thể có của một Strategy.
- **ConcreteStrategy** : cài đặt các hành vi cụ thể của Strategy.
- **Context** : chứa một tham chiếu đến đối tượng Strategy và nhận các yêu cầu từ Client, các yêu cầu này sau đó được ủy quyền cho Strategy thực hiện.

## 3. Ví dụ Strategy Pattern với ứng dụng sắp xếp
Chương trình này cung cấp 3 thuật toán sắp xếp: quick sort, merge sort, selection sort, heap sort, tim sort, …. Tùy thuộc vào loại dữ liệu, số lượng phần tử,... để client có thể thực hiện lựa chọn sao cho hợp lý.
![Ứng dụng sắp xếp](https://gpcoder.com/wp-content/uploads/2019/01/design-patterns-strategy-example.png)

Trong đó: 
- SortStrategy: Đóng vai trò là Strategy
- QuickSort, MergeSort, SelectionSort: Đóng vai trò là các ConcreteStrategy chứa đựng việc triển khai các thuật toán cụ thể
- SortedList: Chứa yêu cầu được gửi đến cho Strategy để thực hiện xử lý, ở đây là các items cần đươc sắp xếp.

## 4. Lợi ích của Strategy Pattern
- Đảm bảo nguyên tắc Single responsibility principle (SRP) : một lớp định nghĩa nhiều hành vi và chúng xuất hiện dưới dạng với nhiều câu lệnh có điều kiện. Thay vì nhiều điều kiện, chúng ta sẽ chuyển các nhánh có điều kiện liên quan vào lớp Strategy riêng lẻ của nó.
- Đảm bảo nguyên tắc Open/Closed Principle (OCP) : chúng ta dễ dàng mở rộng và kết hợp hành vi mới mà không thay đổi ứng dụng.
- Cung cấp một sự thay thế cho kế thừa.
  
## 5. Sử dụng Strategy Pattern khi nào?
- Khi muốn có thể thay đổi các thuật toán được sử dụng bên trong một đối tượng tại thời điểm run-time.
- Khi có một đoạn mã dễ thay đổi, và muốn tách chúng ra khỏi chương trình chính để dễ dàng bảo trì.
- Tránh sự rắc rối, khi phải hiện thực một chức năng nào đó qua quá nhiều lớp con.
- Cần che dấu sự phức tạp, cấu trúc bên trong của thuật toán.

## Tài liệu tham khảo
[1] https://gpcoder.com/4796-huong-dan-java-design-pattern-strategy/

[2] https://viblo.asia/p/tim-hieu-strategy-pattern-znmMdy7YGr69