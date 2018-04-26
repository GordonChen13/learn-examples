public class PhraseOmatic {
    public static void main(String[] args) {

        String[] workListOne = {"man", "student", "women", "kids", "boys", "girls",
        "teacher", "archer", "police", "writer", "coder", "programemr", "boss"};

        String[] workListTwo = {"work", "study", "learn", "run", "code", "play", "shopping",
        "slepping", "coding", "kidding"};

        String[] workListThree = {"hard", "today", "tommorrow", "yesterday", "weekend", "sunday",
        "holliday", "now", "a moment ago"};

        int oneLength = workListOne.length;
        int twoLength = workListTwo.length;
        int threeLength = workListThree.length;

        int rand1 = (int) (Math.random() * oneLength);
        int rand2 = (int) (Math.random() * twoLength);
        int rand3 = (int) (Math.random() * threeLength);

        String phrase = workListOne[rand1] + " " + workListTwo[rand2] + " " + workListThree[rand3];

        System.out.println("What I know is " + phrase);
    }
}