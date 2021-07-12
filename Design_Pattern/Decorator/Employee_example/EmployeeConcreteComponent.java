import java.util.Date;

public class EmployeeConcreteComponent implements EmployeeComponent {

    private String name;

    public EmployeeConcreteComponent(String name){
        this.name = name;
    }

    @Override
    public String getName() {
        // TODO Auto-generated method stub
        return this.name;
    }

    @Override
    public void doTask() {
        // TODO Auto-generated method stub

    }

    @Override
    public void join(Date joinDate) {
        // TODO Auto-generated method stub
        System.out.println(this.getName()+" joined on "+ formatDate(joinDate));
    }

    @Override
    public void terminate(Date terminateDate) {
        // TODO Auto-generated method stub
        System.out.println(this.getName()+" terminated on "+ formatDate(terminateDate));
    }
    
}