use nom::bytes::complete::*;
use nom::character::complete::*;
use nom::number::complete::*;
// use nom::bits::complete::*;

use nom::branch::*;
use nom::combinator::*;
use nom::error::*;
use nom::multi::*;
use nom::sequence::*;
use nom::*;

use std::io::Read;

mod class_parser;
mod class_file;
mod makaira_object;
mod inst_parser;

fn main() {
    let mut path = std::env::home_dir().unwrap();
    path.push("Gitbox/code/makaira/class_parser/Test.class.sample");
    let mut f = std::fs::File::open(path).unwrap();
    let mut input = Vec::new();
    f.read_to_end(&mut input);

    let mut class_file = {
        match class_parser::parse(&input) {
            Ok((_, class_file)) => class_file,
            Err(_) => panic!(""),
        }
    };

    // println!("{:#?}", class_file);

    let makaira_class = makaira_object::MakairaClass::new(class_file);
    println!("{:#?}", makaira_class);
}
