// 2018-12-24 18:07
use std::cell::RefCell;
use std::collections::VecDeque;
use std::rc::Rc;

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}

pub struct BinaryTree {}
impl BinaryTree {
    pub fn new(v: Vec<Option<i32>>) -> (Option<Rc<RefCell<TreeNode>>>) {
        if v.is_empty() {
            return None;
        }
        let mut q = VecDeque::new();
        let mut index = 0;
        let root = Rc::new(RefCell::new(TreeNode::new(v[index].unwrap())));
        q.push_back(root.clone());
        index += 1;

        while index < v.len() {
            let top = q.pop_front().unwrap();
            if let Some(x) = v[index] {
                top.borrow_mut().left = Some(Rc::new(RefCell::new(TreeNode::new(x))));
                q.push_back(top.borrow().left.clone().unwrap());
            }
            index += 1;
            if index >= v.len() {
                break;
            }
            if let Some(x) = v[index] {
                top.borrow_mut().right = Some(Rc::new(RefCell::new(TreeNode::new(x))));
                q.push_back(top.borrow().right.clone().unwrap());
            }
            index += 1;
        }
        Some(root.clone())
    }
}

#[macro_export]
macro_rules! tree {
    ($($v:tt),+,) => {
        tree!($($v),*)
    };
    (None;$vvv:expr) => {
        $vvv.push(None);
    };
    ($v:expr;$vvv:expr) => {
        $vvv.push(Some($v));
    };
    ($($v:tt),*) => {
        {
            let mut v = vec![];
            $(
                tree!($v;v);
            )*;
            BinaryTree::new(v)
        }
    };
}
