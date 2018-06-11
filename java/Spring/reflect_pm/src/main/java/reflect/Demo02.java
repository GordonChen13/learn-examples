package reflect;

import java.util.Scanner;

public class Demo02 {
	public static void main(String[] args) 
		throws Exception{
		Scanner in = new Scanner(System.in);
		//动态加载类
		String className = in.nextLine();
		Class cls = Class.forName(className);
		//动态创建对象
		Object obj = cls.newInstance();
		System.out.println(obj);
		
	}
}







