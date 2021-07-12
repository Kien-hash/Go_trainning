public class OpenAccount implements Command{

    private Account account;

    public OpenAccount( Account account){
        this.account = account;
    }

    @Override
    public void execute() {
        // TODO Auto-generated method stub
        account.open();
    }
    
}