use nom::bytes::complete::*;
use nom::character::complete::*;
use nom::combinator::*;
use nom::multi::*;
use nom::sequence::*;
use nom::*;

#[derive(Debug)]
enum Gender {
    MALE,
    FEMALE,
}

impl Default for Gender {
    fn default() -> Self {
        Gender::MALE
    }
}

#[derive(Debug, Default)]
struct Person {
    name: String,
    age: i32,
    gender: Gender,
}

fn delim(input: &str) -> IResult<&str, &str> {
    let (input, (_, _)) = pair(tag(","), space0)(input)?;
    Ok((input, ""))
}

fn parse_person(input: &str) -> IResult<&str, Person> {
    let mut person = Person::default();
    let (input, (_, name, _)) = tuple((tag("name:"), alpha0, delim))(input)?;
    person.name = name.to_owned();
    let (input, (_, age, _)) = tuple((
        tag("age:"),
        map_res(digit0, |s: &str| s.parse::<i32>()),
        delim,
    ))(input)?;
    person.age = age;
    let (input, (_, gender)) = pair(
        tag("gender:"),
        map(alpha0, |s: &str| match s {
            "male" => Gender::MALE,
            "female" => Gender::FEMALE,
            _ => Gender::MALE,
        }),
    )(input)?;
    person.gender = gender;

    let (input, (_)) = space0(input)?;

    Ok((input, person))
}

fn main() {
    let s = "name:erq,age:20,gender:mail name:sunxm,age:20,gender:female";
    println!("{:?}", many0(parse_person)(s));
}
