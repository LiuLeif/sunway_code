// Accepted
// @lc id=110 lang=rust
// problem: balanced_binary_tree
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
type Link = Option<Rc<RefCell<TreeNode>>>;
impl Solution {
    fn dfs(root: Link) -> (bool, i32) {
        if root.is_none() {
            return (true, 0);
        }
        let root = root.unwrap();
        let (left, right) = (root.borrow().left.clone(), root.borrow().right.clone());
        let (left_balanced, left_height) = Solution::dfs(left);
        let (right_balanced, right_height) = Solution::dfs(right);
        (
            left_balanced && right_balanced && (left_height - right_height).abs() <= 1,
            std::cmp::max(left_height, right_height) + 1,
        )
    }

    pub fn is_balanced(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        let (ret, _) = Solution::dfs(root);
        ret
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::is_balanced(tree![1, 2, 2, 3, 3, None, None, 4, 4])
    );
}
