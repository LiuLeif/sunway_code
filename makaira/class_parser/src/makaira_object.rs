// 2019-09-02 22:47
use crate::class_file::*;
use crate::inst_parser;
use crate::inst_parser::*;

#[derive(Debug)]
pub struct MakairaClass {
    name: String,
    super_name: String,
    access_flags: u16,
    fields: Vec<MakairaField>,
    methods: Vec<MakairaMethod>,
}

impl MakairaClass {
    pub fn new(class_file: ClassFile) -> Self {
        let name = class_file.get_class_name(class_file.this_class as usize);
        let super_name = class_file.get_class_name(class_file.super_class as usize);

        let fields = class_file
            .fields
            .iter()
            .map(|field_info| MakairaField {
                access_flags: field_info.access_flags,
                name: class_file.get_string(field_info.name_index as usize),
                descriptor: class_file.get_string(field_info.descriptor_index as usize),
            })
            .collect::<Vec<_>>();

        let methods = class_file
            .methods
            .iter()
            .map(|method_info| MakairaMethod {
                access_flags: method_info.access_flags,
                name: class_file.get_string(method_info.name_index as usize),
                descriptor: class_file.get_string(method_info.descriptor_index as usize),
                code: {
                    let code = class_file.get_code(&method_info.attributes);
                    MakairaCode {
                        max_stack: code.max_stack,
                        max_locals: code.max_locals,
                        code_length: code.code_length,
                        insts: inst_parser::parse(&code.code).unwrap().1,
                    }
                },
            })
            .collect::<Vec<_>>();

        MakairaClass {
            name,
            super_name,
            access_flags: 0,
            fields,
            methods,
        }
    }
}

#[derive(Debug)]
struct MakairaMethod {
    access_flags: u16,
    name: String,
    descriptor: String,
    code: MakairaCode,
}

#[derive(Debug)]
struct MakairaCode {
    pub max_stack: u16,
    pub max_locals: u16,
    pub code_length: u32,
    pub insts: Vec<MakairaInst>,
}

#[derive(Debug)]
struct MakairaField {
    access_flags: u16,
    name: String,
    descriptor: String,
}
