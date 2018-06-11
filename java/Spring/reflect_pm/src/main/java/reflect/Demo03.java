package reflect;

import java.lang.reflect.Field;
import java.util.Scanner;

public class Demo03 {

	public static void main(String[] args)
		throws Exception{
		//显示 某个类 的属性信息
		Scanner in = new Scanner(System.in);
		String className = in.nextLine();
		//动态加载类
		Class cls = Class.forName(className);
		//获取这个类的全部属性信息
		/*
		 * getDeclaredFields 用于获取类中声明的
		 * 属性信息
		 */
		Field[] flds=cls.getDeclaredFields();
		for(Field f:flds){
			System.out.println(f);
		}
		
		
	}

}



