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

mod parser;
mod class_file;
mod makaira_object;

fn main() {
    let mut path = std::env::home_dir().unwrap();
    path.push("Gitbox/code/makaira/class_parser/Test.class.sample");
    let mut f = std::fs::File::open(path).unwrap();
    let mut input = Vec::new();
    f.read_to_end(&mut input);

    let mut class_file = {
        match parser::parse_class(&input) {
            Ok((_, class_file)) => class_file,
            Err(_) => panic!(""),
        }
    };

    let makaira_class = makaira_object::MakairaClass::new(class_file);
    println!("{:#?}", makaira_class);
}
