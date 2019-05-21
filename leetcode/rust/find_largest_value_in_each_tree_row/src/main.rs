// Accepted
// @lc id=515 lang=rust
// problem: find_largest_value_in_each_tree_row
extern crate util;
use std::collections::VecDeque;
use util::binary_tree::*;
struct Solution {}
use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn largest_values(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut ret = vec![];
        if root == None {
            return ret;
        }
        let mut queue = VecDeque::new();
        queue.push_back(root.clone());
        queue.push_back(None);
        let mut curr_max = std::i32::MIN;
        while queue.len() != 0 {
            let top = queue.pop_front().unwrap();
            if top == None {
                ret.push(curr_max);
                curr_max = std::i32::MIN;
                if queue.len() != 0 {
                    queue.push_back(None);
                }
                continue;
            }
            let node = top.unwrap();
            curr_max = std::cmp::max(curr_max, node.borrow().val);
            let left = node.borrow().left.clone();
            let right = node.borrow().right.clone();
            if left != None {
                queue.push_back(left);
            }
            if right != None {
                queue.push_back(right);
            }
        }
        return ret;
    }
}

fn main() {
    let root = BinaryTree::new(vec![Some(1), Some(-1)]);
    println!("{:?}", Solution::largest_values(root));
}
