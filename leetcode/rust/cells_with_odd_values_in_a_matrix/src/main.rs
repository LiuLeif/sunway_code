// Accepted
struct Solution {}
/*
 * @lc app=leetcode id=1252 lang=rust
 *
 * [1252] Cells with Odd Values in a Matrix
 */

// @lc code=start
impl Solution {
    pub fn odd_cells(n: i32, m: i32, indices: Vec<Vec<i32>>) -> i32 {
        let mut rows = vec![0; n as usize];
        let mut cols = vec![0; m as usize];

        indices.iter().for_each(|p| {
            rows[p[0] as usize] += 1;
            cols[p[1] as usize] += 1;
        });

        let row = rows.into_iter().filter(|i| i % 2 == 1).count() as i32;
        let col = cols.into_iter().filter(|i| i % 2 == 1).count() as i32;

        return row * m + col * n - row * col * 2;
    }
}
// @lc code=end

fn main() {
    println!(
        "{:?}",
        Solution::odd_cells(2, 2, vec![vec![1, 1], vec![0, 0]])
    );
}
