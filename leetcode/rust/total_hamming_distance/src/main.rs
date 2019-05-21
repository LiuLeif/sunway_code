// Accepted
// @lc id=477 lang=rust
// problem: total_hamming_distance

struct Solution {}
impl Solution {
    pub fn total_hamming_distance(nums: Vec<i32>) -> i32 {
        let mut ret = 0;
        for i in 0..32 {
            let x = nums.iter().filter(|&n| n & (1 << i) != 0).count();
            ret += x * (nums.len() - x);
        }
        return ret as i32;
    }
}

fn main() {
    println!("{:?}", Solution::total_hamming_distance(vec![4, 14, 2]));
}
