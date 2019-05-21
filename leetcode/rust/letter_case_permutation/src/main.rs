// Accepted
// 2019-01-07 00:56
// @lc id=784 lang=rust
// problem: letter_case_permutation
struct Solution {}
impl Solution {
    fn backtrack(ret: &mut Vec<String>, curr: String, remain: &str) -> () {
        if remain == "" {
            ret.push(curr);
            return;
        }
        let c = remain.chars().next().unwrap();
        Solution::backtrack(ret, format!("{}{}", curr, c), &remain[1..]);
        if c.is_ascii_lowercase() {
            Solution::backtrack(ret, format!("{}{}", curr, c.to_uppercase()), &remain[1..]);
        } else if c.is_ascii_uppercase() {
            Solution::backtrack(ret, format!("{}{}", curr, c.to_lowercase()), &remain[1..]);
        }
    }

    pub fn letter_case_permutation(s: String) -> Vec<String> {
        let mut ret = vec![];
        Solution::backtrack(&mut ret, "".to_owned(), s.as_str());
        return ret;
    }
}

fn main() {
    println!("{:?}", Solution::letter_case_permutation("abc".to_owned()));
}
