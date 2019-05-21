// Accepted
// 2019-01-05 20:35
// @lc id=693 lang=rust
// problem: binary_number_with_alternating_bits
struct Solution {}
struct BitIter(i32);
impl Iterator for BitIter {
    type Item = i32;
    fn next(&mut self) -> Option<i32> {
        if self.0 == 0 {
            return None;
        }
        let ret = self.0 & 1;
        self.0 = self.0 >> 1;
        return Some(ret);
    }
}
impl Solution {
    pub fn has_alternating_bits(n: i32) -> bool {
        let iter = BitIter(n);
        iter.collect::<Vec<_>>().windows(2).all(|w| w[0] != w[1])
    }
}

fn main() {
    println!("{:?}", Solution::has_alternating_bits(10));
    println!("{:?}", Solution::has_alternating_bits(11));
}
