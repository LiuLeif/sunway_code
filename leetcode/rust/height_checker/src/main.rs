// Accepted
// @lc id=1051 lang=rust
// problem: height_checker
struct Solution {}
impl Solution {
    pub fn height_checker(heights: Vec<i32>) -> i32 {
        let mut sorted_heights = heights.clone();
        sorted_heights.sort();

        heights
            .into_iter()
            .zip(sorted_heights.into_iter())
            .map(|(a, b)| if a == b { 0 } else { 1 })
            .sum()
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::height_checker(vec![1, 2, 1, 2, 1, 1, 1, 2, 1])
    );
}
