public class Logger implements Observer{
    
    @Override
    public void update(User user) {
        // TODO Auto-generated method stub
        System.out.println("Logger: " + user + "(" +
        user.getEmail() + ", " + user.getIp() + ", " + user.getStatus()+ ")");
    }
}