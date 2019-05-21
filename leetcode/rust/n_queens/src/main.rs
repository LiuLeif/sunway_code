// Accepted
// @lc id=51 lang=rust
// problem: n_queens
use std::collections::HashSet;
struct Solution {}
impl Solution {
    fn backtrack(
        ret: &mut Vec<Vec<String>>,
        curr: &mut Vec<String>,
        col_count: &mut HashSet<i32>,
        slant_count_2: &mut HashSet<i32>,
        slant_count_1: &mut HashSet<i32>,
        i: i32,
        n: i32,
    ) {
        if i == n {
            ret.push(curr.clone());
            return;
        }
        for j in 0..n {
            if col_count.contains(&j)
                || slant_count_2.contains(&(i - j))
                || slant_count_1.contains(&(j + i))
            {
                continue;
            }
            let mut tmp = vec!['.'; n as usize];
            tmp[j as usize] = 'Q';
            curr.push(tmp.into_iter().collect::<String>());
            col_count.insert(j);
            slant_count_2.insert(i - j);
            slant_count_1.insert(j + i);
            Solution::backtrack(ret, curr, col_count, slant_count_2, slant_count_1, i + 1, n);
            curr.pop();
            col_count.remove(&j);
            slant_count_2.remove(&(i - j));
            slant_count_1.remove(&(j + i));
        }
    }

    pub fn solve_n_queens(n: i32) -> Vec<Vec<String>> {
        let mut ret = vec![];
        let mut col_count = HashSet::new();
        let mut slant_count_2 = HashSet::new();
        let mut slant_count_1 = HashSet::new();

        Solution::backtrack(
            &mut ret,
            &mut vec![],
            &mut col_count,
            &mut slant_count_2,
            &mut slant_count_1,
            0,
            n,
        );
        ret
    }
}

fn main() {
    println!("{:?}", Solution::solve_n_queens(4).len());
}
