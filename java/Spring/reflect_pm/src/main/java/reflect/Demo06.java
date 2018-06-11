package reflect;

import java.lang.reflect.Method;
import java.util.Scanner;

public class Demo06 {

	public static void main(String[] args)
		throws Exception {
		Scanner in = new Scanner(System.in);
		String className = in.nextLine();
		Class cls=Class.forName(className);
		//找到被执行的方法信息
		Method m = cls.getDeclaredMethod(
			"add", int.class, String.class);
		Object obj = cls.newInstance();
		//执行私有方法
		m.setAccessible(true);//破坏私有功能
		Object val = m.invoke(obj, 5, "Tom");
		System.out.println(val); 
	}

}




