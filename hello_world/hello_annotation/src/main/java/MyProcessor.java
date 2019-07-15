// 2019-07-11 14:08
import java.io.PrintWriter;
import java.util.HashMap;
import java.util.Map;
import java.util.Set;
import javax.annotation.processing.*;
import javax.lang.model.element.*;
import javax.tools.*;
import javax.lang.model.*;

@SupportedSourceVersion(SourceVersion.RELEASE_8)
@SupportedAnnotationTypes({"Author", "RuntimeAuthor"})
public class MyProcessor extends AbstractProcessor {
    @Override
    public boolean process(Set<? extends TypeElement> annotations, RoundEnvironment roundEnv) {
        if (annotations.isEmpty()) {
            return true;
        }
        System.out.println("process");
        StringBuilder builder = new StringBuilder()
                                .append("import java.util.HashMap;")
                                .append("public class GeneratedAuthors {\n")
                                .append(" private static HashMap<String,String> sAuthors=new HashMap(); ")
                                .append(" public static String dump() {return sAuthors.toString();}")
                                .append(" static {");

        Map<String, String> authors = new HashMap();
        for (TypeElement annotation : annotations ) {
            for ( Element element : roundEnv.getElementsAnnotatedWith(annotation) ) {
                Author author = element.getAnnotation(Author.class);
                processingEnv.getMessager().printMessage(Diagnostic.Kind.NOTE, "found @Author at " + element + " value: " + author.value());

                RuntimeAuthor runtimeAuthor = element.getAnnotation(RuntimeAuthor.class);
                if (runtimeAuthor != null) {
                    processingEnv.getMessager().printMessage(Diagnostic.Kind.NOTE, "found @RuntimeAuthor at " + element + " value: " + runtimeAuthor.name());
                }

                // process 可以中断编译
                if (!author.value().equals("sunway")) {
                    processingEnv.getMessager().printMessage(Diagnostic.Kind.ERROR, "author must be sunway");
                }
                authors.put(element.toString(), author.value());
                builder.append("sAuthors.put(")
                        .append("\"" + element.toString() + "\"," + "\"" + author.value() + "\"" + ");");
            }
        }
        builder.append("}}");
        // generate code
        try {
            JavaFileObject sourceFile = processingEnv.getFiler().createSourceFile(
                "GeneratedAuthors");

            PrintWriter out = new PrintWriter(sourceFile.openWriter());
            out.write(builder.toString());
            out.close();

        } catch (Exception e) {
        }

        return true;
    }

}
