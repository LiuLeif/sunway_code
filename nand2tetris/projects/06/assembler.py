#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-08-31 16:34
import ply.lex as lex

reserved = {
    "A",
    "D",
    "M",
    "AD",
    "AM",
    "AMD",
    "MD",
    "JGT",
    "JEQ",
    "JGE",
    "JLT",
    "JNE",
    "JLE",
    "JMP",
}

tokens = [
    "LEFT_P",
    "RIGHT_P",
    "SEMI_COLON",
    "AND",
    "OR",
    "NOT",
    "AT",
    "NUMBER",
    "EQUAL",
    "PLUS",
    "MINUS",
    "LABEL",
] + list(reserved)

t_AND = r"&"
t_OR = r"\|"
t_NOT = r"!"
t_NUMBER = r"\d+"
t_PLUS = r"\+"
t_MINUS = r"-"
t_AT = r"@"
t_EQUAL = r"="
t_LEFT_P = r"\("
t_RIGHT_P = r"\)"
t_SEMI_COLON = r";"
t_ignore_space = r"\ +"
t_ignore_comment = r"//.*"

def t_newline(t):
    r"\n+"
    t.lexer.lineno += len(t.value)


def t_LABEL(t):
    r"[a-zA-Z_][a-zA-Z0-9_\.\$]*"
    if t.value in reserved:
        t.type = t.value
    else:
        t.type = "LABEL"
    return t


def t_error(t):
    print("Illegal character '%s'" % t.value[0])
    t.lexer.skip(1)


import ply.yacc as yacc

symbol_table = {
    "R0": 0,
    "R1": 1,
    "R2": 2,
    "R3": 3,
    "R4": 4,
    "R5": 5,
    "R6": 6,
    "R7": 7,
    "R8": 8,
    "R9": 9,
    "R10": 10,
    "R11": 11,
    "R12": 12,
    "R13": 13,
    "R14": 14,
    "R15": 15,
    "SCREEN": 16384,
    "KBD": 24576,
    "SP": 0,
    "LCL": 1,
    "ARG": 2,
    "THIS": 3,
    "THAT": 4,
}


def p_expression_stmt(p):
    """stmt : stmt stmt
    """


global_lineno = 0


def p_expression_LABEL(p):
    """stmt : LEFT_P LABEL RIGHT_P"""
    global global_lineno
    if first_pass:
        symbol_table[p[2]] = global_lineno


def encode_A_instr(num):
    encoding = [0] * 16
    i = 15
    while num != 0:
        encoding[i] = num & 1
        num >>= 1
        i -= 1
    return encoding


def emit(p):
    if not first_pass:
        print("".join([str(i) for i in p]))


def p_expression_A_NUMBER(p):
    """stmt : AT NUMBER
    """
    global global_lineno
    global_lineno += 1

    p[0] = encode_A_instr(int(p[2]))
    emit(p[0])


global_label_index = 16


def p_expression_A_LABEL(p):
    """stmt : AT LABEL
    """
    global symbol_table, global_label_index, global_lineno
    global_lineno += 1

    label = p[2]
    if not first_pass and label not in symbol_table:
        symbol_table[label] = global_label_index
        global_label_index += 1

    if not first_pass:
        num = symbol_table[label]
        p[0] = encode_A_instr(num)
    emit(p[0])


def p_expression_action_1(p):
    "action : dest EQUAL comp"
    p[0] = p[3] + p[1]


def p_expression_action_2(p):
    "action : comp"
    p[0] = p[1] + [0, 0, 0]


def p_expression_C_1(p):
    "stmt : action SEMI_COLON jump"
    global global_lineno
    global_lineno += 1
    p[0] = [1, 1, 1] + p[1] + p[3]
    emit(p[0])


def p_expression_C_2(p):
    "stmt : action"
    global global_lineno
    global_lineno += 1
    p[0] = [1, 1, 1] + p[1] + [0, 0, 0]
    emit(p[0])


jump_instruction = {
    "JGT": [0, 0, 1],
    "JEQ": [0, 1, 0],
    "JGE": [0, 1, 1],
    "JLT": [1, 0, 0],
    "JNE": [1, 0, 1],
    "JLE": [1, 1, 0],
    "JMP": [1, 1, 1],
}


