public class SimpleDotCom {
    int[] localtionCells;
    int numOfHits = 0;

    public void setLocaltionCells(int[] cells) {
        this.localtionCells = cells;
    }

    public static checkYourself(String stringGuess) {
        int guess = Integer.parseInt(stringGuess);
        String result = "miss";

        for (int cell : localtionCells) {
            if (guess == cell) {
                result = "hit";
                numOfHits++ ;
                break;
            }
        }

        if (numOfHits == localtionCells.length) {
            result = "kill";
        }

        System.out.println(result);
        return result;
    }
}