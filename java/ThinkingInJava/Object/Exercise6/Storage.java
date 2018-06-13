public class Storage {
    public static void main(String[] args) {
        int length = storage("how are you?");
        System.out.println("'how are you?' length times 2 is: " + length);
    }

    public static int storage(String s) {
        return s.length() * 2;
    }
}