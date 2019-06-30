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
}
tokens = ["NUMBER", "LABEL"] + [s.upper() for s in reserved]

t_ignore_comment = r"//.*"
t_ignore_space = r"\s+"
t_NUMBER = r"\d+"


def t_newline(t):
    r"\n+"
    t.lexer.lineno += len(t.value)


def t_LABEL(t):
    r"[a-zA-Z_][a-zA-Z0-9_\.\$]*"
    if t.value in reserved:
        t.type = t.value.upper()
    else:
        t.type = "LABEL"
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

    def close(self):
        self.output.close()

    def emit_inst(self, inst):
        self.output.write(inst)

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
        op = {"add": "D=M+D", "sub": "D=M-D", "and": "D=D&M", "or": "D=D|M"}[op]

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
A=M+D
D=M
{self.snippet_push_d}
        """
        elif memory == "static":
            inst = f"""
@static.{num}
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
A=A+D
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
D=M+D
@R13
M=D

{self.snippet_pop_d}
@R13
A=M
M=D
            """

        elif memory == "static":
            inst = f"""
{self.snippet_pop_d}
@static.{num}
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
D=A+D
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


parser = yacc.yacc()

import sys

if len(sys.argv) != 3:
    print("usage: VMTranslator.py <input> <output>")
    sys.exit(1)

input_file = sys.argv[1]
output_file = sys.argv[2]

emitter = Emitter(output_file)
with open(input_file) as f:
    data = f.read()

parser.parse(data)
emitter.close()
