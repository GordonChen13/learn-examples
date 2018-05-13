public class MakeCanine {
    public void go() {
        Canine c;
        c = new Dog();

        // class Canine is marked abstract, so the compiler will NOT let you do this.
        c = new Canine();
        c.roam();
    }
}
