// Accepted
// @lc id=60 lang=rust
// problem: permutation_sequence
struct Solution {}
impl Solution {
    fn dfs(mut nums: Vec<String>, k: i32) -> String {
        if nums.is_empty() {
            return "".to_owned();
        }
        let n = nums.len();
        let fact = (1..n as i32).product::<i32>();
        let index = (k / fact) as usize;
        let k = k % fact;
        let head = nums[index].clone();
        nums.remove(index);
        head + &Solution::dfs(nums, k)
    }

    pub fn get_permutation(n: i32, k: i32) -> String {
        let nums = (1..=n).map(|x| x.to_string()).collect::<Vec<String>>();
        Solution::dfs(nums, k - 1)
    }
}

fn main() {
    println!("{:?}", Solution::get_permutation(3, 6));
}
