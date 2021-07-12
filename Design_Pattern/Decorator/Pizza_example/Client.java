public class Client {
    public static void main(String[] args) {
        Pizza chicken = new ChickenPizza();
        System.out.println(chicken.doPizza());

        PepperDecorator pepper = new PepperDecorator(chicken);
        System.out.println(pepper.doPizza());

        CheeseDecorator cheese = new CheeseDecorator(chicken);
        System.out.println(cheese.doPizza());

        PepperDecorator all = new PepperDecorator(cheese);
        System.out.println(all.doPizza());
    }
}