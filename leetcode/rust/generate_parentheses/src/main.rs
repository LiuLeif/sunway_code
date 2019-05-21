// Accepted
// @lc id=22 lang=rust
// problem: generate_parentheses
struct Solution {}
impl Solution {
    // [
    //   "((()))",
    //   "(()())",
    //   "(())()",
    //   "()(())",
    //   "()()()"
    // ]
    fn backtrack(ret: &mut Vec<String>, curr: String, left: i32, right: i32, n: i32) {
        if left == n && right == n {
            ret.push(curr);
            return;
        }
        if left < n {
            Solution::backtrack(ret, format!("{}(", curr), left + 1, right, n);
        }

        if left > right {
            Solution::backtrack(ret, format!("{})", curr), left, right + 1, n);
        }
    }

    pub fn generate_parenthesis(n: i32) -> Vec<String> {
        let mut ret = vec![];
        Solution::backtrack(&mut ret, "".to_owned(), 0, 0, n);
        ret
    }
}

fn main() {
    println!("{:?}", Solution::generate_parenthesis(3));
}
