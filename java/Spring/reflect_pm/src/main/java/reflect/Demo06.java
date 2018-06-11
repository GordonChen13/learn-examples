package reflect;

import java.lang.reflect.Method;
import java.util.Scanner;

public class Demo06 {

	public static void main(String[] args)
		throws Exception {
		Scanner in = new Scanner(System.in);
		String className = in.nextLine();
		Class cls=Class.forName(className);
		//�ҵ���ִ�еķ�����Ϣ
		Method m = cls.getDeclaredMethod(
			"add", int.class, String.class);
		Object obj = cls.newInstance();
		//ִ��˽�з���
		m.setAccessible(true);//�ƻ�˽�й���
		Object val = m.invoke(obj, 5, "Tom");
		System.out.println(val); 
	}

}




