
@17
D=A
@SP
AM=M+1
A=A-1
M=D

            
@17
D=A
@SP
AM=M+1
A=A-1
M=D

            
@SP
AM=M-1
D=M

@SP
AM=M-1

D=M-D
@BOOL_TRUE_1
D;JEQ
@0
D=A
@BOOL_END_1
0;JMP
(BOOL_TRUE_1)
@1
D=-A
(BOOL_END_1)
@SP
AM=M+1
A=A-1
M=D

        
@17
D=A
@SP
AM=M+1
A=A-1
M=D

            
@16
D=A
@SP
AM=M+1
A=A-1
M=D

            
@SP
AM=M-1
D=M

@SP
AM=M-1

D=M-D
@BOOL_TRUE_2
D;JEQ
@0
D=A
@BOOL_END_2
0;JMP
(BOOL_TRUE_2)
@1
D=-A
(BOOL_END_2)
@SP
AM=M+1
A=A-1
M=D

        
@16
D=A
@SP
AM=M+1
A=A-1
M=D

            
@17
D=A
@SP
AM=M+1
A=A-1
M=D

            
@SP
AM=M-1
D=M

@SP
AM=M-1

D=M-D
@BOOL_TRUE_3
D;JEQ
@0
D=A
@BOOL_END_3
0;JMP
(BOOL_TRUE_3)
@1
D=-A
(BOOL_END_3)
@SP
AM=M+1
A=A-1
M=D

        
@892
D=A
@SP
AM=M+1
A=A-1
M=D

            
@891
D=A
@SP
AM=M+1
A=A-1
M=D

            
@SP
AM=M-1
D=M

@SP
AM=M-1

D=M-D
@BOOL_TRUE_4
D;JLT
@0
D=A
@BOOL_END_4
0;JMP
(BOOL_TRUE_4)
@1
D=-A
(BOOL_END_4)
@SP
AM=M+1
A=A-1
M=D

        
@891
D=A
@SP
AM=M+1
A=A-1
M=D

            
@892
D=A
@SP
AM=M+1
A=A-1
M=D

            
@SP
AM=M-1
D=M

@SP
AM=M-1

D=M-D
@BOOL_TRUE_5
D;JLT
@0
D=A
@BOOL_END_5
0;JMP
(BOOL_TRUE_5)
@1
D=-A
(BOOL_END_5)
@SP
AM=M+1
A=A-1
M=D

        
@891
D=A
@SP
AM=M+1
A=A-1
M=D

            
@891
D=A
@SP
AM=M+1
A=A-1
M=D

            
@SP
AM=M-1
D=M

@SP
AM=M-1

D=M-D
@BOOL_TRUE_6
D;JLT
@0
D=A
@BOOL_END_6
0;JMP
(BOOL_TRUE_6)
@1
D=-A
(BOOL_END_6)
@SP
AM=M+1
A=A-1
M=D

        
@32767
D=A
@SP
AM=M+1
A=A-1
M=D

            
@32766
D=A
@SP
AM=M+1
A=A-1
M=D

            
@SP
AM=M-1
D=M

@SP
AM=M-1

D=M-D
@BOOL_TRUE_7
D;JGT
@0
D=A
@BOOL_END_7
0;JMP
(BOOL_TRUE_7)
@1
D=-A
(BOOL_END_7)
@SP
AM=M+1
A=A-1
M=D

        
@32766
D=A
@SP
AM=M+1
A=A-1
M=D

            
@32767
D=A
@SP
AM=M+1
A=A-1
M=D

            
@SP
AM=M-1
D=M

@SP
AM=M-1

D=M-D
@BOOL_TRUE_8
D;JGT
@0
D=A
@BOOL_END_8
0;JMP
(BOOL_TRUE_8)
@1
D=-A
(BOOL_END_8)
@SP
AM=M+1
A=A-1
M=D

        
@32766
D=A
@SP
AM=M+1
A=A-1
M=D

            
@32766
D=A
@SP
AM=M+1
A=A-1
M=D

            
@SP
AM=M-1
D=M

@SP
AM=M-1

D=M-D
@BOOL_TRUE_9
D;JGT
@0
D=A
@BOOL_END_9
0;JMP
(BOOL_TRUE_9)
@1
D=-A
(BOOL_END_9)
@SP
AM=M+1
A=A-1
M=D

        
@57
D=A
@SP
AM=M+1
A=A-1
M=D

            
@31
D=A
@SP
AM=M+1
A=A-1
M=D

            
@53
D=A
@SP
AM=M+1
A=A-1
M=D

            
@SP
AM=M-1
D=M

@SP
AM=M-1

D=M+D
@SP
AM=M+1
A=A-1
M=D

        
@112
D=A
@SP
AM=M+1
A=A-1
M=D

            
@SP
AM=M-1
D=M

@SP
AM=M-1

D=M-D
@SP
AM=M+1
A=A-1
M=D

        
@SP
A=M-1
M=-M
        
@SP
AM=M-1
D=M

@SP
AM=M-1

D=D&M
@SP
AM=M+1
A=A-1
M=D

        
@82
D=A
@SP
AM=M+1
A=A-1
M=D

            
@SP
AM=M-1
D=M

@SP
AM=M-1

D=D|M
@SP
AM=M+1
A=A-1
M=D

        
@SP
A=M-1
M=!M
        