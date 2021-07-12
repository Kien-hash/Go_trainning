public class CheeseDecorator extends PizzaDecorator{

    protected CheeseDecorator(Pizza pizza) {
        super(pizza);
        // TODO Auto-generated constructor stub
    }

    private String addCheese(){
        return " + Cheese";
    }
    
    @Override
    public String doPizza(){
        return pizza.doPizza() + this.addCheese();
    }
}