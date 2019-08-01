class Main {
    static void build (build_settings) {
        def spec=new BuildSpec();

        build_settings=build_settings.rehydrate(spec,spec,spec);
        build_settings();
    }

    static void main(String[] args) {
        build {
            command "compile it"
            depends "a","b","c"
        }
    }
}

class BuildSpec {
    void command(String x) {
        println("command:"+x)
    }

    void depends(String... x) {
        println("depends:"+x)
    }
}
