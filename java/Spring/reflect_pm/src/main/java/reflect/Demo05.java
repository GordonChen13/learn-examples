package reflect;

import java.lang.reflect.Method;
import java.util.Scanner;

public class Demo05 {

	public static void main(String[] args) 
		throws Exception {
		//找到一个类中以test为开头方法
		Scanner in = new Scanner(System.in);
		String className = in.nextLine();
		Class cls = Class.forName(className);
		Method[] methods =
			cls.getDeclaredMethods();
		Object obj = cls.newInstance();
		for(Method m:methods){
			//getName() 获取方法信息中的方法名
			String name = m.getName();
			if(name.startsWith("test")){
				//m 是以test为开头方法
				System.out.println(m); 
				//num 代表一个方法的参数个数
				int num = 
					m.getParameterTypes().length;
				if(num==0){
					//执行以test为开头的方法
					Object val=m.invoke(obj);
					System.out.println(val); 
				}
			}
		}
	}

}