def p_expression_jump(p):
    """jump : JGT
            | JEQ
            | JGE
            | JLT
            | JNE
            | JLE
            | JMP
    """
    p[0] = jump_instruction[p[1]]


def p_expression_dest(p):
    """dest : M
            | D
            | MD
            | A
            | AM
            | AD 
            | AMD"""
    dest = p[1]
    p[0] = [0, 0, 0]
    if "A" in dest:
        p[0][0] = 1
    if "D" in dest:
        p[0][1] = 1
    if "M" in dest:
        p[0][2] = 1


def p_expression_comp_0(p):
    """comp : NUMBER
    """
    if p[1] == "0":
        p[0] = [0, 1, 0, 1, 0, 1, 0]
    elif p[1] == "1":
        p[0] = [0, 1, 1, 1, 1, 1, 1]


def p_expression_comp_1(p):
    """comp : MINUS NUMBER
    """
    if p[2] == "1":
        p[0] = [0, 1, 1, 1, 0, 1, 0]


def p_expression_comp_2(p):
    """comp : D
            | A
            | M
    """
    if p[1] == "D":
        p[0] = [0, 0, 0, 1, 1, 0, 0]
    else:
        p[0] = [0, 1, 1, 0, 0, 0, 0]
    if p[1] == "M":
        p[0][0] = 1


def p_expression_comp_3(p):
    """comp : NOT D
            | NOT A
            | NOT M
    """
    if p[2] == "D":
        p[0] = [0, 0, 0, 1, 1, 0, 1]
    else:
        p[0] = [0, 1, 1, 0, 0, 0, 1]

    if p[2] == "M":
        p[0][0] = 1


def p_expression_comp_4(p):
    """comp : MINUS D
            | MINUS A
            | MINUS M
    """
    if p[2] == "D":
        p[0] = [0, 0, 0, 1, 1, 1, 1]
    else:
        p[0] = [0, 1, 1, 0, 0, 1, 1]

    if p[1] == "M":
        p[0][0] = 1


def p_expression_comp_5(p):
    """comp : D PLUS NUMBER
            | A PLUS NUMBER
            | M PLUS NUMBER
            | D MINUS NUMBER
            | A MINUS NUMBER
            | M MINUS NUMBER    
    """
    if p[2] == "+":
        if p[1] == "D":
            p[0] = [0, 0, 1, 1, 1, 1, 1]
        else:
            p[0] = [0, 1, 1, 0, 1, 1, 1]
    else:
        if p[1] == "D":
            p[0] = [0, 0, 0, 1, 1, 1, 0]
        else:
            p[0] = [0, 1, 1, 0, 0, 1, 0]
    if p[1] == "M":
        p[0][0] = 1


def p_expression_comp_6(p):
    """comp : D PLUS A
            | D PLUS M 
            | D MINUS A
            | D MINUS M
            | D AND A
            | D AND M
            | D OR A
            | D OR M    
    """
    if p[2] == "+":
        p[0] = [0, 0, 0, 0, 0, 1, 0]
    elif p[2] == "-":
        p[0] = [0, 0, 1, 0, 0, 1, 1]
    elif p[2] == "&":
        p[0] = [0, 0, 0, 0, 0, 0, 0]
    elif p[2] == "|":
        p[0] = [0, 0, 1, 0, 1, 0, 1]

    if p[3] == "M":
        p[0][0] = 1


def p_expression_comp_7(p):
    """comp : A MINUS D
            | M MINUS D
    """
    p[0] = [0, 0, 0, 0, 1, 1, 1]
    if p[1] == "M":
        p[0][0] = 1


def p_error(p):
    print("Syntax error in input!", p)


first_pass = True

parser = yacc.yacc()

import sys

with open(sys.argv[1]) as f:
    prog = f.read()
    # first pass
    lexer = lex.lex()
    parser.parse(prog)
    global_lineno = 0
    # second pass
    first_pass = False
    lexer = lex.lex()
    parser.parse(prog)
