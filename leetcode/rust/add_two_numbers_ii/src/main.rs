// Accepted
// @lc id=445 lang=rust
// problem: add_two_numbers_ii
use util::linked_list::*;
use util::*;

struct Solution {}

impl Solution {
    fn add(l1: Option<Box<ListNode>>, l2: Option<Box<ListNode>>) -> (Option<Box<ListNode>>, i32) {
        if l1.is_none() {
            return (None, 0);
        }
        let (l1, l2) = (l1.unwrap(), l2.unwrap());
        let (node, carry) = Solution::add(l1.next, l2.next);
        let sum = l1.val + l2.val + carry;
        let mut ret = ListNode::new(sum % 10);
        ret.next = node;

        (Some(Box::new(ret)), sum / 10)
    }

    fn length(head: Option<&Box<ListNode>>) -> i32 {
        let mut ret = 0;
        let mut head = head;
        while head.is_some() {
            ret += 1;
            head = head.unwrap().next.as_ref();
        }
        ret
    }

    fn append_zero(head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
        let mut fake_head = ListNode::new(0);
        let mut prev = &mut fake_head;
        for _ in 0..n {
            prev.next = Some(Box::new(ListNode::new(0)));
            prev = prev.next.as_mut().unwrap();
        }
        prev.next = head;
        fake_head.next
    }

    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let l1_len = Solution::length(l1.as_ref());
        let l2_len = Solution::length(l2.as_ref());
        let l1 = {
            if l1_len >= l2_len {
                l1
            } else {
                Solution::append_zero(l1, l2_len - l1_len)
            }
        };
        let l2 = {
            if l2_len >= l1_len {
                l2
            } else {
                Solution::append_zero(l2, l1_len - l2_len)
            }
        };
        let (node, carry) = Solution::add(l1, l2);
        if carry == 0 {
            return node;
        }
        let mut ret = ListNode::new(carry);
        ret.next = node;
        Some(Box::new(ret))
    }
}

fn main() {
    println!("{:?}", Solution::add_two_numbers(list![10], list![1]));
}
