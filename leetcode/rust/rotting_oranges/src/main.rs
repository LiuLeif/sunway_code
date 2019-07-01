// Accepted
// @lc id=994 lang=rust
// problem: rotting_oranges
use std::collections::VecDeque;

struct Solution {}
impl Solution {
    pub fn oranges_rotting(grid: Vec<Vec<i32>>) -> i32 {
        if grid.is_empty() {
            return 0;
        }
        let (m, n) = (grid.len() as i32, grid[0].len() as i32);
        let mut grid = grid;
        let mut queue = VecDeque::new();
        let mut total_one = 0;
        for i in 0..grid.len() {
            for j in 0..grid[0].len() {
                if grid[i][j] == 2 {
                    queue.push_back((i as i32, j as i32));
                }
                if grid[i][j] == 1 {
                    total_one += 1;
                }
            }
        }
        let mut step = 0;
        while !queue.is_empty() && total_one != 0 {
            step += 1;
            for _ in 0..queue.len() {
                let top = queue.pop_front().unwrap();
                for i in -1..=1 {
                    for j in -1..=1 {
                        if i == 0 && j == 0 || i != 0 && j != 0 {
                            continue;
                        }
                        let (x, y) = (top.0 + i, top.1 + j);
                        if x < 0 || x >= m || y < 0 || y >= n {
                            continue;
                        }
                        if grid[x as usize][y as usize] != 1 {
                            continue;
                        }
                        total_one -= 1;
                        grid[x as usize][y as usize] = 0;
                        queue.push_back((x, y));
                    }
                }
            }
        }

        if total_one == 0 {
            return step;
        }
        return -1;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::oranges_rotting(vec![vec![2, 1, 1], vec![1, 1, 0], vec![0, 1, 1],])
    );
}
