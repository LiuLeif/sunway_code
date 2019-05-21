// Accepted
// @lc id=547 lang=rust
// problem: friend_circles
use std::collections::HashSet;
struct Solution {}
impl Solution {
    fn dfs(x: usize, m: &Vec<Vec<i32>>, visited: &mut HashSet<usize>) -> () {
        if visited.contains(&x) {
            return;
        }
        visited.insert(x);
        for i in 0..m.len() {
            if m[x][i] != 0 {
                Solution::dfs(i, m, visited);
            }
        }
    }

    pub fn find_circle_num(m: Vec<Vec<i32>>) -> i32 {
        let mut visited = HashSet::<usize>::new();
        let mut ret = 0;
        for i in 0..m.len() {
            if visited.contains(&i) {
                continue;
            }
            ret += 1;
            Solution::dfs(i, &m, &mut visited);
        }
        return ret;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::find_circle_num(vec![vec![1, 1, 0], vec![1, 1, 1], vec![0, 1, 1],])
    );
}
