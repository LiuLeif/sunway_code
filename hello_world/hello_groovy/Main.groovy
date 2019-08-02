class Main {
    static void build (build_settings) {
        def spec=new BuildSpec();

        build_settings=build_settings.rehydrate(spec,spec,spec);
        build_settings();
    }

    static void main(String[] args) {
        build {
            // property 方法
            property "none"
            // property 属性
            property="none"
            command "compile it"
            depends "a","b","c"
        }
    }
}

class BuildSpec {
    // 方法和属性可以重名
    String property

    void property(String x) {
        println("property:"+x)
    }

    void command(String x) {
        println("command:"+x)
    }

    void depends(String... x) {
        println("depends:"+x)
    }
}
