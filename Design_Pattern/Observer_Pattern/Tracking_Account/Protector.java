public class Protector implements Observer{

    @Override
    public void update(User user) {
        // TODO Auto-generated method stub
        if (user.getStatus() == LoginStatus.INVALID) {
            System.out.println("Protector: User " + user.getEmail() + " is invalid. "
                    + "IP " + user.getIp() + " is blocked");
        }
    }
    
}