// Accepted
// @lc id=129 lang=rust
// problem: sum_root_to_leaf_numbers
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
type Link = Option<Rc<RefCell<TreeNode>>>;
impl Solution {
    fn dfs(root: Link, sum: i32) -> i32 {
        if root.is_none() {
            return 0;
        }
        let root = root.unwrap();
        let (left, right, val) = (
            root.borrow().left.clone(),
            root.borrow().right.clone(),
            root.borrow().val,
        );
        if left.is_none() && right.is_none() {
            return sum * 10 + val;
        }
        Solution::dfs(left, sum * 10 + val) + Solution::dfs(right, sum * 10 + val)
    }

    pub fn sum_numbers(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        Solution::dfs(root, 0)
    }
}

fn main() {
    println!("{:?}", Solution::sum_numbers(tree![4, 9, 0, 5, 1]));
}
