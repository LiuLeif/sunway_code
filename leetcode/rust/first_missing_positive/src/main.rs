// Accepted
// @lc id=41 lang=rust
// problem: first_missing_positive
struct Solution {}
impl Solution {
    pub fn first_missing_positive(nums: Vec<i32>) -> i32 {
        // set nums[x-1] to  -nums[x-1]
        let mut nums = nums;
        let n = nums.len();
        for i in 0..n {
            if nums[i] <= 0 {
                nums[i] = n as i32 + 1;
            }
        }

        for i in 0..n {
            let index = nums[i].abs() as usize;
            if index > n {
                continue;
            }
            nums[index - 1] = -nums[index - 1].abs();
        }

        match nums.into_iter().enumerate().skip_while(|t| t.1 < 0).next() {
            Some((x, _)) => x as i32 + 1,
            None => n as i32 + 1,
        }
    }
}

fn main() {
    println!("{:?}", Solution::first_missing_positive(vec![3, 4, -1, 1]));
    println!("{:?}", Solution::first_missing_positive(vec![-2, 1, -3]));
    println!("{:?}", Solution::first_missing_positive(vec![]));
}
