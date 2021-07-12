import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class CheckNumberAdaptee {
    public boolean check(String num){
        Pattern pattern = Pattern.compile("\\d{3}-\\d{7}");
        Matcher matcher = pattern.matcher(num); 
        if(matcher.matches()){
            return true;
        } else return false;
    }
}