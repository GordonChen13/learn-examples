import java.awt.*;
import javax.swing.*;

public class MyDrawPanel extends JPanel {

    public void paintComponent(Graphics g) {
        Graphics2D g2d = (Graphics2D) g;

        int red = (int) (Math.random() * 255);
        int green = (int) (Math.random() * 255);
        int blue = (int) (Math.random() * 255);
        Color startColor = new Color(red, green, blue);

        int red2 = (int) (Math.random() * 255);
        int green2 = (int) (Math.random() * 255);
        int blue2 = (int) (Math.random() * 255);
        Color endColor= new Color(red2, green2, blue2);

        GradientPaint gradient = new GradientPaint(70, 70, startColor, 150, 150, endColor);
        g2d.setPaint(gradient);
        g2d.fillOval(70, 70, 100, 100);

    }
}