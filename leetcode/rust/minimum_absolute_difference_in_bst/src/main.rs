// Accepted
// @lc id=530 lang=rust
// problem: minimum_absolute_difference_in_bst
extern crate util;
use util::binary_tree::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
type Link = Option<Rc<RefCell<TreeNode>>>;
impl Solution {
    fn dfs(root: Link) -> (i32, i32, i32) {
        if root.is_none() {
            return (0, 0, 0);
        }
        let root = root.unwrap();

        let left = root.borrow().left.clone();
        let right = root.borrow().right.clone();

        let root_value = root.borrow().val;
        let (mut max, mut min, mut result) = (root_value, root_value, std::i32::MAX);
        if left.is_some() {
            let t = Solution::dfs(left);
            result = std::cmp::min(result, root_value - t.1);
            result = std::cmp::min(result, t.2);
            min = t.0;
        }
        if right.is_some() {
            let t = Solution::dfs(right);
            result = std::cmp::min(result, t.0 - root_value);
            result = std::cmp::min(result, t.2);
            max = t.1;
        }
        return (min, max, result);
    }

    pub fn get_minimum_difference(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let (_, _, ret) = Solution::dfs(root);
        ret
    }
}

fn main() {
    let root = BinaryTree::new(vec![Some(1), None, Some(3), Some(2)]);
    println!("{:?}", Solution::get_minimum_difference(root));
}
