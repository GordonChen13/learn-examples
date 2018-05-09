import java.util.*;

public class DotComBust {
    private GameHelper helper = new GameHelper();
    private ArrayList<DotCom> dotComLists = new ArrayList<DotCom>();
    private int numOfGuess = 0;

    private void setUpGame() {
        DotCom one = new DotCom();
        one.setName("Pets.com");
        DotCom two = new DotCom();
        two.setName("eToys.com");
        DotCom three = new DotCom();
        three.setName("Go2.com");

        dotComLists.add(one);
        dotComLists.add(two);
        dotComLists.add(three);

        System.out.println("Your goal is to sink three dot coms");
        System.out.println("Pets.com, eToys.com, Go2.com");
        System.out.println("Try to sink them all in the fewest number of guesses");

        for (DotCom dotComToSet : dotComLists) {
            ArrayList<String> newLocation = helper.placeDotCom(3);
            dotComToSet.setLocationCells(newLocation);
        }
    }

    private void startPlaying() {
        while (!dotComList.isEmpty()) {
            String userGuess = helper.getUserInput("Enter a guess");
            checkUserGuess(userGuess);
        }
        finishGame();
    }

    private void checkUserGuess(String userGuess) {
        numOfGuess++;
        String result = "miss";

        for (DotCom dotComToTest : dotComLists) {
            result = dotComToTest.checkYourself(userGuess);
            if (result.equals("hit")) {
                break;
            }
            if (result.equals("kill")) {
                dotComLists.remove(dotComToTest);
                break;
            }
        }
        System.out.println(result);
    }

    private void finishGame() {
        System.out.println("All dot coms are dead! Your stock is now worthless.");
        if (numOfGuess <= 18) {
            System.out.println("It only took you " + numOfGuess + "guesses.");
            System.out.println("You got out before your opotions sank.");
        } else {
            System.out.println("Took you long enough. " + numOfGuess + "guesses.");
            System.out.println("Fish are dance with your options");
        }
    }

    public static void main(String[] args) {
        DotComBust game = new DotComBust();
        game.setUpGame();
        game.startPlaying();
    }
}