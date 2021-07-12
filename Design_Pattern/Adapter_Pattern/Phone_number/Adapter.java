public class Adapter implements PhoneTarget{

    private CheckNumberAdaptee checkNum;

    public Adapter (CheckNumberAdaptee check1){
        this.checkNum = check1;
    }

    @Override
    public void useCheckNum(String phoneNum) {
        // TODO Auto-generated method stub
        System.out.println(this.checkNum.check(phoneNum));
    }
    
}