/**
 * Cách làm tương tự như Eager initialization chỉ khác phần static block cung cấp thêm lựa chọn cho việc handle exception hay các xử lý khác.
 */

package singleton;

public class StaticBlockSingleton {
    
    private static final StaticBlockSingleton INSTANCE;
 
    private StaticBlockSingleton() {}
 
    // Static block initialization for exception handling
    static {
        try {
            INSTANCE = new StaticBlockSingleton();
        } catch (Exception e) {
            throw new RuntimeException("Exception occured in creating singleton instance");
        }
    }
 
    public static StaticBlockSingleton getInstance() {
        return INSTANCE;
    }
}