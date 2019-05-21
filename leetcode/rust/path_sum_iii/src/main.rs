// Accepted
// @lc id=437 lang=rust
// problem: path_sum_iii
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    fn dfs(root: Option<Rc<RefCell<TreeNode>>>, sum: i32) -> i32 {
        if root.is_none() {
            return 0;
        }
        let root = root.unwrap();
        let (left, right, val) = (
            root.borrow().left.clone(),
            root.borrow().right.clone(),
            root.borrow().val,
        );
        let mut ret = 0;
        if sum == val {
            ret = 1;
        }
        ret + Solution::dfs(left, sum - val) + Solution::dfs(right, sum - val)
    }

    pub fn path_sum(root: Option<Rc<RefCell<TreeNode>>>, sum: i32) -> i32 {
        let ret = Solution::dfs(root.clone(), sum);
        if root.is_none() {
            return ret;
        }
        let root = root.unwrap();
        let (left, right) = (root.borrow().left.clone(), root.borrow().right.clone());
        ret + Solution::path_sum(left.clone(), sum) + Solution::path_sum(right.clone(), sum)
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::path_sum(tree![10, 5, (-3), 3, 2, None, 11, 3, (-2), None, 1], 8)
    );
    println!("{:?}", Solution::path_sum(tree![10, 5], 5));
}
