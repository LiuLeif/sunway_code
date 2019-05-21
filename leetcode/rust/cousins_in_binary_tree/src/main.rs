// Accepted
// @lc id=993 lang=rust
// problem: cousins_in_binary_tree
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
type Link = Option<Rc<RefCell<TreeNode>>>;

impl Solution {
    fn get_path(root: Link, x: i32) -> Vec<i32> {
        if root.is_none() {
            return vec![];
        }
        let root = root.unwrap();
        let (left, right, val) = (
            root.borrow().left.clone(),
            root.borrow().right.clone(),
            root.borrow().val,
        );
        let mut ret = vec![val];
        if val == x {
            return ret;
        }
        let left_path = Solution::get_path(left, x);
        if !left_path.is_empty() {
            ret.extend(left_path);
            return ret;
        }
        let right_path = Solution::get_path(right, x);
        if !right_path.is_empty() {
            ret.extend(right_path);
            return ret;
        }
        return vec![];
    }

    pub fn is_cousins(root: Option<Rc<RefCell<TreeNode>>>, x: i32, y: i32) -> bool {
        let path_x = Solution::get_path(root.clone(), x);
        let path_y = Solution::get_path(root.clone(), y);
        let (lx, ly) = (path_x.len(), path_y.len());
        if lx != ly || lx < 2 {
            return false;
        }
        path_x[lx - 2] != path_y[lx - 2]
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::is_cousins(tree![1, 2, 3, None, 4, None, 5], 5, 4)
    );
    println!("{:?}", Solution::is_cousins(tree![1, 2, 3, None, 4], 2, 3));
}
