// Accepted
// @lc id=865 lang=rust
// problem: smallest_subtree_with_all_the_deepest_nodes
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
type Link = Option<Rc<RefCell<TreeNode>>>;
impl Solution {
    fn dfs(root: Link) -> (Link, i32) {
        if root.is_none() {
            return (None, 0);
        }
        let root = root.unwrap();
        let (left, right) = (root.borrow().left.clone(), root.borrow().right.clone());
        let (left_node, left_depth) = Solution::dfs(left);
        let (right_node, right_depth) = Solution::dfs(right);

        if left_depth == right_depth {
            return (Some(root), left_depth + 1);
        }
        if left_depth > right_depth {
            return (left_node, left_depth + 1);
        }
        return (right_node, right_depth + 1);
    }

    pub fn subtree_with_all_deepest(
        root: Option<Rc<RefCell<TreeNode>>>,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        Solution::dfs(root).0
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::subtree_with_all_deepest(tree![3, 5, 1, 6, 2, 0, 8, None, None, 7, 4])
    );
}
