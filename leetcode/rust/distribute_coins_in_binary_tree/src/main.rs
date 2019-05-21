// Accepted
// @lc id=979 lang=rust
// problem: distribute_coins_in_binary_tree
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
type Link = Option<Rc<RefCell<TreeNode>>>;
impl Solution {
    fn dfs(root: Link) -> (i32, i32) {
        if root.is_none() {
            return (0, 0);
        }
        let root = root.unwrap();
        let (left, right, root_val) = (
            root.borrow().left.clone(),
            root.borrow().right.clone(),
            root.borrow().val,
        );

        let (left_moves, left_delta) = Solution::dfs(left);
        let (right_moves, right_delta) = Solution::dfs(right);

        let moves = left_moves + right_moves + left_delta.abs() + right_delta.abs();
        let delta = left_delta + right_delta + root_val - 1;
        (moves, delta)
    }

    pub fn distribute_coins(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let (moves, _) = Solution::dfs(root);
        moves
    }
}

fn main() {
    println!("{:?}", Solution::distribute_coins(tree![0, 2, 0]));
}
