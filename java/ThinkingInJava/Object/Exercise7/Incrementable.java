public class Incrementable {
    public static void main(String[] args) {
        increment();
    }

    public static void increment() {
        System.out.println("before increment, i = " + StaticTest.i);
        StaticTest.i++;
        System.out.println("after increment, i = " + StaticTest.i);
    }
}