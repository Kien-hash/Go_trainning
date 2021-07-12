public class ShopClient {
    public static void main(String[] args) {
        ShopFacade.getInstance().buyProductByCashWithFreeShipping("email@example.com");
        ShopFacade.getInstance().buyProductByPaypalWithStandardShipping("email@example.com", "0121546216");
    }
}