// Accepted
// 2019-01-04 11:11
// @lc id=872 lang=rust
// problem: leaf_similar_trees
extern crate util;
use util::binary_tree::*;

struct Solution {}
use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn leaf_similar(
        root1: Option<Rc<RefCell<TreeNode>>>,
        root2: Option<Rc<RefCell<TreeNode>>>,
    ) -> bool {
        fn traverse(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
            let mut ret = vec![];
            match root {
                Some(node) => {
                    if node.borrow().left == None && node.borrow().right == None {
                        return vec![node.borrow().val];
                    }
                    if node.borrow().left != None {
                        ret.extend(traverse(node.borrow().left.clone()))
                    }
                    if node.borrow().right != None {
                        ret.extend(traverse(node.borrow().right.clone()))
                    }
                }
                None => (),
            }
            return ret;
        }

        return traverse(root1) == traverse(root2);
    }
}

fn main() {
    let root1 = BinaryTree::new(vec![
        Some(3),
        Some(5),
        Some(1),
        Some(6),
        Some(2),
        Some(9),
        Some(8),
        None,
        None,
        Some(7),
        Some(4),
    ]);
    let root2 = BinaryTree::new(vec![
        Some(3),
        Some(5),
        Some(1),
        Some(6),
        Some(7),
        Some(4),
        Some(2),
        None,
        None,
        None,
        None,
        None,
        None,
        Some(9),
        Some(8),
    ]);
    println!("{:?}", Solution::leaf_similar(root1, root2));
}
