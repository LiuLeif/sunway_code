use crate::class_parser;

#[derive(Default, Debug)]
pub struct ClassFile {
    pub magic: u32,
    pub minor_vesion: u16,
    pub major_vesion: u16,
    pub constant_pool_count: u16,
    pub constant_pool: Vec<ConstantInfo>,
    pub access_flags: u16,
    pub this_class: u16,
    pub super_class: u16,
    pub interface_count: u16,
    pub interfaces: Vec<u16>,
    pub fields_count: u16,
    pub fields: Vec<FieldInfo>,
    pub methods_count: u16,
    pub methods: Vec<MethodInfo>,
    pub attributes_count: u16,
    pub attributes: Vec<AttributeInfo>,
}

#[derive(Debug, Clone, Copy)]
pub enum ConstantType {
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
}

impl ConstantType {
    pub fn from_u8(v: u8) -> Option<Self> {
        match v {
            1 => Some(Self::Utf8),
            3 => Some(Self::Integer),
            4 => Some(Self::Float),
            5 => Some(Self::Long),
            6 => Some(Self::Double),
            7 => Some(Self::Class),
            8 => Some(Self::String),
            9 => Some(Self::Fieldref),
            10 => Some(Self::Methodref),
            11 => Some(Self::InterfaceMethodref),
            12 => Some(Self::NameAndType),
            _ => None,
        }
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
pub enum ConstantInfo {
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
}

#[derive(Default, Debug)]
pub struct FieldInfo {
    pub access_flags: u16,
    pub name_index: u16,
    pub descriptor_index: u16,
    pub attributes_count: u16,
    pub attributes: Vec<AttributeInfo>,
}

#[derive(Default, Debug)]
pub struct MethodInfo {
    pub access_flags: u16,
    pub name_index: u16,
    pub descriptor_index: u16,
    pub attributes_count: u16,
    pub attributes: Vec<AttributeInfo>,
}

#[derive(Debug)]
pub struct AttributeInfo {
    pub attribute_name_index: u16,
    pub attribute_length: u32,
    pub raw_data: Vec<u8>,
}

#[derive(Debug)]
pub struct CodeInfo {
    pub max_stack: u16,
    pub max_locals: u16,
    pub code_length: u32,
    pub code: Vec<u8>,
    pub exception_table_length: u16,
    pub exception_table: Vec<ExceptionTableEntry>,
    pub attributes_count: u16,
    pub attributes: Vec<AttributeInfo>,
}

impl Default for AttributeInfo {
    fn default() -> Self {
        Self {
            attribute_name_index: 0,
            attribute_length: 0,
            raw_data: vec![],
        }
    }
}

#[derive(Debug)]
pub struct ExceptionTableEntry {
    start_pc: u16,
    end_pc: u16,
    handler_pc: u16,
    catch_type: u16,
}

pub enum AccessFlags {
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

impl ClassFile {
    pub fn get_string(&self, index: usize) -> String {
        let utf8_info = &self.constant_pool[index - 1];
        if let ConstantInfo::Utf8Info { tag, length, bytes } = utf8_info {
            String::from_utf8(bytes.to_vec()).unwrap()
        } else {
            panic!("error");
        }
    }
    pub fn get_class_name(&self, index: usize) -> String {
        let class_info = &self.constant_pool[index - 1];
        let name_index = {
            if let ConstantInfo::ClassInfo { tag, name_index } = class_info {
                name_index
            } else {
                panic!("error parse class");
            }
        };
        self.get_string(*name_index as usize)
    }

    pub fn get_code(&self, attributes: &[AttributeInfo]) -> CodeInfo {
        let code_info = attributes
            .iter()
            .find(|raw_info| self.get_string(raw_info.attribute_name_index as usize) == "Code")
            .unwrap();

        class_parser::parse_code(&code_info.raw_data).unwrap().1
    }
}
