public class DataOnly {
    int i;
    double d;
    boolean b;

    public static void main(String[] args) {
        DataOnly data = new DataOnly();
        data.i = 10;
        data.d = 9.99;
        data.b = false;

        System.out.println("i: " + data.i + "\t d: " + data.d + "\t b: " + data.b);
    }
}