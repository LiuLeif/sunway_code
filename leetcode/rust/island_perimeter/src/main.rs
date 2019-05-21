// 2019-01-02 22:29
// @lc id=463 lang=rust
// problem: island_perimeter
struct Solution {}
impl Solution {
    pub fn island_perimeter(grid: Vec<Vec<i32>>) -> i32 {
        let mut ret = 0;
        let (m, n) = (grid.len(), grid[0].len());
        for i in 0..m {
            for j in 0..n {
                if grid[i][j] == 0 {
                    continue;
                }
                ret += 4;
                if i > 0 && grid[i - 1][j] == 1 {
                    ret -= 2;
                }
                if j > 0 && grid[i][j - 1] == 1 {
                    ret -= 2;
                }
            }
        }
        return ret;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::island_perimeter(vec![
            vec![0, 1, 0, 0],
            vec![1, 1, 1, 0],
            vec![0, 1, 0, 0],
            vec![1, 1, 0, 0],
        ])
    );
}
