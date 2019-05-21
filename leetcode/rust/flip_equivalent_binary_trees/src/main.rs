// Accepted
// @lc id=951 lang=rust
// problem: flip_equivalent_binary_trees
extern crate util;
use std::collections::HashMap;
use util::binary_tree::*;
struct Solution {}
use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn flip_equiv(
        root1: Option<Rc<RefCell<TreeNode>>>,
        root2: Option<Rc<RefCell<TreeNode>>>,
    ) -> bool {
        if root1 == None && root2 == None {
            return true;
        }
        if root1 == None || root2 == None {
            return false;
        }
        let root1 = root1.unwrap();
        let root2 = root2.unwrap();

        if root1.borrow().val != root2.borrow().val {
            return false;
        }

        let l1 = root1.borrow().left.clone();
        let r1 = root1.borrow().right.clone();

        let l2 = root2.borrow().left.clone();
        let r2 = root2.borrow().right.clone();

        let a = Solution::flip_equiv(l1.clone(), l2.clone());
        let b = Solution::flip_equiv(r1.clone(), r2.clone());

        let c = Solution::flip_equiv(l2.clone(), r1.clone());
        let d = Solution::flip_equiv(l1.clone(), r2.clone());

        return (a && b) || (c && d);
    }
}

fn main() {
    let root1 = BinaryTree::new(vec![
        Some(1),
        Some(2),
        Some(3),
        Some(4),
        Some(5),
        Some(6),
        None,
        None,
        None,
        Some(7),
        Some(8),
    ]);
    let root2 = BinaryTree::new(vec![
        Some(1),
        Some(3),
        Some(2),
        None,
        Some(6),
        Some(4),
        Some(5),
        None,
        None,
        None,
        None,
        Some(8),
        Some(7),
    ]);
    println!("{:?}", Solution::flip_equiv(root1, root2));
}
