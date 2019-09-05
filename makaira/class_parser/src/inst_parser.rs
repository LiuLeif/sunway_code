// 2019-09-03 13:53
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

#[derive(Debug)]
pub enum MakairaInst {
    AALOAD,
    AASTORE,
    ACONST_NULL,
    ALOAD(u8),
    ALOAD_0,
    ALOAD_1,
    ALOAD_2,
    ALOAD_3,
    ANEWARRAY(u16),
    ARETURN,
    ARRAYLENGTH,
    ASTORE(u8),
    ASTORE_0,
    ASTORE_1,
    ASTORE_2,
    ASTORE_3,
    ATHROW,
    BALOAD,
    BASTORE,
    BIPUSH(u8),
    BREAKPOINT,
    CALOAD,
    CASTORE,
    CHECKCAST(u16),
    D2F,
    D2I,
    D2L,
    DADD,
    DALOAD,
    DASTORE,
    DCMPG,
    DCMPL,
    DCONST_0,
    DCONST_1,
    DDIV,
    DLOAD(u8),
    DLOAD_0,
    DLOAD_1,
    DLOAD_2,
    DLOAD_3,
    DMUL,
    DNEG,
    DREM,
    DRETURN,
    DSTORE(u8),
    DSTORE_0,
    DSTORE_1,
    DSTORE_2,
    DSTORE_3,
    DSUB,
    DUP,
    DUP_X1,
    DUP_X2,
    DUP2,
    DUP2_X1,
    DUP2_X2,
    F2D,
    F2I,
    F2L,
    FADD,
    FALOAD,
    FASTORE,
    FCMPG,
    FCMPL,
    FCONST_0,
    FCONST_1,
    FCONST_2,
    FDIV,
    FLOAD,
    FLOAD_0,
    FLOAD_1,
    FLOAD_2,
    FLOAD_3,
    FMUL,
    FNEG,
    FREM,
    FRETURN,
    FSTORE,
    FSTORE_0,
    FSTORE_1,
    FSTORE_2,
    FSTORE_3,
    FSUB,
    GETFIELD,
    GETSTATIC,
    GOTO,
    GOTO_W,
    I2B,
    I2C,
    I2D,
    I2F,
    I2L,
    I2S,
    IADD,
    IALOAD,
    IAND,
    IASTORE,
    ICONST_M1,
    ICONST_0,
    ICONST_1,
    ICONST_2,
    ICONST_3,
    ICONST_4,
    ICONST_5,
    IDIV,
    IF_ACMPEQ,
    IF_ACMPNE,
    IF_ICMPEQ,
    IF_ICMPGE,
    IF_ICMPGT,
    IF_ICMPLE,
    IF_ICMPLT,
    IF_ICMPNE,
    IFEQ,
    IFGE,
    IFGT,
    IFLE,
    IFLT,
    IFNE,
    IFNONNULL,
    IFNULL,
    IINC,
    ILOAD,
    ILOAD_0,
    ILOAD_1,
    ILOAD_2,
    ILOAD_3,
    IMPDEP1,
    IMPDEP2,
    IMUL,
    INEG,
    INSTANCEOF,
    INVOKEDYNAMIC,
    INVOKEINTERFACE,
    INVOKESPECIAL(u16),
    INVOKESTATIC,
    INVOKEVIRTUAL,
    IOR,
    IREM,
    IRETURN,
    ISHL,
    ISHR,
    ISTORE,
    ISTORE_0,
    ISTORE_1,
    ISTORE_2,
    ISTORE_3,
    ISUB,
    IUSHR,
    IXOR,
    JSR,
    JSR_W,
    L2D,
    L2F,
    L2I,
    LADD,
    LALOAD,
    LAND,
    LASTORE,
    LCMP,
    LCONST_0,
    LCONST_1,
    LDC,
    LDC_W,
    LDC2_W,
    LDIV,
    LLOAD,
    LLOAD_0,
    LLOAD_1,
    LLOAD_2,
    LLOAD_3,
    LMUL,
    LNEG,
    LOOKUPSWITCH,
    LOR,
    LREM,
    LRETURN,
    LSHL,
    LSHR,
    LSTORE,
    LSTORE_0,
    LSTORE_1,
    LSTORE_2,
    LSTORE_3,
    LSUB,
    LUSHR,
    LXOR,
    MONITORENTER,
    MONITOREXIT,
    MULTIANEWARRAY,
    NEW,
    NEWARRAY,
    NOP,
    POP,
    POP2,
    PUTFIELD,
    PUTSTATIC,
    RET,
    RETURN,
    SALOAD,
    SASTORE,
    SIPUSH,
    SWAP,
    TABLESWITCH,
    WIDE,
}

