// Accepted
// @lc id=113 lang=rust
// problem: path_sum_ii
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn path_sum(root: Option<Rc<RefCell<TreeNode>>>, sum: i32) -> Vec<Vec<i32>> {
        if root.is_none() {
            return vec![];
        }
        let root = root.unwrap();
        let (left, right, val) = (
            root.borrow().left.clone(),
            root.borrow().right.clone(),
            root.borrow().val,
        );
        if left.is_none() && right.is_none() && val == sum {
            return vec![vec![val]];
        }
        let mut ret = vec![];
        for p in Solution::path_sum(left, sum - val) {
            let mut tmp = vec![val];
            tmp.extend(p);
            ret.push(tmp);
        }
        for p in Solution::path_sum(right, sum - val) {
            let mut tmp = vec![val];
            tmp.extend(p);
            ret.push(tmp);
        }
        ret
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::path_sum(tree![5, 4, 8, 11, None, 13, 4, 7, 2], 22)
    );
}
