package demo;

import java.util.HashMap;
import java.util.List;

import org.dom4j.Document;
import org.dom4j.Element;
import org.dom4j.io.SAXReader;

public class ApplicationContext {
	//beans ���ڻ��汻����Ķ���
	private HashMap<String, Object> beans=
			new HashMap<String, Object>();
	/**
	 * ���������ļ�����ʼ����������
	 * @param xml
	 */
	public ApplicationContext(String xml) 
		throws Exception {
		/*
		 * ��ȡ���������ļ������������ļ��е���Ϣ
		 * ��̬�����࣬��̬�������󣬽����󻺴浽
		 * beans������
		 */
		//����dom4j ��ȡXML�ļ�
		SAXReader reader = new SAXReader();
		//getClass().getClassLoader()
		//.getResourceAsStream(�ļ���) 
		// ��"��"�ж�ȡ�ļ�, �ļ���"��"��!!!
		Document doc= reader.read(getClass()
				.getClassLoader()
				.getResourceAsStream(xml));
		//����XML������
		System.out.println(doc.asXML());
		//���ʸ�Ԫ�� <beans>
		Element root = doc.getRootElement();
		//��ѯ�� <bean>
		List<Element> list = root.elements();
		//���� bean Ԫ��
		for(Element e:list){
			//e����xml�ļ��е�ÿ��beanԪ��
			//��ȡclass���Ե�ֵ����Ϊ����
			String className=e.attributeValue("class");
			String id = e.attributeValue("id");
			//��̬�����࣬��̬��������
			Class cls = Class.forName(className);
			Object obj = cls.newInstance();
			//�����󻺴浽beans������
			beans.put(id, obj);
			System.out.println(id+":"+obj); 
		}
	}
	
	public Object getBean(String id){
		//��beans�����в���id��Ӧ�Ķ���
		return beans.get(id);
	}
}





