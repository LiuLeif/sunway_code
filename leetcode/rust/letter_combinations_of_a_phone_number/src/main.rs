// Accepted
struct Solution {}
/*
 * @lc app=leetcode id=17 lang=rust
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
impl Solution {
    pub fn letter_combinations(digits: String) -> Vec<String> {
        if digits.is_empty() {
            return vec![];
        }
        Solution::doit(digits)
    }

    pub fn doit(digits: String) -> Vec<String> {
        if digits.is_empty() {
            return vec!["".to_owned()];
        }

        let head = digits.chars().next().unwrap();
        let ret = Solution::doit((&digits[1..]).to_owned());
        ret.into_iter()
            .flat_map(|s| {
                let characters = match head {
                    '2' => "abc",
                    '3' => "def",
                    '4' => "ghi",
                    '5' => "jkl",
                    '6' => "mno",
                    '7' => "pqrs",
                    '8' => "tuv",
                    '9' => "wxyz",
                    _ => panic!(),
                };
                characters.chars().map(move |c| format!("{}{}", c, s))
            })
            .collect()
    }
}
// @lc code=end

fn main() {
    println!("{:?}", Solution::letter_combinations("234".to_owned()));
}
