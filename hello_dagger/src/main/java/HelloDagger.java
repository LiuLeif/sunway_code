// 2019-06-29 19:08
// note: 需要 gradle 4.x, 最新的 gradle 5.x 无法编译
import javax.inject.Inject;
import javax.inject.Scope;
import dagger.Component;
import dagger.Module;
import dagger.Provides;

interface IHttpClient {
  public void download(String url);
}

class HttpClient implements IHttpClient {
  @Inject
  public HttpClient() {}

  @Override
  public void download(String url) {
    System.out.println("download" + url + " using HttpClient");
  }
}

interface ILogger {
  public void log(String s);
}

class ConsoleLogger implements ILogger {
  @Inject
  public ConsoleLogger() {}

  @Override
  public void log(String s) {
    System.out.println(s);
  }
}
// 直接使用 Inject 完成注入
// -------------------------------
@Component
interface BasicDownloaderComponent {
  void inject(BasicDownloader target);
}

class BasicDownloader {
  @Inject
  HttpClient mClient;

  @Inject
  ConsoleLogger mLogger;

  public void download(String url) {
    mClient.download(url);
    mLogger.log("basic download done");
  }
}

// 有两种情况无法直接使用 @Inject:
//
// 1. HttpClient 的实现在第三方库中, 无法给它的构造函数加上 @Inject
// 2. Downloader 需要使用 IHttpClient, 后者无法实例化
//
// 这里需要通过 @Module 和 @Provides 来作一个转换
// -----------------------------------\
@Module
class HttpClientModule {
  @Provides
  IHttpClient provideIHttpClient() {
    return new HttpClient();
  }

  @Provides
  ILogger provideLogger() {
    return new ConsoleLogger();
  }
}

@Component(modules = HttpClientModule.class)
interface AbstractDownloaderComponent {
  void inject(AbstractDownloader target);

}

class AbstractDownloader {
  @Inject
  IHttpClient mClient;

  @Inject
  IHttpClient mClient2;

  @Inject
  ILogger mLogger;

  public void download(String url) {
    mClient.download(url);
    mLogger.log("abstract downloader done");
  }
}
// -------------------------------------
@Scope
@interface ApplicationScope {}

@Module
class ScopedHttpClientModule {
  @Provides
  @ApplicationScope
  IHttpClient provideIHttpClient() {
    return new HttpClient();
  }

  @Provides
  ILogger provideLogger() {
    return new ConsoleLogger();
  }
}

@ApplicationScope
@Component (modules = ScopedHttpClientModule.class)
interface ScopedDownloaderComponent {
  void inject(AbstractDownloader target);
}

public class HelloDagger {
  static void testInjection() {
    System.out.println("-------------------");
    BasicDownloader downloader = new BasicDownloader();
    DaggerBasicDownloaderComponent.create().inject(downloader);
    downloader.download("basic url");
    System.out.println();
  }
  static void testModule () {
    System.out.println("-------------------");
    AbstractDownloader downloader = new AbstractDownloader();
    DaggerAbstractDownloaderComponent.create().inject(downloader);
    downloader.download("abstract url");
    System.out.println();
  }

  static void testScope () {
    System.out.println("-------------------");
    AbstractDownloader downloader = new AbstractDownloader();
    DaggerScopedDownloaderComponent.create().inject(downloader);
    downloader.download("scoped url");
    System.out.println(downloader.mClient);
    System.out.println(downloader.mClient2);
    System.out.println();
  }

  public static void main(String argv[]) {
    testInjection();
    testModule();
    testScope();
  }
}
