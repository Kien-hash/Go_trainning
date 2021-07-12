package abstractfactory.factory;

import abstractfactory.chair.*;
import abstractfactory.table.*;

public abstract class FurnitureAbstractFactory {
    public abstract Chair createChair();
 
    public abstract Table createTable();
}