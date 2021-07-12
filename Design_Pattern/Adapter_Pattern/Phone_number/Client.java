public class Client{
    // Do smt here
    public static void main(String[] args) {
        PhoneTarget valid = new Adapter(new CheckNumberAdaptee());
        valid.useCheckNum("013-2652314");
    }
}