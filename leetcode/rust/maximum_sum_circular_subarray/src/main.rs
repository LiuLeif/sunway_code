// Accepted
// @lc id=918 lang=rust
// problem: maximum_sum_circular_subarray
struct Solution {}
impl Solution {
    pub fn max_subarray_sum_circular(a: Vec<i32>) -> i32 {
        // 5 -3 5 5 -3 5
        let (mut currMax, mut max) = (0, std::i32::MIN);
        let (mut currMin, mut min) = (0, std::i32::MAX);
        let mut total = 0;
        for i in 0..a.len() {
            currMax = std::cmp::max(currMax + a[i], a[i]);
            max = std::cmp::max(max, currMax);
            currMin = std::cmp::min(currMin + a[i], a[i]);
            min = std::cmp::min(min, currMin);
            total += a[i];
        }
        if max < 0 {
            return max;
        }
        std::cmp::max(max, total - min)
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::max_subarray_sum_circular(vec![-2, -3, -1])
    );
}
