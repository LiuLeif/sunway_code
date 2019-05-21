// Accepted
// @lc id=5 lang=rust
// problem: longest_palindromic_substring
struct Solution {}
impl Solution {
    pub fn longest_palindrome(s: String) -> String {
        // aacdefcaa
        let s = s.into_bytes();
        let n = s.len();
        let mut dp = vec![vec![false; n]; n];
        let mut ret = "".to_owned();
        for i in 0..n {
            dp[i][i] = true;
            ret = String::from_utf8(s[i..i + 1].to_vec()).unwrap();
        }
        for i in 1..n {
            for j in 0..n - i {
                if s[j] == s[j + i] {
                    if i == 1 {
                        dp[j][j + i] = true;
                    } else {
                        dp[j][j + i] = dp[j + 1][j + i - 1];
                    }
                    if dp[j][j + i] {
                        ret = String::from_utf8(s[j..=j + i].to_vec()).unwrap();
                    }
                }
            }
        }
        ret
    }
}

fn main() {
    println!("{:?}", Solution::longest_palindrome("abxyba".to_owned()));
}
