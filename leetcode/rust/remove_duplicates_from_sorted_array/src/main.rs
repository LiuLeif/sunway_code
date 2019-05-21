// Accepted
// @lc id=26 lang=rust
// problem: remove_duplicates_from_sorted_array
struct Solution {}
impl Solution {
    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
        if nums.is_empty() {
            return 0;
        }
        let mut lo = 0;
        for i in 1..nums.len() {
            if nums[i] != nums[lo] {
                nums.swap(lo + 1, i);
                lo += 1;
            }
        }
        return lo as i32 + 1;
    }
}

fn main() {
    let mut v = vec![1, 1, 1, 2];
    println!("{:?}", Solution::remove_duplicates(&mut v));
    println!("{:?}", v);
}
