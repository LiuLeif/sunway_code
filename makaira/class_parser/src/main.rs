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

use enum_primitive::*;

#[derive(Default, Debug)]
struct ClassFile {
    magic: u32,
    minor_vesion: u16,
    major_vesion: u16,
    constant_pool_count: u16,
    constant_pool: Vec<ConstantInfo>,
    access_flags: u16,
    this_class: u16,
    super_class: u16,
    interface_count: u16,
    interfaces: Vec<u16>,
    fields_count: u16,
    fields: Vec<FieldInfo>,
    methods_count: u16,
    methods: Vec<MethodInfo>,
    attributes_count: u16,
    attributes: Vec<AttributeInfo>,
}

enum_from_primitive! {
    #[derive(Debug, Clone, Copy)]
    enum ConstantType {
        Class = 7,
        Fieldref = 9,
        Methodref = 10,
        InterfaceMethodref = 11,
        String = 8,
        Integer = 3,
        Float = 4,
        Long = 5,
        Double = 6,
        NameAndType = 12,
        Utf8 = 1,
        MethodHandle = 15,
        MethodType = 16,
        InvokeDynamic = 18,
    }
}

impl Default for ConstantType {
    fn default() -> Self {
        Self::Class
    }
}

impl Default for ConstantInfo {
    fn default() -> Self {
        Self::ClassInfo {
            tag: ConstantType::Class,
            name_index: 0,
        }
    }
}

#[derive(Debug)]
enum ConstantInfo {
    ClassInfo {
        tag: ConstantType,
        name_index: u16,
    },
    FieldrefInfo {
        tag: ConstantType,
        class_index: u16,
        name_and_type_index: u16,
    },
    MethodrefInfo {
        tag: ConstantType,
        class_index: u16,
        name_and_type_index: u16,
    },
    InterfaceMethodrefInfo {
        tag: ConstantType,
        class_index: u16,
        name_and_type_index: u16,
    },
    StringInfo {
        tag: ConstantType,
        string_index: u16,
    },
    IntegerInfo {
        tag: ConstantType,
        bytes: u32,
    },
    FloatInfo {
        tag: ConstantType,
        bytes: u32,
    },
    LongInfo {
        tag: ConstantType,
        high_bytes: u32,
        low_bytes: u32,
    },
    DoubleInfo {
        tag: ConstantType,
        high_bytes: u32,
        low_bytes: u32,
    },
    NameAndTypeInfo {
        tag: ConstantType,
        name_index: u16,
        descriptor_index: u16,
    },
    Utf8Info {
        tag: ConstantType,
        length: u16,
        bytes: Vec<u8>,
    },
    MethodHandleInfo {
        tag: ConstantType,
        reference_kind: u8,
        reference_index: u8,
    },
    MethodTypeInfo {
        tag: ConstantType,
        descriptor_index: u16,
    },
    InvokeDynamicInfo {
        tag: ConstantType,
        bootstrap_method_attr_index: u16,
        name_and_type_index: u16,
    },
}

#[derive(Default, Debug)]
struct FieldInfo {
    access_flags: u16,
    name_index: u16,
    descriptor_index: u16,
    attributes_count: u16,
    attribute_info: Vec<AttributeInfo>,
}

#[derive(Default, Debug)]
struct MethodInfo {
    access_flags: u16,
    name_index: u16,
    descriptor_index: u16,
    attributes_count: u16,
    attributes: Vec<AttributeInfo>,
}

#[derive(Debug)]
enum AttributeInfo {
    ConstantValue {
        attribute_name_index: u16,
        attribute_length: u32,
        constantvalue_index: u16,
    },
    Code {
        attribute_name_index: u16,
        attribute_length: u32,
        max_stack: u16,
        max_locals: u16,
        code_length: u32,
        code: Vec<u8>,
        exception_table_length: u16,
        exception_table: Vec<ExceptionTableEntry>,
        attributes_count: u16,
        attributes: Vec<AttributeInfo>,
    },
}

impl Default for AttributeInfo {
    fn default() -> Self {
        Self::ConstantValue {
            attribute_name_index: 0,
            attribute_length: 0,
            constantvalue_index: 0,
        }
    }
}

