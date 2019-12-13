// Accepted
use util::linked_list::*;
use util::*;

struct Solution {}
/*
 * @lc app=leetcode id=203 lang=rust
 *
 * [203] Remove Linked List Elements
 */

impl Solution {
    pub fn remove_elements(head: Option<Box<ListNode>>, val: i32) -> Option<Box<ListNode>> {
        let mut fake_head = ListNode::new(0);
        let mut prev = &mut fake_head;
        let mut head = head;
        while let Some(mut node) = head {
            head = node.next.take();
            if node.val != val {
                prev.next = Some(node);
                prev = prev.next.as_mut().unwrap();
            }
        }
        prev.next = None;
        return fake_head.next;
    }
}
// @lc code=end

fn main() {
    let head = list!(1, 2, 3, 4);
    println!("{:?}", Solution::remove_elements(head, 3));
}
