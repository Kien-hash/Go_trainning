/**
 * Đây là hàm main, trong đây gồm 2 hàm. 
 * checkSingleton() : Hàm thực thi việc kiểm tra các instance là duy nhất 
 * reflectionBreakSingleton(): Hàm thực hiện việc phá vỡ cấu trúc của Singleton Pattern bằng reflect 
 */


import singleton.*;
import java.lang.reflect.Constructor;
import java.lang.reflect.InvocationTargetException;

public class Main {
    public static void main(final String[] args)
            throws InstantiationException, IllegalAccessException, InvocationTargetException {
        Main.checkSingleton();
    }

    public static void checkSingleton() {
        SerializedSingleton instanceOne = SerializedSingleton.getInstance();                
        SerializedSingleton instanceTwo = SerializedSingleton.getInstance();        
        if(instanceOne==instanceTwo){
            System.out.println("This is same instance");
        } else {
            System.out.println("This is 2 difference instance");
        }
        System.out.println(instanceOne.hashCode());
        System.out.println(instanceTwo.hashCode());
    }

    public static void reflectionBreakSingleton()throws InstantiationException, IllegalAccessException, InvocationTargetException{
        StaticBlockSingleton instanceOne = StaticBlockSingleton.getInstance();
        StaticBlockSingleton instanceTwo = StaticBlockSingleton.getInstance();
        
        // Sử dụng reflect để có thể tạo ra một đối tượng giống với đối tượng của class chứa Singleton
        Constructor<?>[] constructors = StaticBlockSingleton.class.getDeclaredConstructors();
        for (Constructor<?> constructor : constructors) {
            constructor.setAccessible(true);
            instanceTwo = (StaticBlockSingleton) constructor.newInstance();
        }
 
        System.out.println(instanceOne.hashCode());
        System.out.println(instanceTwo.hashCode());
    }
}