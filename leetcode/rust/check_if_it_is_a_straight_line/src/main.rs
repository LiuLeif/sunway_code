// Accepted
struct Solution {}
/*
 * @lc app=leetcode id=1232 lang=rust
 *
 * [1232] Check If It Is a Straight Line
 */

// @lc code=start
impl Solution {
    pub fn check_straight_line(coordinates: Vec<Vec<i32>>) -> bool {
        let (mut prev_x, mut prev_y) = (0, 0);
        for w in coordinates.windows(2) {
            let (x, y) = (&w[0], &w[1]);

            let delta_x = x[0] - y[0];
            let delta_y = x[1] - y[1];

            if (delta_x == 0 && delta_y == 0) {
                continue;
            }

            if (prev_x == 0 && prev_y == 0) {
                prev_x = delta_x;
                prev_y = delta_y;
                continue;
            }

            if (delta_x * prev_y != delta_y * prev_x) {
                return false;
            }
        }
        return true;
    }
}
// @lc code=end

fn main() {
    println!(
        "{:?}",
        Solution::check_straight_line(vec![
            vec![1, 2],
            vec![2, 3],
            vec![3, 4],
            vec![4, 5],
            vec![5, 6],
            vec![6, 7]
        ])
    );

    println!(
        "{:?}",
        Solution::check_straight_line(vec![vec![1, 2], vec![0, 2], vec![2, 2],])
    );
}
