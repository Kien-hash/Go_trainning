public class CloseAccount implements Command{

    private Account account;

    public CloseAccount(Account account){
        this.account = account;
    }

    @Override
    public void execute() {
        // TODO Auto-generated method stub
        account.close();
    }
    
}