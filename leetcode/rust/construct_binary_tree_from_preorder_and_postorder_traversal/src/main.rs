// Accepted
// @lc id=889 lang=rust
// problem: construct_binary_tree_from_preorder_and_postorder_traversal
use util::binary_tree::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
type Link = Option<Rc<RefCell<TreeNode>>>;
impl Solution {
    fn dfs(pre: &[i32], post: &[i32]) -> Link {
        if pre.is_empty() {
            return None;
        }
        let root = Rc::new(RefCell::new(TreeNode::new(pre[0])));
        let n = pre.len();
        if n == 1 {
            return Some(root);
        }
        let key = pre[1];
        let index = post.iter().position(|&x| x == key).unwrap() + 1;
        root.borrow_mut().left = Solution::dfs(&pre[1..=index], &post[..index]);
        root.borrow_mut().right = Solution::dfs(&pre[index + 1..], &post[index..n - 1]);
        Some(root)
    }

    pub fn construct_from_pre_post(pre: Vec<i32>, post: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
        Solution::dfs(pre.as_slice(), post.as_slice())
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::construct_from_pre_post(vec![1, 2, 4, 5, 3, 6, 7], vec![4, 5, 2, 6, 7, 3, 1])
    );
}
