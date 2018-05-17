class Player {
    static int playerCount = 0;
    private String name;

    public Player(String name) {
        name = name;
        playerCount++;
    }
}