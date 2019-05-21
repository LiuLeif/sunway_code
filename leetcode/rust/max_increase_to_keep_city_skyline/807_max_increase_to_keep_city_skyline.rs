// 2018-12-24 15:29
// @lc id=807 lang=rust
struct Solution {}
impl Solution {
    pub fn max_increase_keeping_skyline(grid: Vec<Vec<i32>>) -> i32 {
        let mut row_max = vec![0; grid.len()];
        let mut col_max = vec![0; grid[0].len()];
        for i in 0..grid.len() {
            for j in 0..grid[0].len() {
                if grid[i][j] > row_max[i] {
                    row_max[i] = grid[i][j]
                }
                if grid[i][j] > col_max[j] {
                    col_max[j] = grid[i][j]
                }
            }
        }
        let mut ret = 0;
        for i in 0..grid.len() {
            for j in 0..grid[0].len() {
                let min = std::cmp::min(row_max[i], col_max[j]);
                if grid[i][j] < min {
                    ret += min - grid[i][j]
                }
            }
        }
        return ret;
    }
}

fn main() {
    let grid = vec![
        vec![3, 0, 8, 4],
        vec![2, 4, 5, 7],
        vec![9, 2, 6, 3],
        vec![0, 3, 1, 0],
    ];
    println!("{:?}", Solution::max_increase_keeping_skyline(grid));
}
