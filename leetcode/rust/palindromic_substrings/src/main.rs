// Accepted
// 2019-01-07 14:38
// @lc id=647 lang=rust
// problem: palindromic_substrings
struct Solution {}
impl Solution {
    fn is_palindrom(dp: &mut Vec<Vec<bool>>, s: &[u8], f: usize, t: usize) -> bool {
        if f == t {
            return true;
        }
        let ret = {
            if s[f] != s[t] {
                false
            } else {
                if f + 1 == t {
                    true
                } else {
                    dp[f + 1][t - 1]
                }
            }
        };
        dp[f][t] = ret;
        return ret;
    }

    pub fn count_substrings(s: String) -> i32 {
        let mut ret = 0;
        let mut dp = vec![vec![false; s.len()]; s.len()];
        for i in 0..s.len() {
            dp[i][i] = true;
        }
        for i in 0..s.len() {
            for j in 0..s.len() - i {
                if Solution::is_palindrom(&mut dp, s.as_bytes(), j, j + i) {
                    ret += 1
                }
            }
        }
        return ret;
    }
}

fn main() {
    println!("{:?}", Solution::count_substrings("abc".to_owned()));
}
