// Accepted
// @lc id=398 lang=rust
// problem: random_pick_index
use rand::prelude::*;
use std::cell::RefCell;

struct Solution {
    nums: Vec<i32>,
    rng: RefCell<ThreadRng>,
}

/** You can modify self type for your need. */
impl Solution {
    fn new(nums: Vec<i32>) -> Self {
        Solution {
            nums: nums,
            rng: RefCell::new(thread_rng()),
        }
    }
    fn pick(&self, target: i32) -> i32 {
        self.nums
            .iter()
            .enumerate()
            .filter(|&(i, &v)| v == target)
            .fold((0, 0), |(total, index), (i, _)| {
                if self.rng.borrow_mut().gen_range(0, total + 1) == 0 {
                    (total + 1, i)
                } else {
                    (total + 1, index)
                }
            })
            .1 as i32
    }
}

fn main() {
    let picker = Solution::new(vec![1, 2, 3, 3, 3]);
    println!("{:?}", picker.pick(1));
}
