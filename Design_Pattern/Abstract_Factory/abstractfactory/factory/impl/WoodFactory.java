package abstractfactory.factory.impl;

import abstractfactory.chair.*;
import abstractfactory.table.*;
import abstractfactory.factory.FurnitureAbstractFactory;

public class WoodFactory extends FurnitureAbstractFactory{
    @Override
    public Chair createChair() {
        return new WoodChair();
    }
 
    @Override
    public Table createTable() {
        return new WoodTable();
    }
}