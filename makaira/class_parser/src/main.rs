use nom::bytes::complete::*;
use nom::character::complete::*;
use nom::sequence::*;

fn main() {
    // hello:abc+123
    let x = tuple((tag("hello"), tag::<_, &str, ()>("world"), tag("hello")))("helloworldhello");
    let (_, (a, b, c)) = x.unwrap();
    println!("{:?}", (a, b, c));
    // println!("{:?}", type_of(&x));

    // println!("{:?}", type_of(&parser));
    // parser("hello:abc+123")?;
}
