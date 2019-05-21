// Accepted
// @lc id=179 lang=rust
// problem: largest_number
struct Solution {}
impl Solution {
    pub fn largest_number(nums: Vec<i32>) -> String {
        let mut nums = nums;
        nums.sort_by(|a, b| format!("{}{}", b, a).cmp(&format!("{}{}", a, b)));
        let mut ret = nums
            .into_iter()
            .map(|x| x.to_string())
            .collect::<Vec<String>>()
            .join("");
        ret = ret.trim_start_matches("0").to_string();
        if ret.is_empty() {
            return "0".to_owned();
        }
        ret
    }
}

fn main() {
    println!("{:?}", Solution::largest_number(vec![3, 30, 34, 5, 9]));
    println!("{:?}", Solution::largest_number(vec![0, 0]));
}
