// Accepted
// @lc id=4 lang=rust
// problem: median_of_two_sorted_arrays
struct Solution {}
impl Solution {
    fn find(nums1: &[i32], nums2: &[i32], n: usize) -> i32 {
        if nums1.is_empty() {
            return nums2[n - 1];
        }
        if nums2.is_empty() {
            return nums1[n - 1];
        }
        let (m1, m2) = (nums1.len() / 2, nums2.len() / 2);
        let (n1, n2) = (nums1[m1], nums2[m2]);
        if m1 + m2 >= n - 1 {
            if n1 > n2 {
                return Solution::find(&nums1[..m1], nums2, n);
            } else {
                return Solution::find(nums1, &nums2[..m2], n);
            }
        } else {
            if n1 > n2 {
                return Solution::find(nums1, &nums2[m2 + 1..], n - m2 - 1);
            } else {
                return Solution::find(&nums1[m1 + 1..], nums2, n - m1 - 1);
            }
        }
    }

    pub fn find_median_sorted_arrays(nums1: Vec<i32>, nums2: Vec<i32>) -> f64 {
        let (l1, l2) = (nums1.len(), nums2.len());
        if (l1 + l2) % 2 == 0 {
            return (Solution::find(&nums1, &nums2, (l1 + l2) / 2) as f64
                + Solution::find(&nums1, &nums2, (l1 + l2) / 2 + 1) as f64)
                / 2.0;
        } else {
            return Solution::find(&nums1, &nums2, (l1 + l2) / 2 + 1) as f64;
        }
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::find_median_sorted_arrays(vec![1, 2], vec![3, 4])
    );
}
