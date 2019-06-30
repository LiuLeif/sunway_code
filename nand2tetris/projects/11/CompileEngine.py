#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-09-13 14:38
from SymbolTable import SymbolTable


class CompileEngine:
    def __init__(self):
        self.output = None

        self.insts = []
        self.class_symbol_table = SymbolTable()
        self.subroutine_symbol_table = SymbolTable()
        self.clazz = ""

        self.subroutine_header_index = 0
        self.subroutine_header_dict = {}
        self.subroutine_category = None

        self.while_max_index = -1
        self.if_max_index = -1
        self.while_indexes = []
        self.if_indexes = []

    def set_output(self, file):
        self.output = open(file, "w")
        pass

    def finish(self):
        for i in self.insts:
            self.output.write(i + "\n")
            print(i)
        self.output.close()

    def set_class_name(self, clazz):
        self.clazz = clazz

    def emit_inst(self, inst):
        self.insts.append(inst)
        print(inst)

    def reset_subroutine(self):
        self.subroutine_header_dict = {}
        self.subroutine_header_index = 0
        self.subroutine_category = None
        self.subroutine_symbol_table.reset()

    def emit_subroutine_header(self, category, name):
        self.subroutine_category = category
        self.subroutine_header_index = len(self.insts)
        self.emit_inst("function {clazz}.{name} {n_local}")
        self.subroutine_header_dict = {"clazz": self.clazz, "name": name}
        if category == "constructor":
            self.emit_inst(
                f"push constant {self.class_symbol_table.get_count('field')}"
            )
            self.emit_inst("call Memory.alloc 1\npop pointer 0")
        elif category == "method":
            self.emit_inst("push argument 0\npop pointer 0")

    def fix_subroutine_header(self):
        self.subroutine_header_dict["n_local"] = self.subroutine_symbol_table.get_count(
            "local"
        )
        self.insts[self.subroutine_header_index] = self.insts[
            self.subroutine_header_index
        ].format(**self.subroutine_header_dict)

    def emit_number(self, number):
        inst = f"""push constant {number}"""
        self.emit_inst(inst)

    def emit_unary_op(self, unary_op):
        unary_ops = {"~": "not", "-": "neg"}
        self.emit_inst(unary_ops[unary_op])

    def emit_binary_op(self, binary_op):
        binary_ops = {
            "+": "add",
            "-": "sub",
            "*": "call Math.multiply 2",
            "/": "call Math.divide 2",
            ">": "gt",
            "<": "lt",
            "&": "and",
            "=": "eq",
            "|": "or",
        }
        self.emit_inst(binary_ops[binary_op])

    def emit_method_call(self, function, n_args):
        self.emit_inst(f"call {self.clazz}.{function} {n_args+1}")

    def emit_function_call_ident(self, clazz):
        if self.class_symbol_table.contains(
            clazz
        ) or self.subroutine_symbol_table.contains(clazz):
            self.emit_push_symbol(clazz)

    def emit_method_call_ident(self):
        self.emit_inst("push pointer 0")

    def emit_function_call(self, clazz, function, n_args):
        if self.class_symbol_table.contains(clazz):
            inst = f"call {self.class_symbol_table.get_symbol_type(clazz)}.{function} {n_args+1}"
        elif self.subroutine_symbol_table.contains(clazz):
            inst = f"call {self.subroutine_symbol_table.get_symbol_type(clazz)}.{function} {n_args+1}"
        else:
            inst = f"""call {clazz}.{function} {n_args}"""
        self.emit_inst(inst)

    def mark_class_symbol(self, t, name, symbol_type):
        self.class_symbol_table.mark_symbol(t, name, symbol_type)

    def mark_subroutine_symbol(self, t, name, symbol_type):
        self.subroutine_symbol_table.mark_symbol(t, name, symbol_type)

    def emit_string(self, string):
        string = string.strip('"')
        self.emit_inst(f"push constant {len(string)}\ncall String.new 1")
        for c in string:
            self.emit_inst(f"push constant {ord(c)}\ncall String.appendChar 2")

    def emit_pop_symbol(self, symbol):
        if self.class_symbol_table.contains(symbol):
            symbol_type = self.class_symbol_table.get_type(symbol)
            symbol_index = self.class_symbol_table.get_index(symbol)
            if symbol_type == "static":
                self.emit_inst(f"pop static {symbol_index}")
            else:
                symbol_index = self.class_symbol_table.get_index(symbol)
                self.emit_inst(f"pop this {symbol_index}")

        if self.subroutine_symbol_table.contains(symbol):
            symbol_type = self.subroutine_symbol_table.get_type(symbol)
            symbol_index = self.subroutine_symbol_table.get_index(symbol)
            self.emit_inst(f"pop {symbol_type} {symbol_index}")

    def emit_while_start(self):
        self.while_max_index += 1
        self.while_indexes.append(self.while_max_index)
        self.emit_inst(f"label WHILE_EXP{self.while_max_index}")

    def emit_while_end(self):
        self.emit_inst(f"not\nif-goto WHILE_END{self.while_max_index}")

    def emit_while_finish(self):
        index = self.while_indexes.pop()
        self.emit_inst(f"goto WHILE_EXP{index}\nlabel WHILE_END{index}")

    def emit_push_symbol(self, symbol):
        if self.class_symbol_table.contains(symbol):
            symbol_type = self.class_symbol_table.get_type(symbol)
            symbol_index = self.class_symbol_table.get_index(symbol)
            if symbol_type == "static":
                self.emit_inst(f"push static {symbol_index}")
            else:
                self.emit_inst(f"push this {symbol_index}")
        if self.subroutine_symbol_table.contains(symbol):
            symbol_type = self.subroutine_symbol_table.get_type(symbol)
            symbol_index = self.subroutine_symbol_table.get_index(symbol)
            self.emit_inst(f"push {symbol_type} {symbol_index}")

    def emit_pop_symbol_array(self, array):
        self.emit_inst("pop temp 0")
        self.emit_push_symbol(array)
        inst = "add\npop pointer 1\npush temp 0\npop that 0"
        self.emit_inst(inst)

    def emit_array(self, array):
        self.emit_push_symbol(array)
        inst = "add\npop pointer 1\npush that 0"
        self.emit_inst(inst)

    def emit_this(self):
        if self.subroutine_category == "constructor":
            self.emit_inst("push pointer 0")
        else:
            self.emit_inst("push argument 0")

    def emit_if_true(self):
        self.if_max_index += 1
        self.if_indexes.append(self.if_max_index)
        self.emit_inst(
            f"if-goto IF_TRUE{self.if_max_index}\ngoto IF_FALSE{self.if_max_index}\nlabel IF_TRUE{self.if_max_index}"
        )

    def emit_if_false(self):
        index = self.if_indexes.pop()
        self.emit_inst(f"label IF_FALSE{index}")

    def emit_if_end(self):
        index = self.if_indexes.pop()
        self.emit_inst(f"label IF_END{index}")

    def emit_else(self):
        index = self.if_indexes[-1]
        self.emit_inst(f"goto IF_END{index}")
        self.emit_inst(f"label IF_FALSE{index}")
