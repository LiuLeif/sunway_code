// Accepted
// @lc id=695 lang=rust
// problem: max_area_of_island
struct Solution {}
impl Solution {
    fn dfs(x: i32, y: i32, grid: &mut Vec<Vec<i32>>) -> i32 {
        let dir = vec![(0, 1), (0, -1), (-1, 0), (1, 0)];
        let mut ret = 1;
        for d in dir {
            let (x, y) = (x + d.0, y + d.1);
            if x < 0 || x >= grid.len() as i32 || y < 0 || y >= grid[0].len() as i32 {
                continue;
            }
            if grid[x as usize][y as usize] == 0 {
                continue;
            }
            grid[x as usize][y as usize] = 0;
            ret += Solution::dfs(x, y, grid);
        }
        ret
    }

    pub fn max_area_of_island(grid: Vec<Vec<i32>>) -> i32 {
        let (m, n) = (grid.len(), grid[0].len());
        let mut ret = 0;
        let mut grid = grid;
        for i in 0..m {
            for j in 0..n {
                if grid[i][j] == 0 {
                    continue;
                }
                grid[i][j] = 0;
                ret = std::cmp::max(ret, Solution::dfs(i as i32, j as i32, &mut grid))
            }
        }
        return ret;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::max_area_of_island(vec![
            vec![0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0],
            vec![0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0],
            vec![0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0],
            vec![0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0],
            vec![0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0],
            vec![0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0],
            vec![0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0],
            vec![0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0],
        ])
    );
}
