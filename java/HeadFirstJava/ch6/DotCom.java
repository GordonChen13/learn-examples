import java.util.*;

public class DotCom {
    private String name;
    private ArrayList<String> locationCells;

    public void setLocationCells(ArrayList<String> cells) {
        locationCells = cells;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String checkYourself(String userGuess) {
        String result = "miss";
        int index = locationCells.indexOf(userGuess);

        if (index >= 0) {
            locationCells.remove(userGuess);

            if (locationCells.isEmpty()) {
                result = "kill";
                System.out.println("Ouch! You sunk " + name + " : ( ");
            } else {
                result = "hit";
            }
        }

        return result;
    }
}