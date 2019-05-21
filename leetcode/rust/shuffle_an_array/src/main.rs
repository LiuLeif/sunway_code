// Accepted
// @lc id=384 lang=rust
// problem: shuffle_an_array
extern crate rand;
use rand::prelude::*;
use std::cell::RefCell;

struct Solution {
    v: Vec<i32>,
    rng: RefCell<ThreadRng>,
}

use rand::Rng;

/** You can modify self type for your need. */
impl Solution {
    fn new(nums: Vec<i32>) -> Self {
        return Solution {
            v: nums,
            rng: RefCell::new(thread_rng()),
        };
    }

    fn reset(&self) -> Vec<i32> {
        return self.v.clone();
    }

    fn shuffle(&self) -> Vec<i32> {
        let mut ret = self.v.clone();
        for i in 0..ret.len() {
            let index = self.rng.borrow_mut().gen_range(0, i + 1);
            ret.swap(i, index);
        }
        return ret;
    }
}

fn main() {
    let x = Solution::new(vec![1, 2, 3, 4]);
    println!("{:?}", x.shuffle());
    println!("{:?}", x.shuffle());
}
