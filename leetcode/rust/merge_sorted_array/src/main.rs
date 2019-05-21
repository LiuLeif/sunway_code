// Accepted
// @lc id=88 lang=rust
// problem: merge_sorted_array
struct Solution {}
impl Solution {
    pub fn merge(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
        let mut m = m as usize;
        let mut n = n as usize;
        for i in (0..m + n).rev() {
            if m == 0 || (n != 0 && nums2[n - 1] >= nums1[m - 1]) {
                nums1[i] = nums2[n - 1];
                n -= 1;
            } else {
                nums1[i] = nums1[m - 1];
                m -= 1;
            }
        }
    }
}

fn main() {
    let mut a = vec![1, 2, 3, 0, 0, 0];
    Solution::merge(&mut a, 3, &mut vec![2, 5, 6], 3);
    println!("{:?}", a);
}
