public class Exercise2 {
    float f;

    public static void main(String[] args) {
        Exercise2 test1 = new Exercise2();
        Exercise2 test2 = new Exercise2();

        test1.f = 9.9f;
        test2.f = 10.1f;

        System.out.println("before aliasing, f1 = " + test1.f + ", f2 = " + test2.f);

        test1 = test2;
        System.out.println("after aliasing, f1 = " + test1.f + ", f2 = " + test2.f);
    }
}
