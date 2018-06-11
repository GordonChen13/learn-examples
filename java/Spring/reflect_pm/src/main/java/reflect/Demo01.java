package reflect;

import java.util.Scanner;

public class Demo01 {

	public static void main(String[] args)
		throws Exception{
		/*
		 * 动态加载类到内存中
		 */
		Scanner in = new Scanner(System.in);
		//运行期间从控制台“动态”获取“类名”
		String className = in.nextLine();
		//在程序运行之前， 是不指定类名是什么的
		/*
		 * 动态加载类到方法区中，当类名错误时候，
		 * 类名对应的磁盘上没有class文件,就发生
		 * 类没有找到异常！
		 */
		Class cls = Class.forName(className);
		//检查加载的结果
		System.out.println(cls); 
	}

}





