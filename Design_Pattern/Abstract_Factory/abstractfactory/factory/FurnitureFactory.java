package abstractfactory.factory;

import abstractfactory.MaterialType;
import abstractfactory.factory.impl.FlasticFactory;
import abstractfactory.factory.impl.WoodFactory;


public class FurnitureFactory {
    private FurnitureFactory(){}

    public static FurnitureAbstractFactory getFactory(MaterialType materialType){
        switch (materialType) {
            case FLASTIC:
                return new FlasticFactory();
            case WOOD:
                return new WoodFactory();
            default:
                throw new UnsupportedOperationException("This furniture is unsupported ");
        }
    }
}