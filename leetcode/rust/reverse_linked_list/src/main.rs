// Accepted
// @lc id=206 lang=rust
// problem: reverse_linked_list
extern crate util;
use util::linked_list::*;
struct Solution {}

impl Solution {
    pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut prev: Option<Box<ListNode>> = None;
        let mut curr = head;
        while let Some(mut curr_node) = curr {
            let next = curr_node.next;
            curr_node.next = prev;
            prev = Some(curr_node);
            curr = next;
        }
        prev
    }
}

fn main() {
    let head = LinkedList::new(vec![1, 2, 3, 4]);
    println!("{:?}", Solution::reverse_list(head));
}
