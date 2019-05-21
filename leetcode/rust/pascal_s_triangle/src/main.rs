// Accepted
// @lc id=118 lang=rust
// problem: pascal_s_triangle
struct Solution {}
impl Solution {
    pub fn generate(num_rows: i32) -> Vec<Vec<i32>> {
        if num_rows == 0 {
            return vec![];
        }
        let num_rows = num_rows as usize;
        let mut ret = vec![vec![]; num_rows];
        ret[0] = vec![1];
        for i in 1..num_rows {
            ret[i] = vec![0; i + 1];
            for j in 0..=i {
                let a = {
                    if j == 0 {
                        0
                    } else {
                        ret[i - 1][j - 1]
                    }
                };
                let b = {
                    if j == i {
                        0
                    } else {
                        ret[i - 1][j]
                    }
                };
                ret[i][j] = a + b;
            }
        }
        ret
    }
}

fn main() {
    println!("{:?}", Solution::generate(5));
}
