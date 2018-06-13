public class PrimitiveWrapperTest {
    Boolean aBoolean = false;
    Character ch = 'c';
    Byte by = 13;
    Short sh = 1024;
    Integer in = 9999;
    Long lo = 99999999l;
    Float fl = 123.56f;
    Double dou = 789.123456d;

    public static void main(String[] args) {
        PrimitiveWrapperTest test = new PrimitiveWrapperTest();
        System.out.println("Boolean: " + test.aBoolean);
        System.out.println("Character: " + test.ch);
        System.out.println("Byte: " + test.by);
        System.out.println("Short: " + test.sh);
        System.out.println("Integer: " + test.in);
        System.out.println("Long: " + test.lo);
        System.out.println("Float: " + test.fl);
        System.out.println("Double: " + test.dou);
    }
}