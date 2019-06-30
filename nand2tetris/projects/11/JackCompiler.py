#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-09-05 13:33
import ply.lex as lex
import ply.yacc as yacc

keywords = {
    "class",
    "function",
    "method",
    "void",
    "int",
    "char",
    "boolean",
    "var",
    "let",
    "while",
    "do",
    "static",
    "field",
    "return",
    "constructor",
    "this",
    "null",
    "true",
    "false",
    "if",
    "else",
}
tokens = [
    "NUMBER",
    "STRING",
    "OPEN_PARENT",
    "CLOSE_PARENT",
    "OPEN_BRACKET",
    "CLOSE_BRACKET",
    "OPEN_CURLY",
    "CLOSE_CURLY",
    "IDENT",
    "SEMICOLON",
    "DOT",
    "ADD",
    "MINUS",
    "TIMES",
    "DIVIDE",
    "EQUAL",
    "COMMA",
    "AND",
    "OR",
    "NOT",
    "GT",
    "LT",
] + [s.upper() for s in keywords]


def t_comment(t):
    r"//.*"


def t_comment_long(t):
    r"/\*+(.|\n)*?\*+/"
    t.lexer.lineno += t.value.count("\n")


t_ignore_space = r"\s+"

t_COMMA = r","
t_EQUAL = r"="
t_NUMBER = r"\d+"
t_STRING = r'".*"'
t_OPEN_PARENT = r"\("
t_CLOSE_PARENT = r"\)"
t_OPEN_BRACKET = r"\["
t_CLOSE_BRACKET = r"\]"
t_OPEN_CURLY = r"{"
t_CLOSE_CURLY = r"}"
t_SEMICOLON = r";"
t_DOT = r"\."
t_ADD = r"\+"
t_MINUS = r"\-"
t_TIMES = r"\*"
t_DIVIDE = r"\/"
t_AND = r"&"
t_OR = r"\|"
t_NOT = r"~"
t_GT = r">"
t_LT = r"<"


def t_newline(t):
    r"\n+"
    t.lexer.lineno += len(t.value)


def t_IDENT(t):
    r"[a-zA-Z_][a-zA-Z0-9_\-\$]*"
    if t.value in keywords:
        t.type = t.value.upper()
    else:
        t.type = "IDENT"
    return t


def t_error(t):
    print("Illegal character '%s'" % t.value[0])
    t.lexer.skip(1)


lexer = lex.lex()
# with open("Square/SquareGame.jack") as f:
#     data = f.read()
#     lexer.input(data)
# while True:
#     tok = lexer.token()
#     if not tok:
#         break
#     print(tok)


from CompileEngine import CompileEngine


def p_error(p):
    print("Syntax error at ", p)


def p_class(p):
    """class : CLASS class_name OPEN_CURLY class_var_dec subroutine_dec CLOSE_CURLY"""
    global engine
    engine.finish()


def p_class_name(p):
    """class_name : IDENT"""
    global engine
    engine.set_class_name(p[1])


def p_class_var_dec(p):
    """class_var_dec : class_var_dec STATIC type IDENT var_list SEMICOLON class_var_dec
                     | class_var_dec FIELD type IDENT var_list SEMICOLON 
                     | empty
    """
    global engine
    if len(p) != 2:
        engine.mark_class_symbol(p[3], p[4], p[2])
        for i in p[5]:
            engine.mark_class_symbol(p[3], i, p[2])


def p_empty(p):
    """empty :"""


def p_var_list(p):
    """var_list : COMMA IDENT var_list
                | empty
    """
    if len(p) == 2:
        p[0] = []
    else:
        p[0] = [p[2]] + p[3]


def p_type(p):
    """ type : INT
             | BOOLEAN
             | CHAR
             | IDENT
    """
    p[0] = p[1]


def p_subroutine_dec(p):
    """subroutine_dec :  subroutine_header subroutine_body subroutine_dec
    """


def p_subroutine_header(p):
    """subroutine_header : subroutine_category subroutine_return_type IDENT OPEN_PARENT parameter_list CLOSE_PARENT"""
    global engine
    engine.emit_subroutine_header(p[1], p[3])


def p_subroutine_dec_2(p):
    """subroutine_dec : empty
    """


def p_subroutine_category(p):
    """subroutine_category : CONSTRUCTOR
                           | FUNCTION
                           | METHOD
    """
    p[0] = p[1]
    global engine
    if p[1] == "method":
        engine.subroutine_symbol_table.mark_symbol("nil", "this", "argument")


def p_subroutine_return_type(p):
    """subroutine_return_type : VOID
                              | type
    """


def p_parameter_list(p):
    """parameter_list : parameter_var_list type IDENT 
                      | empty
    """
    global engine
    if len(p) != 2:
        engine.mark_subroutine_symbol(p[2], p[3], "argument")


def p_parameter_var_list(p):
    """parameter_var_list : parameter_var_list type IDENT COMMA 
                          | empty
    """
    global engine
    if len(p) != 2:
        engine.mark_subroutine_symbol(p[2], p[3], "argument")


def p_subroutine_body(p):
    """subroutine_body : OPEN_CURLY var_dec stmt CLOSE_CURLY
    """
    global engine
    engine.fix_subroutine_header()
    engine.reset_subroutine()


def p_var_dec(p):
    """var_dec : var_dec VAR type IDENT name_var_list SEMICOLON
    """
    global engine
    engine.mark_subroutine_symbol(p[3], p[4], "local")
    for i in p[5]:
        engine.mark_subroutine_symbol(p[3], i, "local")


