// Accepted
// 2019-01-07 00:47
// @lc id=896 lang=rust
// problem: monotonic_array
struct Solution {}
impl Solution {
    pub fn is_monotonic(a: Vec<i32>) -> bool {
        a.windows(2).all(|w| w[0] >= w[1]) || a.windows(2).all(|w| w[0] <= w[1])
    }
}

fn main() {
    println!("{:?}", Solution::is_monotonic(vec![1, 2, 2, 3]));
    println!("{:?}", Solution::is_monotonic(vec![3, 2, 2, 1]));
    println!("{:?}", Solution::is_monotonic(vec![1, 3, 2]));
}
