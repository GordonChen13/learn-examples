import java.io.IOException;
import java.io.PrintWriter;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class HelloWorldServlet extends HttpServlet {
    private static final long serialVersionUID = 1L;

    @Override
    protected void doGet(HttpServletRequest req, HttpServletResponse resp)
            throws ServletException, IOException {

        resp.setContentType("text/html;charset=gbk");


        PrintWriter out = resp.getWriter();


        String str = "<html><title></title><body>你好Hello World Servlet!!!</body></html>";

        out.println(str);
        out.flush();
        out.close();
        System.out.println("doget...");
    }


    @Override
    protected void doPost(HttpServletRequest req, HttpServletResponse resp)
            throws ServletException, IOException {
        PrintWriter out = resp.getWriter();
        String str = "<html><title></title><body>你好Hello World Servlet!!!</body></html>";
        out.println(str);
        out.flush();
        out.close();
        System.out.println("dopost...");
    }
}