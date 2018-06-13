//: object/HelloDate.java
import java.util.*;

/** The first Thinking in Java example program
 * Dispalys a string adn today's data.
 * @author Gordon
 * @author gordonchen13@test.com
 * @version 4.0
 */
public class HelloDate {
    /** Entry point to class and applicaton.
     * @param args array of string arguments
     */
    public static void main(String[] args) {
        System.out.println("Hello, it's: ");
        System.out.println(new Date());
    }
}
/* Output: (55% match)
Hello, it's:
Wed Oct 05 14:39:36 MDT 2005
 *///:~