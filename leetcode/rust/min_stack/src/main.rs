// Accepted
// @lc id=155 lang=rust
// problem: min_stack
use std::cell::RefCell;
struct MinStack {
    min_stack: RefCell<Vec<i32>>,
    stack: RefCell<Vec<i32>>,
}

impl MinStack {
    fn new() -> Self {
        MinStack {
            min_stack: RefCell::new(vec![]),
            stack: RefCell::new(vec![]),
        }
    }

    fn push(&self, x: i32) {
        self.stack.borrow_mut().push(x);
        let min = *self.min_stack.borrow().last().unwrap_or(&std::i32::MAX);
        if x <= min {
            self.min_stack.borrow_mut().push(x);
        }
    }

    fn pop(&self) {
        let top = self.stack.borrow_mut().pop().unwrap();
        if *self.min_stack.borrow().last().unwrap() == top {
            self.min_stack.borrow_mut().pop();
        }
    }

    fn top(&self) -> i32 {
        *self.stack.borrow().last().unwrap()
    }

    fn get_min(&self) -> i32 {
        *self.min_stack.borrow().last().unwrap()
    }
}

fn main() {
    let stack = MinStack::new();
    stack.push(-2);
    stack.push(0);
    stack.push(-3);
    println!("{:?}", stack.get_min());
    stack.pop();
    println!("{:?}", stack.top());
    println!("{:?}", stack.get_min());
}
