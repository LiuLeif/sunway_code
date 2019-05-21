// Accepted
// @lc id=462 lang=rust
// problem: minimum_moves_to_equal_array_elements_ii
struct Solution {}
impl Solution {
    fn partition(nums: &mut [i32]) -> usize {
        let mut lo = 1;
        let mut p = 0;
        while lo < nums.len() {
            if nums[lo] >= nums[p] {
                lo += 1;
            } else {
                nums.swap(lo, p + 1);
                nums.swap(p, p + 1);
                lo += 1;
                p += 1;
            }
        }
        return p;
    }

    fn quick_select(nums: &mut [i32], k: usize) -> i32 {
        let p = Solution::partition(nums);
        if p == k {
            return nums[p];
        }
        if p > k {
            return Solution::quick_select(&mut nums[..p], k);
        } else {
            return Solution::quick_select(&mut nums[p + 1..], k - p - 1);
        }
    }

    pub fn min_moves2(nums: Vec<i32>) -> i32 {
        let mut nums = nums;
        let n = nums.len();
        let median = Solution::quick_select(&mut nums, n / 2);
        nums.into_iter().map(|x| (x - median).abs()).sum()
    }
}

fn main() {
    println!("{:?}", Solution::min_moves2(vec![1, 2, 3]));
}
