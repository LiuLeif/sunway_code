
// function        
(SimpleFunction.test)
D=0        
        @SP
AM=M+1
A=A-1
M=D

@SP
AM=M+1
A=A-1
M=D


@0
D=A
@LCL
A=M+D
D=M
@SP
AM=M+1
A=A-1
M=D

        
@1
D=A
@LCL
A=M+D
D=M
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

        
@SP
A=M-1
M=!M
        
@0
D=A
@ARG
A=M+D
D=M
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

        
@1
D=A
@ARG
A=M+D
D=M
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

        
// return
        
@LCL
D=M
@5
A=D-A
D=M
@ret_addr
M=D

@ARG
D=M
@old_arg
M=D
        
@SP
AM=M-1
D=M

@ARG
A=M
M=D        
        
@LCL
D=M
@SP
M=D        
        
@SP
AM=M-1
D=M

@THAT
M=D
@SP
AM=M-1
D=M

@THIS
M=D
@SP
AM=M-1
D=M

@ARG
M=D
@SP
AM=M-1
D=M

@LCL
M=D                                
        
@old_arg
D=M+1
@SP
M=D        
        
@ret_addr
A=M        
0;JMP        
        