public abstract class PizzaDecorator implements Pizza{

    protected Pizza pizza;

    protected PizzaDecorator(Pizza pizza){
        this.pizza = pizza;
    }

    @Override
    public String doPizza() {
        // TODO Auto-generated method stub
        return pizza.doPizza();
    }
    
}