
@256
D=A
@SP
M=D
        
// call
        
@0
D=A
@SP
D=M-D
@new_arg
M=D
        
@Sys.init.ret.1
D=A
@SP
AM=M+1
A=A-1
M=D

        
@LCL
D=M
@SP
AM=M+1
A=A-1
M=D

@ARG
D=M
@SP
AM=M+1
A=A-1
M=D

@THIS
D=M
@SP
AM=M+1
A=A-1
M=D

@THAT
D=M
@SP
AM=M+1
A=A-1
M=D

        
@SP
D=M
@LCL
M=D
@new_arg
D=M
@ARG
M=D
        
@Sys.init
0;JMP
        
(Sys.init.ret.1)
        
// function
(Main.fibonacci)
D=0
        
@0
D=A
@ARG
A=M+D
D=M
@SP
AM=M+1
A=A-1
M=D

        
@2
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
D;JLT
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

        
@SP
AM=M-1
D=M

@IF_TRUE
D;JNE
        
@IF_FALSE
0;JMP
        (IF_TRUE)
@0
D=A
@ARG
A=M+D
D=M
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
        (IF_FALSE)
@0
D=A
@ARG
A=M+D
D=M
@SP
AM=M+1
A=A-1
M=D

        
@2
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

        
// call
        
@1
D=A
@SP
D=M-D
@new_arg
M=D
        
@Main.fibonacci.ret.2
D=A
@SP
AM=M+1
A=A-1
M=D

        
@LCL
D=M
@SP
AM=M+1
A=A-1
M=D

@ARG
D=M
@SP
AM=M+1
A=A-1
M=D

@THIS
D=M
@SP
AM=M+1
A=A-1
M=D

@THAT
D=M
@SP
AM=M+1
A=A-1
M=D

        
@SP
D=M
@LCL
M=D
@new_arg
D=M
@ARG
M=D
        
@Main.fibonacci
0;JMP
        
(Main.fibonacci.ret.2)
        
@0
D=A
@ARG
A=M+D
D=M
@SP
AM=M+1
A=A-1
M=D

        
@1
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

        
// call
        
@1
D=A
@SP
D=M-D
@new_arg
M=D
        
@Main.fibonacci.ret.3
D=A
@SP
AM=M+1
A=A-1
M=D

        
@LCL
D=M
@SP
AM=M+1
A=A-1
M=D

@ARG
D=M
@SP
AM=M+1
A=A-1
M=D

@THIS
D=M
@SP
AM=M+1
A=A-1
M=D

@THAT
D=M
@SP
AM=M+1
A=A-1
M=D

        
@SP
D=M
@LCL
M=D
@new_arg
D=M
@ARG
M=D
        
@Main.fibonacci
0;JMP
        
(Main.fibonacci.ret.3)
        
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
        
// function
(Sys.init)
D=0
        
@4
D=A
@SP
AM=M+1
A=A-1
M=D

            
// call
        
@1
D=A
@SP
D=M-D
@new_arg
M=D
        
@Main.fibonacci.ret.4
D=A
@SP
AM=M+1
A=A-1
M=D

        
@LCL
D=M
@SP
AM=M+1
A=A-1
M=D

@ARG
D=M
@SP
AM=M+1
A=A-1
M=D

@THIS
D=M
@SP
AM=M+1
A=A-1
M=D

@THAT
D=M
@SP
AM=M+1
A=A-1
M=D

        
@SP
D=M
@LCL
M=D
@new_arg
D=M
@ARG
M=D
        
@Main.fibonacci
0;JMP
        
(Main.fibonacci.ret.4)
        (WHILE)
@WHILE
0;JMP
        