use capnpc;

fn main() {
    capnpc::CompilerCommand::new()
        .src_prefix("schema")
        .file("schema/data.capnp")
        .output_path("schema/")
        .run()
        .expect("compiling schema");

}
