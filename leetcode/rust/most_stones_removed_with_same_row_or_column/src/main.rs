// Accepted
// @lc id=947 lang=rust
// problem: most_stones_removed_with_same_row_or_column
use std::collections::HashMap;
use std::collections::HashSet;
struct Solution {}
impl Solution {
    fn dfs(
        x: i32,
        y: i32,
        rows: &HashMap<i32, Vec<i32>>,
        cols: &HashMap<i32, Vec<i32>>,
        visited: &mut HashSet<(i32, i32)>,
    ) {
        if visited.contains(&(x, y)) {
            return;
        }
        visited.insert((x, y));
        for i in &rows[&x] {
            Solution::dfs(x, *i, rows, cols, visited);
        }
        for i in &cols[&y] {
            Solution::dfs(*i, y, rows, cols, visited);
        }
    }

    pub fn remove_stones(stones: Vec<Vec<i32>>) -> i32 {
        // x x
        // x   x
        //   x x
        let ret = stones.len() as i32;
        let mut rows = HashMap::new();
        let mut cols = HashMap::new();
        for stone in stones {
            rows.entry(stone[0]).or_insert(vec![]).push(stone[1]);
            cols.entry(stone[1]).or_insert(vec![]).push(stone[0]);
        }

        let mut visited = HashSet::new();

        let mut islands = 0;
        for (x, v) in rows.iter() {
            for y in v {
                if visited.contains(&(*x, *y)) {
                    continue;
                }
                islands += 1;
                Solution::dfs(*x, *y, &rows, &cols, &mut visited);
            }
        }
        ret - islands
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::remove_stones(vec![
            vec![0, 0],
            vec![0, 1],
            vec![1, 0],
            vec![1, 2],
            vec![2, 1],
            vec![2, 2],
        ])
    );
}
