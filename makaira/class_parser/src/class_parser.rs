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

use crate::class_file::*;

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
        ConstantType::Utf8 => {
            let (input, length) = be_u16(input)?;
            let (input, bytes) = take(length)(input)?;
            Ok((
                input,
                ConstantInfo::Utf8Info {
                    tag: constant_type,
                    length,
                    bytes: bytes.to_owned(),
                },
            ))
        } // _ => Err(Err::Error((input, ErrorKind::Tag))),
    }
}

fn parse_constant_pool(input: &[u8]) -> IResult<&[u8], (u16, Vec<ConstantInfo>)> {
    pair(be_u16, many0(parse_constant_info))(input)
}

fn parse_interface(input: &[u8]) -> IResult<&[u8], (u16, Vec<u16>)> {
    let (input, interface_count) = be_u16(input)?;
    let (input, interfaces) = count(be_u16, interface_count as usize)(input)?;
    Ok((input, (interface_count, interfaces)))
}

fn parse_attributes(input: &[u8]) -> IResult<&[u8], (u16, Vec<AttributeInfo>)> {
    let (input, attributes_count) = be_u16(input)?;
    let (input, attribute_info) = count(
        map(
            pair(be_u16, length_data(be_u32)),
            |(attribute_name_index, raw_data)| AttributeInfo {
                attribute_name_index,
                attribute_length: raw_data.len() as u32,
                raw_data: raw_data.to_vec(),
            },
        ),
        attributes_count as usize,
    )(input)?;
    Ok((input, (attributes_count, attribute_info)))
}

fn parse_fields(input: &[u8]) -> IResult<&[u8], (u16, Vec<FieldInfo>)> {
    let (input, fields_count) = be_u16(input)?;
    let (input, fields) = count(
        map(
            tuple((be_u16, be_u16, be_u16, parse_attributes)),
            |(access_flags, name_index, descriptor_index, (attributes_count, attributes))| {
                FieldInfo {
                    access_flags,
                    name_index,
                    descriptor_index,
                    attributes_count,
                    attributes,
                }
            },
        ),
        fields_count as usize,
    )(input)?;
    Ok((input, (fields_count, fields)))
}

fn parse_methods(input: &[u8]) -> IResult<&[u8], (u16, Vec<MethodInfo>)> {
    let (input, methods_count) = be_u16(input)?;
    let (input, methods) = count(
        map(
            tuple((be_u16, be_u16, be_u16, parse_attributes)),
            |(access_flags, name_index, descriptor_index, (attributes_count, attributes))| {
                MethodInfo {
                    access_flags,
                    name_index,
                    descriptor_index,
                    attributes_count,
                    attributes,
                }
            },
        ),
        methods_count as usize,
    )(input)?;
    Ok((input, (methods_count, methods)))
}

pub fn parse(input: &[u8]) -> IResult<&[u8], ClassFile> {
    let mut ret = ClassFile::default();
    let (input, (magic, minor_vesion, major_vesion)) = tuple((be_u32, be_u16, be_u16))(input)?;
    ret.magic = magic;
    ret.minor_vesion = minor_vesion;
    ret.major_vesion = major_vesion;

    let (input, (constant_pool_count, constant_pool)) = parse_constant_pool(input)?;
    ret.constant_pool_count = constant_pool_count;
    ret.constant_pool = constant_pool;

    let (input, (access_flags, this_class, super_class)) = tuple((be_u16, be_u16, be_u16))(input)?;
    ret.access_flags = access_flags;
    ret.this_class = this_class;
    ret.super_class = super_class;

    let (input, (interface_count, interfaces)) = parse_interface(input)?;
    ret.interface_count = interface_count;
    ret.interfaces = interfaces;

    let (input, (fields_count, fields)) = parse_fields(input)?;
    ret.fields_count = fields_count;
    ret.fields = fields;

    let (input, (methods_count, methods)) = parse_methods(input)?;
    ret.methods_count = methods_count;
    ret.methods = methods;

    let (input, (attributes_count, attributes)) = parse_attributes(input)?;
    ret.attributes_count = attributes_count;
    ret.attributes = attributes;

    Ok((input, ret))
}

pub fn parse_code(input: &[u8]) -> IResult<&[u8], CodeInfo> {
    let (input, (max_stack, max_locals, code)) =
        tuple((be_u16, be_u16, length_data(be_u32)))(input)?;
    Ok((
        input,
        CodeInfo {
            max_stack,
            max_locals,
            code_length: code.len() as u32,
            code: code.to_vec(),
            exception_table_length: 0,
            exception_table: vec![],
            attributes_count: 0,
            attributes: vec![],
        },
    ))
}
