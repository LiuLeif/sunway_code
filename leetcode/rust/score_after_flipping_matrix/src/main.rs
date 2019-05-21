// 2019-01-03 10:09
// @lc id=861 lang=rust
// problem: score_after_flipping_matrix
struct Solution {}
impl Solution {
    pub fn matrix_score(a: Vec<Vec<i32>>) -> i32 {
        // [0,0,1,1]
        // [1,0,1,0]
        // [1,1,0,0]
        // ->
        // [1,1,1,1],
        // [1,0,0,1],
        // [1,1,1,1],
        let mut a = a;
        let (m, n) = (a.len(), a[0].len());
        for i in 0..m {
            if a[i][0] == 0 {
                for j in 0..n {
                    a[i][j] = (a[i][j] + 1) % 2;
                }
            }
        }
        for j in 0..n {
            let mut count_one = 0;
            for i in 0..m {
                if a[i][j] == 1 {
                    count_one += 1;
                }
            }
            if count_one <= m / 2 {
                for i in 0..m {
                    a[i][j] = (a[i][j] + 1) % 2;
                }
            }
        }

        let mut ret = 0;
        let mut power = 1;
        for j in 0..n {
            for i in 0..m {
                if a[i][n - j - 1] == 1 {
                    ret += power;
                }
            }
            power *= 2
        }
        return ret;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::matrix_score(vec![vec![0, 0, 1, 1], vec![1, 0, 1, 0], vec![1, 1, 0, 0],])
    );
}
