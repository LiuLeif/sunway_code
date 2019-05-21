// Accepted
// @lc id=912 lang=rust
// problem: sort_an_array
struct Solution {}
impl Solution {
    pub fn sort_array(nums: Vec<i32>) -> Vec<i32> {
        let mut ret = nums.clone();
        ret.sort();
        ret
    }
}

fn main() {
    println!("{:?}", Solution::sort_array(vec![5, 2, 3, 1]));
}
