// Accepted
// 2019-01-04 19:22
// @lc id=931 lang=rust
// problem: minimum_falling_path_sum
#![feature(core_intrinsics)]
fn type_of<T>(_: &T) -> &str {
    unsafe { std::intrinsics::type_name::<T>() }
}

struct Solution {}
impl Solution {
    pub fn min_falling_path_sum(a: Vec<Vec<i32>>) -> i32 {
        let (m, n) = (a.len(), a[0].len());
        let mut dp_prev = a[0].clone();
        let mut dp = a[0].clone();
        for _i in 1..m {
            for j in 0..n {
                let left = {
                    if j > 0 {
                        dp_prev[j - 1]
                    } else {
                        std::i32::MAX
                    }
                };
                let mid = dp_prev[j];
                let right = {
                    if j < n - 1 {
                        dp_prev[j + 1]
                    } else {
                        std::i32::MAX
                    }
                };
                dp[j] = vec![left, mid, right].into_iter().min().unwrap() + a[_i][j];
            }
            std::mem::swap(&mut dp_prev, &mut dp);
        }
        return dp_prev.into_iter().min().unwrap();
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::min_falling_path_sum(vec![vec![1, 2, 3], vec![4, 5, 6], vec![7, 8, 9],])
    );
}
