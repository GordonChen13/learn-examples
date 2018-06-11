package reflect;

import java.lang.reflect.Method;
import java.util.Scanner;

public class Demo05 {

	public static void main(String[] args) 
		throws Exception {
		//�ҵ�һ��������testΪ��ͷ����
		Scanner in = new Scanner(System.in);
		String className = in.nextLine();
		Class cls = Class.forName(className);
		Method[] methods =
			cls.getDeclaredMethods();
		Object obj = cls.newInstance();
		for(Method m:methods){
			//getName() ��ȡ������Ϣ�еķ�����
			String name = m.getName();
			if(name.startsWith("test")){
				//m ����testΪ��ͷ����
				System.out.println(m); 
				//num ����һ�������Ĳ�������
				int num = 
					m.getParameterTypes().length;
				if(num==0){
					//ִ����testΪ��ͷ�ķ���
					Object val=m.invoke(obj);
					System.out.println(val); 
				}
			}
		}
	}

}


