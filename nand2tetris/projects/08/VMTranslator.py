#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-09-05 13:33
import ply.lex as lex
import ply.yacc as yacc

reserved = {
    "push",
    "pop",
    "constant",
    "argument",
    "local",
    "temp",
    "this",
    "that",
    "pointer",
    "static",
    "add",
    "sub",
    "eq",
    "lt",
    "gt",
    "neg",
    "and",
    "or",
    "not",
    "label",
    "if-goto",
    "goto",
    "function",
    "return",
    "call",
}
tokens = ["NUMBER", "BRANCH_LABEL"] + [s.upper().replace("-", "_") for s in reserved]

t_ignore_comment = r"//.*"
t_ignore_space = r"\s+"
t_NUMBER = r"\d+"


def t_newline(t):
    r"\n+"
    t.lexer.lineno += len(t.value)


def t_LABEL(t):
    r"[a-zA-Z_][a-zA-Z0-9_\-\.\$]*"
    if t.value in reserved:
        t.type = t.value.upper().replace("-", "_")
    else:
        t.type = "BRANCH_LABEL"
    return t


def t_error(t):
    print("Illegal character '%s'" % t.value[0])
    t.lexer.skip(1)


lexer = lex.lex()
# with open("MemoryAccess/BasicTest/BasicTest.vm") as f:
#     data = f.read()
# lexer.input(data)
# while True:
#     tok = lexer.token()
#     if not tok:
#         break
#     print(tok)

# emitter
class Emitter:
    snippet_pop_a = "@SP\nAM=M-1\n"

    snippet_pop_d = "@SP\nAM=M-1\nD=M\n"

    snippet_push_d = "@SP\nAM=M+1\nA=A-1\nM=D\n"

    def __init__(self, output):
        self.output = open(output, "w")
        self.bool_index = 0
        self.call_index = 0

    def close(self):
        self.output.close()

    def bootstrap(self):
        inst = """
@256
D=A
@SP
M=D
        """
        self.emit_inst(inst)
        self.emit_call("Sys.init", 0)

    def emit_inst(self, inst):
        self.output.write(inst)

    def emit_function(self, function, n_local):
        # before the function:
        # 1. new SP,LCL,ARG is set
        # 2. old frame is pushed
        #
        # in function:
        # 1. declare function symbol
        # 2. reserve local (with n_args)
        #
        inst = f"""
// function
({function})
D=0
        """
        for i in range(n_local):
            inst += f"{self.snippet_push_d}\n"

        self.emit_inst(inst)

    def emit_return(self):
        # return:
        # 1. save ret_addr and old arg
        # 2. pop ret to *arg
        # 3. reset sp to LCL
        # 4. pop and restore THAT, THIS, ARG, LCL
        # 5. set sp to old arg
        # 6. jump to ret_addr
        inst = f"""
// return
        """
        # save ret_addr and old arg
        inst += f"""
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
        """
        # pop return value to arg
        inst += f"""
{self.snippet_pop_d}
@ARG
A=M
M=D
        """
        # set LCL as SP
        inst += f"""
@LCL
D=M
@SP
M=D
        """
        # pop and restore THAT, THIS, ARG, LCL
        inst += f"""
{self.snippet_pop_d}
@THAT
M=D
{self.snippet_pop_d}
@THIS
M=D
{self.snippet_pop_d}
@ARG
M=D
{self.snippet_pop_d}
@LCL
M=D
        """
        # set old arg as SP
        inst += f"""
@old_arg
D=M+1
@SP
M=D
        """
        # jump to ret_addr
        inst += f"""
@ret_addr
A=M
0;JMP
        """
        self.emit_inst(inst)

    def emit_call(self, label, n_arg):
        # 1. save new_arg
        # 2. push ret_addr
        # 3. push LCL, ARG, THIS, THAT
        # 4. set LCL, SP, ARG
        # 5. jump to label
        # 6. define reg_addr label
        inst = """
// call
        """
        # 1. save new_arg
        inst += f"""
@{n_arg}
D=A
@SP
D=M-D
@new_arg
M=D
        """
        # 2. push ret_addr
        self.call_index += 1

        inst += f"""
@{label}.ret.{self.call_index}
D=A
{self.snippet_push_d}
        """
        # 3. push LCL, ARG, THIS, THAT
        inst += f"""
@LCL
D=M
{self.snippet_push_d}
@ARG
D=M
{self.snippet_push_d}
@THIS
D=M
{self.snippet_push_d}
@THAT
D=M
{self.snippet_push_d}
        """
        # 4. set LCL, SP, ARG
        inst += f"""
@SP
D=M
@LCL
M=D
@new_arg
D=M
@ARG
M=D
        """
        # 5. jump to label
        inst += f"""
@{label}
0;JMP
        """
        # 6. define reg_addr label
        inst += f"""
({label}.ret.{self.call_index})
        """

        self.emit_inst(inst)

    def emit_goto(self, label):
        inst = f"""
@{label}
0;JMP
        """
        self.emit_inst(inst)

    def emit_if_goto(self, label):
        inst = f"""
{self.snippet_pop_d}
@{label}
D;JNE
        """
        self.emit_inst(inst)

    def emit_label(self, label):
        inst = f"""
({label})
        """
        self.emit_inst(inst)

    def emit_bool(self, op):
        self.bool_index += 1
        if op == "eq":
            op = "D;JEQ"
        elif op == "lt":
            op = "D;JLT"
        elif op == "gt":
            op = "D;JGT"

        inst = f"""
{self.snippet_pop_d}
{self.snippet_pop_a}
D=M-D
@BOOL_TRUE_{self.bool_index}
{op}
@0
D=A
@BOOL_END_{self.bool_index}
0;JMP
(BOOL_TRUE_{self.bool_index})
@1
D=-A
(BOOL_END_{self.bool_index})
{self.snippet_push_d}
        """
        self.emit_inst(inst)

    def emit_unary(self, op):
        if op == "not":
            op = "!"
        else:
            op = "-"
        inst = f"""
@SP
A=M-1
M={op}M
        """
        self.emit_inst(inst)

    def emit_binary(self, op):
        # pop a, pop b, push (a+b)
        op = {"add": "D=D+M", "sub": "D=M-D", "and": "D=D&M", "or": "D=D|M"}[op]

        inst = f"""
{self.snippet_pop_d}
{self.snippet_pop_a}
{op}
{self.snippet_push_d}
        """
        self.emit_inst(inst)

    def emit_push(self, memory, num):
        # *sp=*(memory+num); sp++
        memories = {"local": "LCL", "argument": "ARG", "this": "THIS", "that": "THAT"}
        if memory in memories:
            inst = f"""
@{num}
D=A
@{memories[memory]}
A=D+M
D=M
{self.snippet_push_d}
        """
        elif memory == "static":
            global clazz
            inst = f"""
@{clazz}.static.{num}
D=M
{self.snippet_push_d}
            """
        elif memory == "pointer":
            if num == 0:
                dest = "THIS"
            else:
                dest = "THAT"
            inst = f"""
@{dest}
D=M
{self.snippet_push_d}
            """
        elif memory == "temp":
            inst = f"""

@{num}
D=A
@5
A=D+A
D=M
{self.snippet_push_d}
            """
        elif memory == "constant":
            inst = f"""
@{num}
D=A
{self.snippet_push_d}
            """
        self.emit_inst(inst)

    def emit_pop(self, memory, num):
        # sp--; *(memory+num)= *sp
        memories = {"local": "LCL", "argument": "ARG", "this": "THIS", "that": "THAT"}
        if memory in memories:
            inst = f"""
@{num}
D=A
@{memories[memory]}
D=D+M
@R13
M=D

{self.snippet_pop_d}
@R13
A=M
M=D
            """

        elif memory == "static":
            global clazz
            inst = f"""
{self.snippet_pop_d}
@{clazz}.static.{num}
M=D
            """
        elif memory == "pointer":
            if num == 0:
                dest = "THIS"
            else:
                dest = "THAT"
            inst = f"""
{self.snippet_pop_d}
@{dest}
M=D
            """
        elif memory == "temp":
            inst = f"""
@{num}
D=A
@5
D=D+A
@R13
M=D

{self.snippet_pop_d}
@R13
A=M
M=D
            """
        self.emit_inst(inst)


