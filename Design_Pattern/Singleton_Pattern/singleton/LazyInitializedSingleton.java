/**
 * Khi hàm getInstance() được gọi thì đối tượng mới được khởi tạo
 * Là một cách làm mang tính mở rộng hơn so với 2 cách làm trên và hoạt động tốt trong môi trường đơn luồng (single-thread).
 * Trường hợp nếu có nhiều luồng (multi-thread) cùng chạy và cùng gọi hàm getInstance() tại cùng một thời điểm thì có thể có nhiều hơn 1 thể hiện của instance.
 * Đối với thao tác create instance quá chậm thì người dùng có phải chờ lâu cho lần sử dụng đầu tiên
 */

package singleton;

public class LazyInitializedSingleton {
    private static LazyInitializedSingleton instance;
 
    private LazyInitializedSingleton() {
    }
 
    public static LazyInitializedSingleton getInstance() {
        if (instance == null) {
            instance = new LazyInitializedSingleton();
        }
        return instance;
    }
}