// Accepted
// @lc id=934 lang=rust
// problem: shortest_bridge
use std::collections::VecDeque;
struct Solution {}
impl Solution {
    fn dfs(a: &mut Vec<Vec<i32>>, x: i32, y: i32) {
        let (m, n) = (a.len() as i32, a[0].len() as i32);
        a[x as usize][y as usize] = 2;
        for i in vec![-1, 0, 1] {
            for j in vec![-1, 0, 1] {
                if (i == 0 && j == 0) || (i != 0 && j != 0) {
                    continue;
                }
                let (x, y) = (x + i, y + j);
                if x < 0 || x >= m || y < 0 || y >= n {
                    continue;
                }
                if a[x as usize][y as usize] == 1 {
                    Solution::dfs(a, x, y);
                }
            }
        }
    }
    fn mark_first(a: &mut Vec<Vec<i32>>) {
        let (x, y) = {
            let (mut x, mut y) = (0, 0);
            'again: for i in 0..a.len() {
                for j in 0..a[0].len() {
                    if a[i][j] == 1 {
                        x = i;
                        y = j;
                        break 'again;
                    }
                }
            }
            (x as i32, y as i32)
        };

        Solution::dfs(a, x, y);
    }

    fn mark_second(a: &mut Vec<Vec<i32>>) -> i32 {
        let (m, n) = (a.len() as i32, a[0].len() as i32);
        let mut queue = VecDeque::new();
        for i in 0..m {
            for j in 0..n {
                if a[i as usize][j as usize] == 1 {
                    queue.push_back((i, j, 0));
                }
            }
        }
        while !queue.is_empty() {
            let top = queue.pop_front().unwrap();
            for i in vec![-1, 0, 1] {
                for j in vec![-1, 0, 1] {
                    if (i == 0 && j == 0) || (i != 0 && j != 0) {
                        continue;
                    }
                    let (x, y) = (top.0 + i, top.1 + j);
                    if x < 0 || x >= m || y < 0 || y >= n {
                        continue;
                    }
                    if a[x as usize][y as usize] == 1 {
                        continue;
                    }
                    if a[x as usize][y as usize] == 2 {
                        return top.2;
                    }
                    a[x as usize][y as usize] = 1;
                    queue.push_back((x, y, top.2 + 1));
                }
            }
        }
        unreachable!();
    }

    pub fn shortest_bridge(a: Vec<Vec<i32>>) -> i32 {
        let mut a = a;
        Solution::mark_first(&mut a);
        Solution::mark_second(&mut a)
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::shortest_bridge(vec![vec![0, 1, 0], vec![0, 0, 0], vec![0, 0, 1],])
    );
}
