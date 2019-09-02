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
    fn get_string(&self, index: usize) -> String {
        let utf8_info = &self.class_file.constant_pool[index - 1];
        if let ConstantInfo::Utf8Info { tag, length, bytes } = utf8_info {
            String::from_utf8(bytes.to_vec()).unwrap()
        } else {
            panic!("error");
        }
    }

    fn get_class_name (&self, index:usize) -> String {
        let class_info = &self.class_file.constant_pool[index - 1];
        let name_index = {
            if let ConstantInfo::ClassInfo { tag, name_index } = class_info {
                name_index
            } else {
                panic!("error parse class");
            }
        };
        self.get_string(*name_index as usize)
    }

    pub fn new(class_file: ClassFile) -> Self {
        let mut ret = MakairaClass {
            class_file,
            name: "".to_owned(),
            super_name: "".to_owned(),
            access_flags: 0,
            fields: vec![],
            methods: vec![],
        };

        ret.name = ret.get_class_name(ret.class_file.this_class as usize);
        ret.super_name = ret.get_class_name(ret.class_file.super_class as usize);
        ret
    }
}

#[derive(Debug)]
struct MakairaMethod {}

#[derive(Debug)]
struct MakairaField {}
