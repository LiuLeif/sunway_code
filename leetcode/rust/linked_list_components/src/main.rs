// Accepted
// @lc id=817 lang=rust
// problem: linked_list_components
use util::linked_list::*;
use util::*;

use std::collections::HashSet;
struct Solution {}

impl Solution {
    pub fn num_components(head: Option<Box<ListNode>>, g: Vec<i32>) -> i32 {
        let set = g.into_iter().collect::<HashSet<i32>>();

        let mut head = head;
        let mut ret = 0;
        let mut is_boundary = true;

        while let Some(node) = head {
            if !set.contains(&node.val) {
                is_boundary = true;
            } else {
                if is_boundary {
                    ret += 1;
                }
                is_boundary = false;
            }
            head = node.next;
        }
        return ret;
    }
}

fn main() {
    let head = list!(0, 1, 2, 3, 4);
    println!("{:?}", Solution::num_components(head, vec![0, 2, 4]));
}
