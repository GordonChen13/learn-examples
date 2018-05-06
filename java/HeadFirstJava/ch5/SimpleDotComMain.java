//package com.gordonchen13.headfirst;

public class SimpleDotComMain {
    public static void main(String[] args) {
        int numofGuess = 0;

        GameHelper helper = new GameHelper();

        SimpleDotCom dotCom = new SimpleDotCom();

        int randomNum = (int) (Math.random() * 5);

        int[] locations = {randomNum, randomNum + 1, randomNum + 2};

        dotCom.setLocationCells(locations);

        boolean isAlive = true;

        while (isAlive == true) {
            String guess = helper.getUserInput("enter a number");
            String result = dotCom.checkYourself(guess);

            numofGuess++;

            if (result.equals("kill")) {
                isAlive = false;
                System.out.println("You took " + numofGuess + " guesses.");
            }
        }
    }
}

