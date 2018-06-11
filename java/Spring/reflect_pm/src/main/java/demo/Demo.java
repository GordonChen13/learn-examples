package demo;

public class Demo {

	public static void main(String[] args)
		throws Exception {
		String cfg = "context.xml";
		ApplicationContext ctx = 
				new ApplicationContext(cfg);
		Object bean = ctx.getBean("foo");
		System.out.println(bean); 
	}

}
