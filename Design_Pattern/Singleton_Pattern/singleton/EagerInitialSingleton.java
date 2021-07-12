/**
 * Singleton Class được khởi tạo ngay khi được gọi đến. 
 * Đây là cách dễ nhất nhưng nó có một nhược điểm mặc dù instance đã được khởi tạo mà có thể sẽ không dùng tới.
 * Eager initialization là cách tiếp cận tốt, dễ cài đặt, tuy nhiên, nó dễ dàng bị phá vỡ bởi Reflection.
 */

package singleton;

public class EagerInitialSingleton {
    private static EagerInitialSingleton instance = null;

    private EagerInitialSingleton(){}

    public static EagerInitialSingleton getInstance(){
        return instance;
    }

}