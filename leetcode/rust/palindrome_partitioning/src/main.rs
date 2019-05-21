// Accepted
// @lc id=131 lang=rust
// problem: palindrome_partitioning
struct Solution {}
impl Solution {
    fn backtrack(
        ret: &mut Vec<Vec<String>>,
        s: &[u8],
        dp: &Vec<Vec<bool>>,
        curr: &mut Vec<String>,
        index: usize,
    ) {
        if index == s.len() {
            ret.push(curr.clone());
            return;
        }
        for i in index..s.len() {
            if !dp[index][i] {
                continue;
            }
            curr.push(String::from_utf8(s[index..i + 1].to_vec()).unwrap());
            Solution::backtrack(ret, s, dp, curr, i + 1);
            curr.pop();
        }
    }

    pub fn partition(s: String) -> Vec<Vec<String>> {
        let s = s.into_bytes();
        let n = s.len();
        let mut dp = vec![vec![false; n]; n];
        for i in 0..n {
            dp[i][i] = true;
        }
        // abc
        for interval in 1..n {
            for i in 0..n - interval {
                if interval == 1 {
                    dp[i][i + interval] = s[i] == s[i + interval]
                } else {
                    dp[i][i + interval] = (s[i] == s[i + interval]) && dp[i + 1][i + interval - 1];
                }
            }
        }
        let mut ret = vec![];
        Solution::backtrack(&mut ret, &s, &dp, &mut vec![], 0);
        ret
    }
}

fn main() {
    println!("{:?}", Solution::partition("aaxxa".to_owned()));
}
