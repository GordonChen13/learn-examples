package demo;

import java.util.HashMap;
import java.util.List;

import org.dom4j.Document;
import org.dom4j.Element;
import org.dom4j.io.SAXReader;

public class ApplicationContext {
	//beans 用于缓存被管理的对象
	private HashMap<String, Object> beans=
			new HashMap<String, Object>();
	/**
	 * 根据配置文件，初始化容器环境
	 * @param xml
	 */
	public ApplicationContext(String xml) 
		throws Exception {
		/*
		 * 读取遍历配置文件，根据配置文件中的信息
		 * 动态加载类，动态创建对象，将对象缓存到
		 * beans集合中
		 */
		//导入dom4j 读取XML文件
		SAXReader reader = new SAXReader();
		//getClass().getClassLoader()
		//.getResourceAsStream(文件名) 
		// 从"包"中读取文件, 文件在"包"中!!!
		Document doc= reader.read(getClass()
				.getClassLoader()
				.getResourceAsStream(xml));
		//解析XML的内容
		System.out.println(doc.asXML());
		//访问根元素 <beans>
		Element root = doc.getRootElement();
		//查询到 <bean>
		List<Element> list = root.elements();
		//遍历 bean 元素
		for(Element e:list){
			//e代表xml文件中的每个bean元素
			//读取class属性的值，最为类名
			String className=e.attributeValue("class");
			String id = e.attributeValue("id");
			//动态加载类，动态创建对象
			Class cls = Class.forName(className);
			Object obj = cls.newInstance();
			//将对象缓存到beans集合中
			beans.put(id, obj);
			System.out.println(id+":"+obj); 
		}
	}
	
	public Object getBean(String id){
		//从beans集合中查找id对应的对象
		return beans.get(id);
	}
}





