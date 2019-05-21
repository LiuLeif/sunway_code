// 2018-12-29 22:50
// @lc id=876 lang=rust
extern crate util;
use util::linked_list::{LinkedList, ListNode};

struct Solution {}
impl Solution {
    pub fn middle_node(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut curr = head.as_ref();
        let mut count = 0;
        while let Some(ref node) = curr {
            curr = node.next.as_ref();
            count += 1;
        }
        let mut head = head;
        for _i in 0..count / 2 {
            head = head.unwrap().next;
        }
        return head;
    }
}

fn main() {
    let head = LinkedList::new(vec![1, 2, 3, 4, 5]);
    println!("{:?}", Solution::middle_node(head));
}
