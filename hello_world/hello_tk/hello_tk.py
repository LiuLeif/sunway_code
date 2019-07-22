#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2019-07-22 12:57
from tkinter import *
from tkinter.ttk import *
import threading
import time

class Application(Frame):
    def __init__(self, master=None):
        super().__init__(master)
        self.pack(fill="both")
        self.create_widgets()

    def create_widgets(self):
        self.tree = Treeview(self, columns=("pid", "id", "status", "percent"), height=768)
        self.tree.heading("pid", text="group id")
        self.tree.heading("id", text="id")
        self.tree.heading("status", text="status")
        self.tree.heading("percent", text="percent")
        self.tree.column("#0", width=10)
        self.tree.tag_configure("deleted", background='gray')
        self.tree.tag_configure("sched", background='gold')
        self.tree.tag_configure("started", background='cyan')
        self.tree.tag_configure("success", background='green')
        self.tree.pack(fill="both")

class Controller:
    def __init__(self, tree):
        self.id_to_info = {}
        self.id_to_widget = {}
        self.tree = tree;

    def bind_view(self, info, tree_item):
        self.tree.item(tree_item, values=(info.pid, info.id, info.status, "{}%".format(info.percent)))

    def requested(self, pid, id):
        info = Info(id=id, pid=pid, status=190);
        self.id_to_info[id] = info
        if id == pid:
            tree_item = app.tree.insert("", "end", open=True)
        else:
            parent_item = self.id_to_widget.get(pid, -1)
            tree_item = app.tree.insert(parent_item, "end")

        self.id_to_widget[id] = tree_item
        self.bind_view(info, tree_item)

    def status(self, id, status):
        if id not in self.id_to_widget:
            return
        tree_item = self.id_to_widget[id]
        info = self.id_to_info[id]
        info.status = status
        self.bind_view(info, tree_item)
        if status == "200":
            self.tree.item(tree_item, tags=("success"))

    def progress(self, id, current, total):
        if id not in self.id_to_widget:
            return
        tree_item = self.id_to_widget[id]
        info = self.id_to_info[id]
        info.percent = int(int(current) * 100 / (int(total) + 1))
        self.bind_view(info, tree_item)

    def deleted(self, id):
        if id not in self.id_to_widget:
            return
        tree_item = self.id_to_widget[id]
        self.tree.item(tree_item, tags=("deleted"))

    def scheduled(self, id):
        if id not in self.id_to_widget:
            return
        tree_item = self.id_to_widget[id]
        self.tree.item(tree_item, tags=("sched"))

    def started(self, id):
        if id not in self.id_to_widget:
            return
        tree_item = self.id_to_widget[id]
        self.tree.item(tree_item, tags=("started"))

class Info:
    def __init__(self, id, pid, status):
        self.id = id
        self.pid = pid;
        self.status = status;
        self.percent = 0;

def parse_log():
    controller = Controller(app.tree)
    # test
    # for i in range(60):
    #     controller.requested(1, i + 1)

    router = {
        r'.*DownloadStateLog: ([0-9]+) / ([0-9]+) REQ': controller.requested,
        r'.*DownloadStateLog: [0-9]+ / ([0-9]+) SCHED': controller.scheduled,
        r'.*DownloadStateLog: [0-9]+ / ([0-9]+) STARTED': controller.started,
        r'.*DownloadStateLog: [0-9]+ / ([0-9]+) STATUS:([0-9]+)': controller.status,
        r'.*DownloadStateLog: [0-9]+ / ([0-9]+) PROG:([0-9]+):([0-9]+)': controller.progress,
        r'.*DownloadStateLog: [0-9]+ / ([0-9]+) DEL': controller.deleted,
    }

    for line in sys.stdin:
        for key in router:
            m = re.search(key, line)
            if m:
                router[key](*m.groups())
                break


root = Tk()
app = Application(master=root)

log_parser = threading.Thread(target=parse_log)
log_parser.start()
app.mainloop()
