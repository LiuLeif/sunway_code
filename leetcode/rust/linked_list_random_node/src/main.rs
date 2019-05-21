// Accepted
// @lc id=382 lang=rust
// problem: linked_list_random_node
use rand::prelude::*;
use std::cell::RefCell;
use util::linked_list::*;
use util::*;

struct Solution {
    head: Option<Box<ListNode>>,
    rng: RefCell<ThreadRng>,
}

impl Solution {
    fn new(head: Option<Box<ListNode>>) -> Self {
        Solution {
            head,
            rng: RefCell::new(thread_rng()),
        }
    }

    fn get_random(&self) -> i32 {
        let mut head = self.head.as_ref();
        let mut n = 0;
        let mut ret = 0;
        while let Some(node) = head {
            n += 1;
            if self.rng.borrow_mut().gen_range(0, n) == 0 {
                ret = node.val;
            }
            head = node.next.as_ref();
        }
        return ret;
    }
}

fn main() {
    let solution = Solution::new(list![1, 2, 3, 4]);
    for i in 0..10 {
        println!("{:?}", solution.get_random());
    }
}
