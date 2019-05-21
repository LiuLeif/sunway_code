// Accepted
// @lc id=230 lang=rust
// problem: kth_smallest_element_in_a_bst
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
type Link = Option<Rc<RefCell<TreeNode>>>;

enum Ret {
    Result(i32),
    Count(i32),
}

impl Solution {
    fn dfs(root: Link, k: i32) -> Ret {
        if root.is_none() {
            return Ret::Count(0);
        }
        let root = root.unwrap();
        let (left, right, val) = (
            root.borrow().left.clone(),
            root.borrow().right.clone(),
            root.borrow().val,
        );
        let mut total_count = 1;
        let count = {
            match Solution::dfs(left, k) {
                Ret::Result(ret) => return Ret::Result(ret),
                Ret::Count(count) => count,
            }
        };
        if count == k - 1 {
            return Ret::Result(val);
        }
        total_count += count;
        match Solution::dfs(right, k - count - 1) {
            Ret::Result(ret) => return Ret::Result(ret),
            Ret::Count(count) => total_count += count,
        }
        Ret::Count(total_count)
    }

    pub fn kth_smallest(root: Option<Rc<RefCell<TreeNode>>>, k: i32) -> i32 {
        match Solution::dfs(root, k) {
            Ret::Result(ret) => ret,
            _ => 0,
        }
    }
}

fn main() {
    println!("{:?}", Solution::kth_smallest(tree![3, 1, 4, None, 2], 4));
}
