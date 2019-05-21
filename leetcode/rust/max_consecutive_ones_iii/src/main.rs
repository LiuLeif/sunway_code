// Accepted
// @lc id=1004 lang=rust
// problem: max_consecutive_ones_iii
struct Solution {}
impl Solution {
    // pub fn longest_ones_dp(a: Vec<i32>, k: i32) -> i32 {
    //     if a.iter().filter(|&&x| x == 0).count() <= k as usize {
    //         return a.len() as i32;
    //     }
    //     let n = a.len() + 1;
    //     let k = k as usize + 1;
    //     let mut dp = vec![vec![0; k]; n];
    //     let mut ret = 0;
    //     for i in 1..n {
    //         for j in 0..k {
    //             if a[i - 1] == 1 {
    //                 dp[i][j] = dp[i - 1][j] + 1;
    //             } else if j > 0 {
    //                 dp[i][j] = dp[i - 1][j - 1] + 1;
    //             }
    //             ret = std::cmp::max(dp[i][j], ret);
    //         }
    //     }
    //     ret
    // }
    pub fn longest_ones(a: Vec<i32>, k: i32) -> i32 {
        let mut k = k;
        let (mut lo, mut hi) = (0, 0);
        let mut ret = 0;
        while hi < a.len() {
            if k >= 0 {
                ret = std::cmp::max(ret, hi - lo);
                if a[hi] == 0 {
                    k -= 1;
                }
                hi += 1;
            } else {
                if a[lo] == 0 {
                    k += 1;
                }
                lo += 1;
            }
        }
        if k >= 0 {
            ret = std::cmp::max(ret, hi - lo);
        }
        ret as i32
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::longest_ones(vec![1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0], 2)
    );
    println!("{:?}", Solution::longest_ones(vec![1, 0, 0, 0], 3));
}
