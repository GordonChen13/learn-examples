package reflect;

import java.lang.reflect.Method;
import java.util.Scanner;

public class Demo04 {

	public static void main(String[] args)
		throws Exception{
		//��̬������
		Scanner  in = new Scanner(System.in);
		String className = in.nextLine();
		//forName() �������ʱ��������ִ��
		// ���������࣬ʵ����ֻ����һ��
		Class cls = Class.forName(className);
		//��̬��ȡ��ķ�����Ϣ
		Method[] methods = 
			cls.getDeclaredMethods();
		for(Method m:methods){
			System.out.println(m); 
		}
			
	}

}





