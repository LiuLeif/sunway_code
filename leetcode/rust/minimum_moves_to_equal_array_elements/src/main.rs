// Accepted
// @lc id=453 lang=rust
// problem: minimum_moves_to_equal_array_elements
struct Solution {}
impl Solution {
    pub fn min_moves(nums: Vec<i32>) -> i32 {
        let min = nums.iter().min().unwrap();
        nums.iter().map(|x| x - min).sum()
    }
}

fn main() {
    println!("{:?}", Solution::min_moves(vec![1, 1, 2147483647]));
    println!("{:?}", Solution::min_moves(vec![1, 2, 3]));
}