#[derive(Debug)]
struct ExceptionTableEntry {
    start_pc: u16,
    end_pc: u16,
    handler_pc: u16,
    catch_type: u16,
}

enum AccessFlags {
    ACC_PUBLIC = 0x1,
    ACC_PRIVATE = 0x2,
    ACC_PROTECTED = 0x4,
    ACC_STATIC = 0x8,
    ACC_FINAL = 0x10,
    ACC_SUPER = 0x20,
    ACC_VOATIVATE = 0x40,
    ACC_TRANSIENT = 0x80,
    ACC_INTERFACE = 0x200,
    ACC_ABSTRACT = 0x400,
    ACC_SYNTHETIC = 0x1000,
    ACC_ANNOTATION = 0x2000,
    ACC_ENUM = 0x4000,
}

fn parse_constant_info(input: &[u8]) -> IResult<&[u8], ConstantInfo> {
    let (input, constant_type) = map_opt(be_u8, |t| ConstantType::from_u8(t))(input)?;
    match constant_type {
        ConstantType::Class => map(be_u16, |name_index| ConstantInfo::ClassInfo {
            tag: constant_type,
            name_index,
        })(input),
        ConstantType::Fieldref => map(
            pair(be_u16, be_u16),
            |(class_index, name_and_type_index)| ConstantInfo::FieldrefInfo {
                tag: constant_type,
                class_index,
                name_and_type_index,
            },
        )(input),
        ConstantType::Methodref => map(
            pair(be_u16, be_u16),
            |(class_index, name_and_type_index)| ConstantInfo::MethodrefInfo {
                tag: constant_type,
                class_index,
                name_and_type_index,
            },
        )(input),
        ConstantType::InterfaceMethodref => map(
            pair(be_u16, be_u16),
            |(class_index, name_and_type_index)| ConstantInfo::InterfaceMethodrefInfo {
                tag: constant_type,
                class_index,
                name_and_type_index,
            },
        )(input),
        ConstantType::String => map(be_u16, |string_index| ConstantInfo::StringInfo {
            tag: constant_type,
            string_index,
        })(input),
        ConstantType::Integer => map(be_u32, |bytes| ConstantInfo::IntegerInfo {
            tag: constant_type,
            bytes,
        })(input),
        ConstantType::Float => map(be_u32, |bytes| ConstantInfo::FloatInfo {
            tag: constant_type,
            bytes,
        })(input),
        ConstantType::Long => map(pair(be_u32, be_u32), |(high_bytes, low_bytes)| {
            ConstantInfo::LongInfo {
                tag: constant_type,
                high_bytes,
                low_bytes,
            }
        })(input),
        ConstantType::Double => map(pair(be_u32, be_u32), |(high_bytes, low_bytes)| {
            ConstantInfo::DoubleInfo {
                tag: constant_type,
                high_bytes,
                low_bytes,
            }
        })(input),
        ConstantType::NameAndType => map(pair(be_u16, be_u16), |(name_index, descriptor_index)| {
            ConstantInfo::NameAndTypeInfo {
                tag: constant_type,
                name_index,
                descriptor_index,
            }
        })(input),

        _ => Err(Err::Error((input, ErrorKind::Tag))),
    }
}

fn parse_constant_pool(input: &[u8]) -> IResult<&[u8], (u16, Vec<ConstantInfo>)> {
    pair(be_u16, many0(parse_constant_info))(input)
}

fn parse_class(input: &[u8]) -> IResult<&[u8], ClassFile> {
    let mut ret = ClassFile::default();
    let (input, (magic, minor_vesion, major_vesion)) = tuple((be_u32, be_u16, be_u16))(input)?;
    ret.magic = magic;
    ret.minor_vesion = minor_vesion;
    ret.major_vesion = major_vesion;
    let (input, (constant_pool_count, constant_pool)) = parse_constant_pool(input)?;
    ret.constant_pool_count = constant_pool_count;
    ret.constant_pool = constant_pool;
    Ok((&[], ret))
}

fn main() {
    let mut path = std::env::home_dir().unwrap();
    path.push("Gitbox/code/makaira/class_parser/Test.class.sample");
    let mut f = std::fs::File::open(path).unwrap();
    let mut buff = Vec::new();
    f.read_to_end(&mut buff);

    println!("{:?}", parse_class(&buff));
}
