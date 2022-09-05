
// function
(Sys.init)
D=0
        
@4000
D=A
@SP
AM=M+1
A=A-1
M=D

            
@SP
AM=M-1
D=M

@THIS
M=D
            
@5000
D=A
@SP
AM=M+1
A=A-1
M=D

            
@SP
AM=M-1
D=M

@THAT
M=D
            
// call
        
@0
D=A
@SP
D=M-D
@old_arg
M=D
        
@Sys.main.ret.1
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
@old_arg
D=M
@ARG
M=D
        
@Sys.main
0;JMP
        
(Sys.main.ret.1)
        
@1
D=A
@5
D=A+D
@R13
M=D

@SP
AM=M-1
D=M

@R13
A=M
M=D
            (LOOP)
@LOOP
0;JMP
        
// function
(Sys.main)
D=0
        @SP
AM=M+1
A=A-1
M=D

@SP
AM=M+1
A=A-1
M=D

@SP
AM=M+1
A=A-1
M=D

@SP
AM=M+1
A=A-1
M=D

@SP
AM=M+1
A=A-1
M=D


@4001
D=A
@SP
AM=M+1
A=A-1
M=D

            
@SP
AM=M-1
D=M

@THIS
M=D
            
@5001
D=A
@SP
AM=M+1
A=A-1
M=D

            
@SP
AM=M-1
D=M

@THAT
M=D
            
@200
D=A
@SP
AM=M+1
A=A-1
M=D

            
@1
D=A
@LCL
D=M+D
@R13
M=D

@SP
AM=M-1
D=M

@R13
A=M
M=D
            
@40
D=A
@SP
AM=M+1
A=A-1
M=D

            
@2
D=A
@LCL
D=M+D
@R13
M=D

@SP
AM=M-1
D=M

@R13
A=M
M=D
            
@6
D=A
@SP
AM=M+1
A=A-1
M=D

            
@3
D=A
@LCL
D=M+D
@R13
M=D

@SP
AM=M-1
D=M

@R13
A=M
M=D
            
@123
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
@old_arg
M=D
        
@Sys.add12.ret.2
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
@old_arg
D=M
@ARG
M=D
        
@Sys.add12
0;JMP
        
(Sys.add12.ret.2)
        
@0
D=A
@5
D=A+D
@R13
M=D

@SP
AM=M-1
D=M

@R13
A=M
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

        
@2
D=A
@LCL
A=M+D
D=M
@SP
AM=M+1
A=A-1
M=D

        
@3
D=A
@LCL
A=M+D
D=M
@SP
AM=M+1
A=A-1
M=D

        
@4
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
(Sys.add12)
D=0
        
@4002
D=A
@SP
AM=M+1
A=A-1
M=D

            
@SP
AM=M-1
D=M

@THIS
M=D
            
@5002
D=A
@SP
AM=M+1
A=A-1
M=D

            
@SP
AM=M-1
D=M

@THAT
M=D
            
@0
D=A
@ARG
A=M+D
D=M
@SP
AM=M+1
A=A-1
M=D

        
@12
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
        