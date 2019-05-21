// Accepted
// @lc id=508 lang=rust
// problem: most_frequent_subtree_sum
extern crate util;
use util::binary_tree::*;
struct Solution {}

use std::cell::RefCell;
use std::collections::HashMap;
use std::rc::Rc;
type Link = Option<Rc<RefCell<TreeNode>>>;
impl Solution {
    fn dfs(root: Link, counter: &mut HashMap<i32, i32>) -> i32 {
        if root.is_none() {
            return 0;
        }
        let root = root.unwrap();
        let left = root.borrow().left.clone();
        let right = root.borrow().right.clone();
        let v = Solution::dfs(left, counter) + Solution::dfs(right, counter) + root.borrow().val;
        *counter.entry(v).or_insert(0) += 1;
        v
    }

    pub fn find_frequent_tree_sum(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut counter = HashMap::<i32, i32>::new();
        Solution::dfs(root, &mut counter);
        let max_value = counter.values().max().unwrap_or(&0);
        counter
            .iter()
            .filter(|(k, v)| *v == max_value)
            .map(|(k, v)| *k)
            .collect()
    }
}

fn main() {
    let root = BinaryTree::new(vec![]);
    println!("{:?}", Solution::find_frequent_tree_sum(root));
}
