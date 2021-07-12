import Bank.Bank;
import Bank.TPBank;
import Bank.VietcomBank;

public class BankFactory {
    private BankFactory(){}

    public static final Bank getBank(BankType bankType){
        switch (bankType){
            case TPBANK:
                return new TPBank();
            case VIETCOMBANK:
                return new VietcomBank();
            default:
                throw new IllegalArgumentException("This bank is not signed yet!");
        }
        
    }
}