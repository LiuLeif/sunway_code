
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
(Class2.set)
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

        
@SP
AM=M-1
D=M

@Class2.static.0
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

@Class2.static.1
M=D
            
@0
D=A
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
(Class2.get)
D=0
        
@Class2.static.0
D=M
@SP
AM=M+1
A=A-1
M=D

            
@Class2.static.1
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
        
// function
(Sys.init)
D=0
        
@6
D=A
@SP
AM=M+1
A=A-1
M=D

            
@8
D=A
@SP
AM=M+1
A=A-1
M=D

            
// call
        
@2
D=A
@SP
D=M-D
@new_arg
M=D
        
@Class1.set.ret.2
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
        
@Class1.set
0;JMP
        
(Class1.set.ret.2)
        
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
            
@23
D=A
@SP
AM=M+1
A=A-1
M=D

            
@15
D=A
@SP
AM=M+1
A=A-1
M=D

            
// call
        
@2
D=A
@SP
D=M-D
@new_arg
M=D
        
@Class2.set.ret.3
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
        
@Class2.set
0;JMP
        
(Class2.set.ret.3)
        
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
            
// call
        
@0
D=A
@SP
D=M-D
@new_arg
M=D
        
@Class1.get.ret.4
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
        
@Class1.get
0;JMP
        
(Class1.get.ret.4)
        
// call
        
@0
D=A
@SP
D=M-D
@new_arg
M=D
        
@Class2.get.ret.5
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
        
@Class2.get
0;JMP
        
(Class2.get.ret.5)
        (WHILE)
@WHILE
0;JMP
        
// function
(Class1.set)
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

        
@SP
AM=M-1
D=M

@Class1.static.0
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

@Class1.static.1
M=D
            
@0
D=A
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
(Class1.get)
D=0
        
@Class1.static.0
D=M
@SP
AM=M+1
A=A-1
M=D

            
@Class1.static.1
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
        