fn parse_inst(input: &[u8]) -> IResult<&[u8], MakairaInst> {
    use MakairaInst::*;

    let mut input = input;
    let (tmp, opcode) = be_u8(input)?;
    input = tmp;
    let inst = {
        match opcode {
            0x32 => AALOAD,
            0x53 => AASTORE,
            0x1 => ACONST_NULL,
            0x19 => {
                let (tmp, index) = be_u8(input)?;
                input = tmp;
                ALOAD(index)
            }
            0x2a => ALOAD_0,
            0x2b => ALOAD_1,
            0x2c => ALOAD_2,
            0x2d => ALOAD_3,
            0xbd => {
                let (tmp, (hi, lo)) = pair(be_u8, be_u8)(input)?;
                input = tmp;
                ANEWARRAY(((hi as u16) << 8) + lo as u16)
            }
            0xb0 => ARETURN,
            0xbe => ARRAYLENGTH,
            0x3a => {
                let (tmp, index) = be_u8(input)?;
                input = tmp;
                ASTORE(index)
            }
            0x4b => ASTORE_0,
            0x4c => ASTORE_1,
            0x4d => ASTORE_2,
            0x4e => ASTORE_3,
            0xbf => ATHROW,
            0x33 => BALOAD,
            0x54 => BASTORE,
            0x10 => {
                let (tmp, byte) = be_u8(input)?;
                input = tmp;
                BIPUSH(byte)
            },
            0x34 => CALOAD,
            0x55 => CASTORE,
            0xc0 => {
                let (tmp, (hi, lo)) = pair(be_u8, be_u8)(input)?;
                input = tmp;
                CHECKCAST(((hi as u16) << 8) + lo as u16)
            },
            0x90 => D2F,
            0x8e => D2I,
            0x8f => D2L,
            0x63 => DADD,
            0x31 => DALOAD,
            0x52 => DASTORE,
            0x98 => DCMPG,
            0x97 => DCMPL,
            0x0e => DCONST_0,
            0x0f => DCONST_1,
            0x6f => DDIV,
            0x18 => {
                let (tmp, index) = be_u8(input)?;
                input = tmp;
                DLOAD(index)
            },
            0x26 => DLOAD_0,
            0x27 => DLOAD_1,
            0x28 => DLOAD_2,
            0x29 => DLOAD_3,
            0x6b => DMUL,
            0x77 => DNEG,
            0x73 => DREM,
            0xaf => DRETURN,
            0x39 => {
                let (tmp, index) = be_u8(input)?;
                input = tmp;
                DSTORE(index)
            },
            0x47 => DSTORE_0,
            0x48 => DSTORE_1,
            0x49 => DSTORE_2,
            0x50 => DSTORE_3,
            0x67 => DSUB,
            0x59 => DUP,
            0x5a => DUP_X1,
            0x5b => DUP_X2,
            0x5c => DUP2,
            0x5d => DUP2_X1,
            0x5e => DUP2_X2,
            0x8d => F2D,
            0x8b => F2I,
            0x8c => F2L,
            0x62 => FADD,
            0x30 => FALOAD,
            0x51 => FASTORE,
            0x96 => FCMPG,
            0x95 => FCMPL,
            0x0b => FCONST_0,
            0x0c => FCONST_1,
            0x0d => FCONST_2,
            0xb7 => {
                let (tmp, (hi, lo)) = pair(be_u8, be_u8)(input)?;
                input = tmp;
                INVOKESPECIAL(((hi as u16) << 8) + lo as u16)
            }
            0xb1 => RETURN,
            // _ => Err(Err::Error((input, ErrorKind::Tag))),
            _ => panic!("unknown inst"),
        }
    };
    Ok((input, inst))
}

pub fn parse(input: &[u8]) -> IResult<&[u8], Vec<MakairaInst>> {
    many0(parse_inst)(input)
}
