package abstractfactory.factory.impl;

import abstractfactory.chair.*;
import abstractfactory.table.*;
import abstractfactory.factory.FurnitureAbstractFactory;

public class FlasticFactory extends FurnitureAbstractFactory{
    @Override
    public Chair createChair() {
        return new PlasticChair();
    }
 
    @Override
    public Table createTable() {
        return new PlasticTable();
    }
}