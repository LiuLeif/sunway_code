// 2019-01-03 14:56
// @lc id=442 lang=rust
// problem: find_all_duplicates_in_an_array
struct Solution {}
impl Solution {
    pub fn find_duplicates(nums: Vec<i32>) -> Vec<i32> {
        // 1 2 3 4 5 6 7 8
        let mut nums = nums;
        let mut ret = vec![];
        for i in 0..nums.len() {
            let index = nums[i].abs() as usize;
            if nums[index - 1] < 0 {
                ret.push(index as i32);
            } else {
                nums[index - 1] = -nums[index - 1];
            }
        }
        return ret;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::find_duplicates(vec![4, 3, 2, 7, 8, 2, 3, 1])
    );
}
