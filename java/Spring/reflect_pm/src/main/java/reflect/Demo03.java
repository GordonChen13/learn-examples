package reflect;

import java.lang.reflect.Field;
import java.util.Scanner;

public class Demo03 {

	public static void main(String[] args)
		throws Exception{
		//��ʾ ĳ���� ��������Ϣ
		Scanner in = new Scanner(System.in);
		String className = in.nextLine();
		//��̬������
		Class cls = Class.forName(className);
		//��ȡ������ȫ��������Ϣ
		/*
		 * getDeclaredFields ���ڻ�ȡ����������
		 * ������Ϣ
		 */
		Field[] flds=cls.getDeclaredFields();
		for(Field f:flds){
			System.out.println(f);
		}
		
		
	}

}