def p_error(p):
    print("Syntax error at ", p)


def p_program(p):
    """stmt : stmt stmt"""


def p_label(p):
    """stmt : LABEL BRANCH_LABEL"""
    global emitter
    emitter.emit_label(p[2])


def p_goto(p):
    """stmt : GOTO BRANCH_LABEL"""
    global emitter
    emitter.emit_goto(p[2])


def p_if_goto(p):
    """stmt : IF_GOTO BRANCH_LABEL"""
    global emitter
    emitter.emit_if_goto(p[2])


def p_bool_op(p):
    """stmt : EQ
            | GT
            | LT
    """
    global emitter
    emitter.emit_bool(p[1])


def p_unary_op(p):
    """stmt : NEG
            | NOT
    """
    global emitter
    emitter.emit_unary(p[1])


def p_binary_op(p):
    """stmt : ADD
            | SUB
            | OR
            | AND
    """
    global emitter
    emitter.emit_binary(p[1])


def p_push_op(p):
    """stmt : PUSH LOCAL NUMBER
            | PUSH ARGUMENT NUMBER
            | PUSH THIS NUMBER
            | PUSH THAT NUMBER
            | PUSH TEMP NUMBER
            | PUSH CONSTANT NUMBER
            | PUSH POINTER NUMBER
            | PUSH STATIC NUMBER
    """
    global emitter
    emitter.emit_push(p[2], int(p[3]))


def p_pop_op(p):
    """stmt : POP LOCAL NUMBER
            | POP ARGUMENT NUMBER
            | POP THIS NUMBER
            | POP THAT NUMBER
            | POP TEMP NUMBER
            | POP CONSTANT NUMBER
            | POP POINTER NUMBER
            | POP STATIC NUMBER
    """
    global emitter
    emitter.emit_pop(p[2], int(p[3]))


def p_function(p):
    """stmt : FUNCTION BRANCH_LABEL NUMBER"""
    global emitter
    emitter.emit_function(p[2], int(p[3]))


def p_return(p):
    """stmt : RETURN"""
    global emitter
    emitter.emit_return()


def p_call(p):
    """stmt : CALL BRANCH_LABEL NUMBER"""
    global emitter
    emitter.emit_call(p[2], int(p[3]))


parser = yacc.yacc()

import sys

if len(sys.argv) != 3:
    print("usage: VMTranslator.py <input> <output>")
    sys.exit(1)

import glob, os

input_file = sys.argv[1]
output_file = sys.argv[2]
emitter = Emitter(output_file)

if os.path.isdir(input_file):
    input_files = glob.glob(input_file + "/*.vm")
    emitter.bootstrap()
else:
    input_files = [input_file]

clazz = ""
for input_file in input_files:
    with open(input_file) as f:
        clazz = os.path.basename(input_file).replace(".vm", "")
        data = f.read()
        parser.parse(data)
        
emitter.close()
