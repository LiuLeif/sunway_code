// Accepted
// @lc id=173 lang=rust
// problem: binary_search_tree_iterator
use std::cell::RefCell;
use std::rc::Rc;
use util::binary_tree::*;
use util::*;

type Link = Option<Rc<RefCell<TreeNode>>>;

struct BSTIterator {
    stack: RefCell<Vec<Link>>,
}

impl BSTIterator {
    fn new(root: Option<Rc<RefCell<TreeNode>>>) -> Self {
        let mut stack = vec![];
        if root.is_some() {
            stack.push(root);
        }
        BSTIterator {
            stack: RefCell::new(stack),
        }
    }

    fn next(&self) -> i32 {
        while !self.stack.borrow().is_empty() {
            let top = self.stack.borrow_mut().pop().unwrap().unwrap();
            let (left, right, val) = (
                top.borrow().left.clone(),
                top.borrow().right.clone(),
                top.borrow().val,
            );
            if left.is_none() && right.is_none() {
                return val;
            }
            if right.is_some() {
                self.stack.borrow_mut().push(right);
            }
            top.borrow_mut().left = None;
            top.borrow_mut().right = None;
            self.stack.borrow_mut().push(Some(top));
            if left.is_some() {
                self.stack.borrow_mut().push(left);
            }
        }
        unreachable!();
    }

    fn has_next(&self) -> bool {
        !self.stack.borrow().is_empty()
    }
}

fn main() {
    let iter = BSTIterator::new(tree![2, 1, 3]);
    println!("{:?}", iter.next());
    println!("{:?}", iter.next());
    println!("{:?}", iter.next());
    println!("{:?}", iter.has_next());
}
