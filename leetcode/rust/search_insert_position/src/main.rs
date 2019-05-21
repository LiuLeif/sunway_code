// Accepted
// @lc id=35 lang=rust
// problem: search_insert_position
struct Solution {}
impl Solution {
    pub fn search_insert(nums: Vec<i32>, target: i32) -> i32 {
        let (mut lo, mut hi) = (0i32, nums.len() as i32 - 1);
        while lo <= hi {
            let mid = (lo + hi) / 2;
            if nums[mid as usize] >= target {
                hi = mid - 1;
            } else {
                lo = mid + 1;
            }
        }
        hi as i32 + 1
    }
}

fn main() {
    println!("{:?}", Solution::search_insert(vec![1, 3, 5, 6], 0));
}
