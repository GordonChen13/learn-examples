public class Duck {
    int size;

    public Duck() {
        size = 20;

        System.out.println("the default size is " + size);
    }

    public Duck(int duckSize) {
        System.out.println("Quack");

        size = duckSize;

        System.out.println("size is " + size);
    }
}
