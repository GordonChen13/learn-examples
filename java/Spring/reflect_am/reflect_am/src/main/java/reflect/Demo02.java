package reflect;

import java.util.Scanner;

public class Demo02 {
	public static void main(String[] args) 
		throws Exception{
		Scanner in = new Scanner(System.in);
		//��̬������
		String className = in.nextLine();
		Class cls = Class.forName(className);
		//��̬��������
		Object obj = cls.newInstance();
		System.out.println(obj);
		
	}
}







