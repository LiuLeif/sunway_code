// Accepted
// @lc id=200 lang=rust
// problem: number_of_islands
struct Solution {}
impl Solution {
    fn dfs(grid: &mut Vec<Vec<char>>, x: usize, y: usize) {
        if grid[x][y] == '0' {
            return;
        }
        grid[x][y] = '0';

        let (m, n) = (grid.len() as i32, grid[0].len() as i32);
        for delta_x in vec![-1, 0, 1] {
            for delta_y in vec![-1, 0, 1] {
                if (delta_x == 0 && delta_y == 0) || (delta_x != 0 && delta_y != 0) {
                    continue;
                }
                let (x, y) = (x as i32 + delta_x, y as i32 + delta_y);
                if x < 0 || x >= m || y < 0 || y >= n {
                    continue;
                }
                Solution::dfs(grid, x as usize, y as usize);
            }
        }
    }

    pub fn num_islands(grid: Vec<Vec<char>>) -> i32 {
        if grid.is_empty() {
            return 0;
        }
        let mut grid = grid;
        let (m, n) = (grid.len(), grid[0].len());
        let mut ret = 0;
        for i in 0..m {
            for j in 0..n {
                if grid[i][j] == '0' {
                    continue;
                }
                Solution::dfs(&mut grid, i, j);
                ret += 1;
            }
        }
        ret
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::num_islands(vec![
            vec!['1', '1', '0', '0', '0'],
            vec!['1', '1', '0', '0', '0'],
            vec!['0', '0', '1', '0', '0'],
            vec!['0', '0', '0', '1', '1'],
        ])
    );
}
