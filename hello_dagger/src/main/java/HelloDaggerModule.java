// 2019-06-29 19:08
// note: 需要 gradle 4.x, 最新的 gradle 5.x 无法编译
import javax.inject.Inject;
import javax.inject.Scope;
import dagger.Component;
import dagger.Module;
import dagger.Provides;

class BBB {
  public BBB() {}

}

// -----------------------------------
@Module
class TestModule {
  @Provides
  BBB provideBBB() {
    return new BBB();
  }
}

@Component(modules = TestModule.class)
interface AAAComponent {
  void inject(AAA aaa);

}

class AAA {
  @Inject
  BBB mClient;
}

public class HelloDaggerModule {
  static void testModule () {
    AAA aaa = new AAA();
    DaggerAAAComponent.create().inject(aaa);
  }

  public static void main(String argv[]) {
    testModule();
  }
}
