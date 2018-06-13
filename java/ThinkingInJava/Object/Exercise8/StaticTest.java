public class StaticTest {
    static int i = 10;

    public static void main(String[] args) {
        for (int a = 0; a < 10; a ++) {
            StaticTest test = new StaticTest();
            System.out.println("i = " + test.i);
        }
    }
}
