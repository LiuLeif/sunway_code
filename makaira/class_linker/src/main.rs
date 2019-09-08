use class_parser;
mod makaira_object;

fn main() {
    let mut path = std::env::home_dir().unwrap();
    path.push("Gitbox/code/makaira/example/Test.class");
    let class_file = class_parser::parse(path.to_str().unwrap());

    let makaira_class = makaira_object::MakairaClass::new(class_file);
    println!("{:#?}", makaira_class);
}
