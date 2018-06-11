package reflect;

import java.lang.reflect.Method;
import java.util.Scanner;

public class Demo04 {

	public static void main(String[] args)
		throws Exception{
		//动态加载类
		Scanner  in = new Scanner(System.in);
		String className = in.nextLine();
		//forName() 加载类的时候，如果多次执行
		// 方法加载类，实际上只加载一次
		Class cls = Class.forName(className);
		//动态获取类的方法信息
		Method[] methods = 
			cls.getDeclaredMethods();
		for(Method m:methods){
			System.out.println(m); 
		}
			
	}

}





