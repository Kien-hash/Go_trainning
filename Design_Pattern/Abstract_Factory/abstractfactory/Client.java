package abstractfactory;

import abstractfactory.chair.Chair;
import abstractfactory.factory.FurnitureAbstractFactory;
import abstractfactory.factory.FurnitureFactory;
import abstractfactory.table.Table;



public class Client{
    public static void main (String[] args) {
 
            FurnitureAbstractFactory factory = FurnitureFactory.getFactory(MaterialType.FLASTIC);
     
            Chair chair = factory.createChair();
            chair.create(); // Create plastic chair
     
            Table table = factory.createTable();
            table.create(); // Create plastic table
        
    }
}