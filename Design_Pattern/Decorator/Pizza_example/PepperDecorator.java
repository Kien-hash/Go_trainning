public class PepperDecorator extends PizzaDecorator{

    protected PepperDecorator(Pizza pizza) {
        super(pizza);
        // TODO Auto-generated constructor stub
    }
    
    public String addPepper(){
        return " + Pepper";
    }

    @Override
    public String doPizza(){
        String s = pizza.doPizza();
        return s + this.addPepper();
    }
}