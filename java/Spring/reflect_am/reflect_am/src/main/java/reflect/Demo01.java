package reflect;

import java.util.Scanner;

public class Demo01 {

	public static void main(String[] args)
		throws Exception{
		/*
		 * ��̬�����ൽ�ڴ���
		 */
		Scanner in = new Scanner(System.in);
		//�����ڼ�ӿ���̨����̬����ȡ��������
		String className = in.nextLine();
		//�ڳ�������֮ǰ�� �ǲ�ָ��������ʲô��
		/*
		 * ��̬�����ൽ�������У�����������ʱ��
		 * ������Ӧ�Ĵ�����û��class�ļ�,�ͷ���
		 * ��û���ҵ��쳣��
		 */
		Class cls = Class.forName(className);
		//�����صĽ��
		System.out.println(cls); 
	}

}





