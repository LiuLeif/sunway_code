// Accepted
// @lc id=152 lang=rust
// problem: maximum_product_subarray
struct Solution {}
impl Solution {
    pub fn max_product(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut dp_min = 1;
        let mut dp_max = 1;
        let mut ret = std::i32::MIN;
        for i in 1..n + 1 {
            let min = vec![dp_min * nums[i - 1], dp_max * nums[i - 1], nums[i - 1]]
                .into_iter()
                .min()
                .unwrap();

            let max = vec![dp_min * nums[i - 1], dp_max * nums[i - 1], nums[i - 1]]
                .into_iter()
                .max()
                .unwrap();
            dp_min = min;
            dp_max = max;
            ret = std::cmp::max(ret, dp_max);
        }
        ret
    }
}

fn main() {
    println!("{:?}", Solution::max_product(vec![-2, -2, -3]));
}
