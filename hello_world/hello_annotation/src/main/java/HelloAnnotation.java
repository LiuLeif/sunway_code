// 2019-06-29 19:08
import java.lang.annotation.Annotation;
import java.lang.reflect.Method;
import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

// 1. 编译 annotation processor
// javac MyProcessor.java
// 2. 编译其它 java 文件, 并指定 processor
// javac -cp . -processor MyProcessor HelloAnnotation.java Author.java
// 3. 运行
// java -cp . HelloAnnotation

@Author("sunway")
class HelloAnnotation2 {}

@Author("sunway")
@RuntimeAuthor (name = "sunway", age = 16)
class HelloAnnotation3 {}

public class HelloAnnotation {
    public static void main(String argv[]) {
        System.out.println(GeneratedAuthors.dump());
        for (Annotation a : HelloAnnotation3.class.getAnnotations()) {
            if (a instanceof RuntimeAuthor) {
                System.out.println(((RuntimeAuthor)a).name());
                System.out.println(((RuntimeAuthor)a).age());
            }
            if (a instanceof Author) {
                System.out.println(((Author)a).value());
            }
        }

    }
}
