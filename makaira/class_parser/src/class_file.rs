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
pub enum AttributeInfo {
    RawInfo {
        attribute_name_index: u16,
        attribute_length: u32,
        raw_data: Vec<u8>,
    },
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
