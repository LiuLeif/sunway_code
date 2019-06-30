// 2019-06-29 19:08
import javax.inject.Inject;
import dagger.Component;

class SellMoe {
  @Inject
  public SellMoe() {
  }

  public void sellMoe() {
    System.out.println("赶紧卖了个大萌");
  }
}

@Component
interface HelloDaggerComponent {
  void inject(HelloDagger t);
}

public class HelloDagger {
  @Inject
  SellMoe moe;
  public static void main(String argv[]) {
    HelloDagger dagger = new HelloDagger();
    DaggerHelloDaggerComponent.create().inject(dagger);
    dagger.moe.sellMoe();
  }
}
