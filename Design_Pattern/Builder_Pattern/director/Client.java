package director;

import concretebuilder.*;
import product.type.*;
import product.Order;

public class Client {
    public static void main(String[] args) {
        Order order = new FastFoodOrderBuilder()
        .orderType(OrderType.ON_SITE)
        .orderBread(BreadType.OMELETTE)
        .orderSauce(SauceType.SOY_SAUCE)
        .build();

        
        System.out.println(order);
    }
}