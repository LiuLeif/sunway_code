// Accepted
// @lc id=553 lang=rust
// problem: optimal_division
struct Solution {}
impl Solution {
    pub fn optimal_division(nums: Vec<i32>) -> String {
        if nums.len() == 1 {
            return nums[0].to_string();
        }
        if nums.len() == 2 {
            return nums[0].to_string() + "/" + nums[1].to_string().as_str();
        }
        nums[0].to_string()
            + "/("
            + nums
                .iter()
                .skip(1)
                .map(|i| i.to_string())
                .collect::<Vec<String>>()
                .join("/")
                .as_str()
            + ")"
    }
}

fn main() {
    println!("{:?}", Solution::optimal_division(vec![1000, 100, 10, 2]));
    println!("{:?}", Solution::optimal_division(vec![1000]));
    println!("{:?}", Solution::optimal_division(vec![1000, 2]));
}
