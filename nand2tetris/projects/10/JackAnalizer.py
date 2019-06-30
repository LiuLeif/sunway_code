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


def p_error(p):
    print("Syntax error at ", p)


def p_class(p):
    """class : CLASS IDENT OPEN_CURLY class_var_dec subroutine_dec CLOSE_CURLY"""


def p_class_var_dec(p):
    """class_var_dec : STATIC type IDENT var_list SEMICOLON class_var_dec
                     | FIELD type IDENT var_list SEMICOLON class_var_dec
                     | empty
    """


def p_empty(p):
    """empty :"""


def p_var_list(p):
    """var_list : COMMA IDENT var_list
                | empty
    """


def p_type(p):
    """ type : INT
             | BOOLEAN
             | CHAR
             | IDENT
    """


def p_subroutine_dec(p):
    """subroutine_dec : subroutine_category subroutine_return_type IDENT OPEN_PARENT parameter_list CLOSE_PARENT subroutine_body subroutine_dec
                      | empty
    """


def p_subroutine_category(p):
    """subroutine_category : CONSTRUCTOR
                           | FUNCTION
                           | METHOD
    """


def p_subroutine_return_type(p):
    """subroutine_return_type : VOID
                              | type
    """


def p_parameter_list(p):
    """parameter_list : type IDENT parameter_var_list
                      | empty
    """


def p_parameter_var_list(p):
    """parameter_var_list : COMMA type IDENT parameter_var_list
                          | empty
    """


def p_subroutine_body(p):
    """subroutine_body : OPEN_CURLY var_dec stmt CLOSE_CURLY
    """


def p_var_dec(p):
    """var_dec : VAR type IDENT name_var_list SEMICOLON var_dec
               | empty
    """

def p_name_var_list(p):
    """name_var_list : COMMA IDENT name_var_list
                     | empty
    """


def p_stmt(p):
    """stmt : stmt stmt"""


def p_let_stmt(p):
    """stmt : LET IDENT EQUAL expression SEMICOLON"""


def p_let_array_stmt(p):
    """stmt : LET IDENT OPEN_BRACKET expression CLOSE_BRACKET EQUAL expression SEMICOLON"""


def p_if_stmt(p):
    """stmt : IF OPEN_PARENT expression CLOSE_PARENT OPEN_CURLY stmt CLOSE_CURLY"""


def p_if_else_stmt(p):
    """stmt : IF OPEN_PARENT expression CLOSE_PARENT OPEN_CURLY stmt CLOSE_CURLY ELSE OPEN_CURLY stmt CLOSE_CURLY"""


def p_while_stmt(p):
    """stmt : WHILE OPEN_PARENT expression CLOSE_PARENT OPEN_CURLY stmt CLOSE_CURLY"""


def p_do_stmt(p):
    """stmt : DO subroutine_call SEMICOLON"""


def p_return_stmt(p):
    """stmt : RETURN expression SEMICOLON
            | RETURN SEMICOLON
    """


def p_expression(p):
    """expression : unary_op expression
                  | expression binary_op expression
                  | IDENT OPEN_BRACKET expression CLOSE_BRACKET
                  | OPEN_PARENT expression CLOSE_PARENT    
    """


def p_expression_single(p):
    """expression : NUMBER
                  | STRING
                  | TRUE
                  | FALSE
                  | NULL
                  | THIS
                  | IDENT
                  | subroutine_call
    """
    
def p_unary_op(p):
    """unary_op : NOT
                | MINUS
    """


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


def p_subroutine_call(p):
    """subroutine_call : IDENT OPEN_PARENT argument_list CLOSE_PARENT
                       | IDENT DOT IDENT OPEN_PARENT argument_list CLOSE_PARENT
    """


def p_argument_list(p):
    """argument_list : expression argument_var_list
                     | empty
    """


def p_argument_var_list(p):
    """argument_var_list : COMMA expression argument_var_list
                         | empty
    """


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
        data = f.read()
        parser.parse(data)
