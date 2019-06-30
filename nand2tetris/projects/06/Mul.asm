(TEST)
@R0
D=M
@ADD
D;JGT
@END
0;JMP

(ADD)
// do add
@R1
D=M
@R2
M=D+M
// dec r0
@R0
M=M-1
@TEST
0;JMP

(END)
@END
0;JMP