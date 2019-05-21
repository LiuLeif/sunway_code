// Accepted
// 2019-01-07 11:04
// @lc id=856 lang=rust
// problem: score_of_parentheses
struct Solution {}
impl Solution {
    pub fn score_of_parentheses(s: String) -> i32 {
        let mut stack = vec![];
        let mut curr_score = 0;
        s.chars().for_each(|c| match c {
            '(' => {
                stack.push(curr_score);
                curr_score = 0;
            }
            ')' => {
                if curr_score == 0 {
                    curr_score = 1
                } else {
                    curr_score *= 2;
                }
                if let Some(top) = stack.pop() {
                    curr_score += top;
                }
            }
            _ => (),
        });
        curr_score
    }
}

fn main() {
    println!("{:?}", Solution::score_of_parentheses("()".to_owned()));
    println!(
        "{:?}",
        Solution::score_of_parentheses("(()(()))".to_owned())
    );
}
