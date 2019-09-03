// 2019-09-03 13:53
pub enum MakairaInstruction {
    ALOAD_0,
    INVOKE_SPECIAL(u16),
    RETURN,
}

fn parse(input: &[u8]) -> IResult<&[u8], Vec<MakairaInstruction>> {}