def p_var_dec_empty(p):
    """var_dec : empty
    """


def p_name_var_list(p):
    """name_var_list : COMMA IDENT name_var_list
    """
    p[0] = [p[2]] + p[3]


def p_name_var_list_empty(p):
    """name_var_list : empty
    """
    p[0] = []


def p_stmt(p):
    """stmt : stmt stmt"""


def p_let_stmt(p):
    """stmt : LET IDENT EQUAL expression SEMICOLON"""
    global engine
    engine.emit_pop_symbol(p[2])


def p_let_array_stmt(p):
    """stmt : LET IDENT OPEN_BRACKET expression CLOSE_BRACKET EQUAL expression SEMICOLON"""
    global engine
    engine.emit_pop_symbol_array(p[2])


def p_if_stmt(p):
    """stmt : IF OPEN_PARENT if_expression CLOSE_PARENT OPEN_CURLY stmt CLOSE_CURLY"""
    global engine
    engine.emit_if_false()


def p_if_expression(p):
    """if_expression : expression"""
    global engine
    engine.emit_if_true()


def p_if_else_stmt(p):
    """stmt : IF OPEN_PARENT if_expression CLOSE_PARENT OPEN_CURLY stmt CLOSE_CURLY else OPEN_CURLY stmt CLOSE_CURLY"""
    global engine
    engine.emit_if_end()


def p_else(p):
    """else : ELSE"""
    global engine
    engine.emit_else()


def p_while_stmt(p):
    """stmt : while OPEN_PARENT while_expression CLOSE_PARENT OPEN_CURLY stmt CLOSE_CURLY"""
    global engine
    engine.emit_while_finish()


def p_while_expression(p):
    """while_expression : expression"""
    global engine
    engine.emit_while_end()


def p_while(p):
    "while : WHILE"
    global engine
    engine.emit_while_start()


def p_do_stmt(p):
    """stmt : DO subroutine_call SEMICOLON"""
    global engine
    engine.emit_inst("pop temp 0")


def p_return_stmt(p):
    """stmt : RETURN expression SEMICOLON
            | RETURN SEMICOLON
    """
    global engine
    if len(p) == 3:
        engine.emit_inst("push constant 0")
    engine.emit_inst("return")


def p_expression(p):
    """expression : unary_op expression
    """
    global engine
    engine.emit_unary_op(p[1])


def p_expression_binary_op(p):
    """expression : expression binary_op expression
    """
    global engine
    engine.emit_binary_op(p[2])


def p_expression_2(p):
    """expression : OPEN_PARENT expression CLOSE_PARENT    
    """


def p_expression_single(p):
    """expression : subroutine_call
    """


def p_expression_true_false(p):
    """expression : TRUE
                  | FALSE
                  | NULL
    """
    global engine
    if p[1] == "true":
        engine.emit_inst("push constant 1\nneg")
    else:
        engine.emit_inst("push constant 0")


def p_expression_this(p):
    """expression : THIS"""
    global engine
    engine.emit_this()


def p_expression_array(p):
    """expression : IDENT OPEN_BRACKET expression CLOSE_BRACKET"""
    global engine
    engine.emit_array(p[1])


def p_expression_ident(p):
    """expression : IDENT"""
    global engine
    engine.emit_push_symbol(p[1])


def p_expression_string(p):
    """expression : STRING"""
    global engine
    engine.emit_string(p[1])


def p_expression_single_1(p):
    """expression : NUMBER
    """
    global engine
    engine.emit_number(p[1])


def p_unary_op(p):
    """unary_op : NOT
                | MINUS
    """
    p[0] = p[1]


def p_binary_op(p):
    """binary_op : ADD
                 | MINUS
                 | TIMES
                 | DIVIDE
                 | AND
                 | OR
                 | GT
                 | LT
                 | EQUAL
    """
    p[0] = p[1]


def p_method_call(p):
    """subroutine_call : method_call_ident OPEN_PARENT argument_list CLOSE_PARENT
    """
    global engine
    engine.emit_method_call(p[1], p[3])


def p_method_call_ident(p):
    """method_call_ident : IDENT"""
    global engine
    engine.emit_method_call_ident()
    p[0] = p[1]


def p_function_call(p):
    """subroutine_call : function_call_ident OPEN_PARENT argument_list CLOSE_PARENT
    """
    global engine
    engine.emit_function_call(p[1][0], p[1][1], p[3])


def p_function_call_ident(p):
    """function_call_ident : IDENT DOT IDENT"""
    global engine
    engine.emit_function_call_ident(p[1])
    p[0] = (p[1], p[3])


def p_argument_list(p):
    """argument_list : expression argument_var_list
                     | empty
    """
    if len(p) == 3:
        p[0] = p[2] + 1
    else:
        p[0] = 0


def p_argument_var_list(p):
    """argument_var_list : COMMA expression argument_var_list
                         | empty
    """
    if len(p) == 4:
        p[0] = p[3] + 1
    else:
        p[0] = 0


parser = yacc.yacc()

import sys

if len(sys.argv) != 2:
    print("usage: JackAnalizer.py <input>")
    sys.exit(1)

import glob, os

input_file = sys.argv[1]

if os.path.isdir(input_file):
    input_files = glob.glob(input_file + "/*.jack")
else:
    input_files = [input_file]

for input_file in input_files:
    with open(input_file) as f:
        engine = CompileEngine()
        engine.set_output(input_file.replace("jack", "vm"))
        data = f.read()
        parser.parse(data)
