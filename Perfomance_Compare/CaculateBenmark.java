import java.time.Duration;
import java.time.Instant;
import java.util.Random;

public class CaculateBenmark {
    static final double N1 = 500000;
    static final double N2 = 1000000;
    static final double N3 = 1000000000;
    static final double N4 = 1e10;
    static final double N5 = 1e11;
    static final String charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
    static int charset_Size = charset.length();

    public static void main(final String[] args) {

        Instant start = Instant.now();

        {
            // Do smt to comsume time
        }

        Instant finish = Instant.now();
        long timeElapsed = Duration.between(start, finish).toMillis();
        System.out.println("Time consumming: " + timeElapsed / 1000.0);
    }

    /*********************************************************************
     * Ham tinh tong tu 1 den N
     *********************************************************************/
    public static double sumBenmark(final double N) {
        double sum = 0;
        double count = 1;
        do {
            sum += count;
            // count += 1;
        } while (count++ < N);
        return sum;
    }

    /*********************************************************************
     * Ham tinh tong cua chuoi phan so
     *********************************************************************/
    public static double cal_benmark(final double N) {
        double sum = 0;
        double count = 1;
        double num;
        int sign = -1;
        do {
            // printf("%d\n",sign);
            num = count / ((count + 1) * (count + 2));
            sum += sign * num;
            sign = -sign;
            // printf("\n%f",sum);
        } while (count++ < N);
        return sum;
    }

    /*********************************************************************
     * Ham tinh so thu N trong day Fibonaci
     *********************************************************************/
    public static double Fibonaci(final int sequence) {
        if (sequence == 1 || sequence == 2) {
            return 1;
        } else {
            return Fibonaci(sequence - 1) + Fibonaci(sequence - 2);
        }
    }

    /*********************************************************************
     * Ham sinh ra chuoi co N ki tu ngau nhien
     *********************************************************************/
    public static String randomString(int length) {
        if (length < 1)
            throw new IllegalArgumentException();

        StringBuilder sb = new StringBuilder(length);
        for (int i = 0; i < length; i++) {
            int rndCharAt = random.nextInt(charset_Size);
            char rndChar = charset.charAt(rndCharAt);
            sb.append(rndChar);
        }
        return sb.toString();
    }

    /*********************************************************************
     * Ham so sanh 2 chuoi co cung do dai la N
     *********************************************************************/
    public static int strCompare(String str1, String str2) {
        int count = 0;
        for (int i = 0; i < str1.length(); i++) {
            if (str1.charAt(i) == str2.charAt(i))
                count++;
        }
        return count;
    }
}
