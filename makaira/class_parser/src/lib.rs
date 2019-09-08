use std::io::Read;

pub mod class_parser;
pub mod class_file;
pub mod inst_parser;

pub fn parse (path: &str) -> class_file::ClassFile {
    let mut f = std::fs::File::open(path).unwrap();
    let mut input = Vec::new();
    f.read_to_end(&mut input);

    let mut class_file = {
        match class_parser::parse(&input) {
            Ok((_, class_file)) => class_file,
            Err(_) => panic!(""),
        }
    };
    class_file
}

#[test]
fn test_class_parser() {
    let mut path = std::env::home_dir().unwrap();
    path.push("Gitbox/code/makaira/example/Test.class");
    println!("{:#?}", parse(path.to_str().unwrap()));

    // let makaira_class = makaira_object::MakairaClass::new(class_file);
    // println!("{:#?}", makaira_class);
}
