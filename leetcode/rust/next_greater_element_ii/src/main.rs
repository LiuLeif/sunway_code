// Accepted
// @lc id=503 lang=rust
// problem: next_greater_element_ii
struct Solution {}
impl Solution {
    pub fn next_greater_elements(nums: Vec<i32>) -> Vec<i32> {
        // 1,2,1
        // 2,-1,2
        let mut stack: Vec<usize> = vec![];
        let mut nums = nums;
        nums.extend(nums.clone());

        let mut ret = vec![-1; nums.len() * 2];
        for i in 0..nums.len() {
            while !stack.is_empty() && nums[*stack.last().unwrap()] < nums[i] {
                ret[stack.pop().unwrap() as usize] = nums[i];
            }
            stack.push(i);
        }
        ret[..nums.len() / 2].to_vec()
    }
}

fn main() {
    println!("{:?}", Solution::next_greater_elements(vec![1, 2, 1]));
}
