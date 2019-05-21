// Accepted
// @lc id=238 lang=rust
// problem: product_of_array_except_self
struct Solution {}
impl Solution {
    pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
        // 1,2,3,4
        let mut left = vec![0; nums.len()];
        let mut right = vec![0; nums.len()];
        let mut product = 1;
        for i in 0..nums.len() {
            product *= nums[i];
            left[i] = product;
        }

        let mut product = 1;
        for i in (0..nums.len()).rev() {
            product *= nums[i];
            right[i] = product;
        }

        left.insert(0, 1);
        right.push(1);

        left.into_iter()
            .take(nums.len())
            .zip(right.into_iter().skip(1))
            .map(|t| t.0 * t.1)
            .collect()
    }
}

fn main() {
    println!("{:?}", Solution::product_except_self(vec![1, 2, 3, 4]));
}
