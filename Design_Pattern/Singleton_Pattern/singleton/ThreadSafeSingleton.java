/**
 * Cách đơn giản nhất là chúng ta gọi phương thức synchronized của hàm getInstance() 
 * như vậy hệ thống đảm bảo rằng tại cùng một thời điểm chỉ có thể có 1 luồng có thể truy cập vào hàm getInstance() 
 * và đảm bảo rằng chỉ có duy nhất 1 thể hiện của class.
 * Cách này có nhược điểm là một phương thức synchronized sẽ chạy rất chậm và tốn hiệu năng, 
 * bất kỳ Thread nào gọi đến đều phải chờ nếu có một Thread khác đang sử dụng
 */

package singleton;

public class ThreadSafeSingleton{
    private static volatile ThreadSafeSingleton instance ;
 
    private ThreadSafeSingleton() {
    }
 
    public static synchronized ThreadSafeSingleton getInstance() {
        if (instance == null) {
            instance = new ThreadSafeSingleton();
        }
        return instance;
    }
}