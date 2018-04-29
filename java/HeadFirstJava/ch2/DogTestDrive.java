class DogTestDrive {
    public static void main(String[] args) {

        Dog d = new Dog();
//        Dog d2 = new Dog();
        Dog d2 = d;

        d.size = 40;
        d.bark();

        d2.size = 50;

        System.out.println("new size of d is: " + d.size);
    }
}