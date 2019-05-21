// Accepted
// @lc id=894 lang=rust
// problem: all_possible_full_binary_trees
extern crate util;
use util::binary_tree::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn all_possible_fbt(n: i32) -> Vec<Option<Rc<RefCell<TreeNode>>>> {
        if n % 2 == 0 {
            return vec![];
        }
        if n == 1 {
            return vec![Some(Rc::new(RefCell::new(TreeNode::new(0))))];
        }
        let mut ret = vec![];
        for i in (1..n - 1).step_by(2) {
            for left in Solution::all_possible_fbt(i) {
                for right in Solution::all_possible_fbt(n - 1 - i) {
                    let root = Rc::new(RefCell::new(TreeNode::new(0)));
                    root.borrow_mut().left = left.clone();
                    root.borrow_mut().right = right.clone();
                    ret.push(Some(root));
                }
            }
        }
        return ret;
    }
}                             

fn main() {
    let ret = Solution::all_possible_fbt(7);
    println!("{:?}", ret.len());
}
