import java.util.Date;

public class EmployeeDecorator implements EmployeeComponent {

    protected EmployeeComponent employee;

    protected EmployeeDecorator(EmployeeComponent employee){
        this.employee = employee;
    }

    @Override
    public String getName() {
        // TODO Auto-generated method stub
        return employee.getName();
    }

    @Override
    public void doTask() {
        // TODO Auto-generated method stub
        
    }

    @Override
    public void join(Date joinDate) {
        // TODO Auto-generated method stub
        employee.join(joinDate);
    }

    @Override
    public void terminate(Date terminateDate) {
        // TODO Auto-generated method stub
        employee.terminate(terminateDate);
    }
    
}