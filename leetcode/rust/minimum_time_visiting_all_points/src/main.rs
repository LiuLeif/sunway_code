// Accepted
struct Solution {}
/*
 * @lc app=leetcode id=1266 lang=rust
 *
 * [1266] Minimum Time Visiting All Points
 */

// @lc code=start
impl Solution {
    pub fn min_time_to_visit_all_points(points: Vec<Vec<i32>>) -> i32 {
        points
            .windows(2)
            .map(|p| {
                let delta_x = (p[0][0] - p[1][0]).abs();
                let delta_y = (p[0][1] - p[1][1]).abs();
                std::cmp::max(delta_x, delta_y)
            })
            .sum()
    }
}
// @lc code=end

fn main() {
    println!(
        "{:?}",
        Solution::min_time_to_visit_all_points(vec![vec![1, 1], vec![3, 4], vec![-1, 0]])
    );
}
