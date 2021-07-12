public class Mailer implements Observer{

    @Override
    public void update(User user) {
        // TODO Auto-generated method stub
        if (user.getStatus() == LoginStatus.EXPIRED) {
            System.out.println("Mailer: User " + user.getEmail() + " is expired. An email was sent!");
        }
    }
    
}