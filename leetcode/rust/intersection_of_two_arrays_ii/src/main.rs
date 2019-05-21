// Accepted
// @lc id=350 lang=rust
// problem: intersection_of_two_arrays_ii
struct Solution {}
impl Solution {
    pub fn intersect(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
        let mut nums1 = nums1;
        nums1.sort();
        let mut nums2 = nums2;
        nums2.sort();
        let mut ret = vec![];
        let (mut i, mut j) = (0, 0);
        while i < nums1.len() && j < nums2.len() {
            if nums1[i] == nums2[j] {
                ret.push(nums1[i]);
                i += 1;
                j += 1;
            } else if nums1[i] > nums2[j] {
                j += 1;
            } else {
                i += 1;
            }
        }
        ret
    }
}

fn main() {
    println!("{:?}", Solution::intersect(vec![1, 2, 2, 1], vec![2, 1]));
}
