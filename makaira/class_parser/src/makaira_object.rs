// 2019-09-02 22:47
use crate::class_file::*;

#[derive(Debug)]
pub struct MakairaClass {
    class_file: ClassFile,
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

        let mut ret = MakairaClass {
            class_file,
            name,
            super_name,
            access_flags: 0,
            fields: vec![],
            methods: vec![],
        };
        ret
    }
}

#[derive(Debug)]
struct MakairaMethod {}

#[derive(Debug)]
struct MakairaField {
    access_flags: u16,
    name: String,
    descriptor: String,
}
