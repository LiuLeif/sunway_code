// Accepted
// @lc id=2 lang=rust
// problem: add_two_numbers
use util::linked_list::*;
use util::*;

struct Solution {}

impl Solution {
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut fake_head = ListNode::new(0);
        let mut prev = &mut fake_head;
        let (mut l1, mut l2) = (l1.as_ref(), l2.as_ref());
        let mut carry = 0;
        while l1.is_some() || l2.is_some() {
            let sum = l1.map_or(0, |node| node.val) + l2.map_or(0, |node| node.val) + carry;
            l1 = l1.and_then(|node| node.next.as_ref());
            l2 = l2.and_then(|node| node.next.as_ref());

            carry = sum / 10;
            prev.next = Some(Box::new(ListNode::new(sum % 10)));
            prev = prev.next.as_mut().unwrap();
        }
        if carry != 0 {
            prev.next = Some(Box::new(ListNode::new(carry)));
        }
        fake_head.next
    }
}

fn main() {
    let l1 = list![5];
    let l2 = list![5];
    println!("{:?}", Solution::add_two_numbers(l1, l2));
}
